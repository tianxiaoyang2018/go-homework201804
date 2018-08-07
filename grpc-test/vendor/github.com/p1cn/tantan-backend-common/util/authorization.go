package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rc4"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/icza/bitio"
)

const (
	AuthTypeBasic  = "Basic"
	AuthTypeBearer = "Bearer"
	AuthTypeMac    = "MAC"

	MagicUserIDForAntispam              = 53060000
	MaxFailedTimesForBuildInfoHashCheck = 7
)

type HMacSecrets struct {
	IOS     map[string]HMacSecret
	Android map[string]HMacSecret
}

type HMacSecret struct {
	Secret          string
	AppVersionRange [2]string
}

type appVersionRange [][2]string

func (s appVersionRange) Len() int           { return len(s) }
func (s appVersionRange) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s appVersionRange) Less(i, j int) bool { return VersionGreaterThanOrEqualTo(s[j][0], s[i][0]) }

func (s HMacSecrets) IsValid() bool {
	return !s.isAppVersionOverlapped(s.Android) && !s.isAppVersionOverlapped(s.IOS)
}

func (s HMacSecrets) isAppVersionOverlapped(secrets map[string]HMacSecret) bool {
	versionRanges := make(appVersionRange, 0, len(secrets))
	for _, v := range secrets {
		versionRanges = append(versionRanges, v.AppVersionRange)
	}

	sort.Sort(versionRanges)

	for i := 0; i < len(versionRanges)-1; i++ {
		if VersionGreaterThan(versionRanges[i][1], versionRanges[i+1][0]) {
			return true
		}
	}
	return false
}

type HiddenError string

func (he HiddenError) Error() string {
	return string(he)
}

var (
	ErrBadHMAC                      = errors.New("bad hmac")
	ErrBadBoundary                  = HiddenError("bad boundary")
	ErrAndroidBoundaryBotMarkCP     = HiddenError("android boundary botMark copy-paste")
	ErrAndroidBoundaryBotMarkMotion = HiddenError("android boundary botMark motion")
	ErrAndroidBoundaryBotMarkCPM    = HiddenError("android boundary botMark cp and motion")
	ErrAndroidBoundaryXPosed        = HiddenError("android boundary xposed")
	ErrAndroidBoundaryInEmulator    = HiddenError("android boundary emulator")
)

func IsClientOSSecretVersionMismatch(secretVer, clientOS string) error {
	var osInSecretVer string
	if strings.HasPrefix(secretVer, OSAndroid) {
		osInSecretVer = OSAndroid
	}
	if strings.HasPrefix(secretVer, OSIOS) {
		osInSecretVer = OSIOS
	}
	if osInSecretVer != clientOS {
		return HiddenError(fmt.Sprintf("secret version %s client os %s mismatch", secretVer, clientOS))
	}
	return nil
}

func IsAppVersionSecretVersionMismatch(secretVer, clientOS, appVer string, secrets HMacSecrets) error {
	if clientOS == OSAndroid {
		hMacSecrets := secrets.Android
		for secretVersion, secret := range hMacSecrets {
			if VersionGreaterThanOrEqualTo(appVer, secret.AppVersionRange[0]) &&
				VersionGreaterThan(secret.AppVersionRange[1], appVer) {
				if secretVer != secretVersion {
					return HiddenError(fmt.Sprintf("secret version %s app version %s mismatch", secretVer, appVer))
				} else {
					return nil
				}
			}
		}
	}

	if clientOS == OSIOS {
		hMacSecrets := secrets.IOS
		for secretVersion, secret := range hMacSecrets {
			if VersionGreaterThanOrEqualTo(appVer, secret.AppVersionRange[0]) &&
				VersionGreaterThan(secret.AppVersionRange[1], appVer) {
				if secretVer != secretVersion {
					return HiddenError(fmt.Sprintf("secret version %s app version %s mismatch", secretVer, appVer))
				} else {
					return nil
				}
			}
		}
	}

	return nil
}

func IsHiddenError(err error) (string, bool) {
	he, ok := err.(HiddenError)
	return string(he), ok
}

type MacHeader struct {
	Version          string
	SecretVersion    string
	Timestamp        int64
	TimestampMs      int64
	TimestampNs      int64
	AccessToken      string
	BuildInfoHash    string
	DeviceIdentifier string
	UriPath          string
	Message          string
	LogMessage       string
	Mac              string
	AppVersion       string
	ClientOS         string
}

func ParseAuthBasic(header http.Header) string {
	return parseAuthHeader(header, AuthTypeBasic)
}

func ParseAuthMac(request *http.Request, appVer string, clientOS string, allowedTimeDiff int, secrets HMacSecrets) (*MacHeader, error) {
	token := ParseAuthBearerAccessToken(request.Header)
	if token != "" {
		return &MacHeader{Version: "0", AccessToken: token, AppVersion: appVer}, nil
	}

	macHeader, err := parseAuthMacHeader(request)
	if err != nil {
		return macHeader, err
	}
	if macHeader != nil && macHeader.SecretVersion != "" {
		macHeader.AppVersion = appVer
		macHeader.ClientOS = clientOS
		err := macHeader.validate(allowedTimeDiff, secrets)
		return macHeader, err
	}

	return macHeader, nil
}

func ParseAuthBoundary(boundary, appVer string, hmacTime int64) error {
	if boundary == "" || VersionGreaterThan("2.4.0", appVer) {
		return nil
	}
	if len(boundary) != 36 {
		return ErrBadBoundary
	}
	tmpbb := []byte(boundary)
	bb := tmpbb[:0]
	for i, v := range tmpbb {
		switch i {
		case 8, 13, 14, 18, 19, 23:
		default:
			bb = append(bb, v)
		}
	}
	if len(bb) != 30 {
		return ErrBadBoundary
	}
	keyAndValue := make([]byte, 15)
	hex.Decode(keyAndValue, bb)

	cipher, _ := rc4.NewCipher(keyAndValue[:7])
	cipher.XORKeyStream(keyAndValue[7:], keyAndValue[7:])

	if VersionGreaterThanOrEqualTo(appVer, "2.4.2") {
		if keyAndValue[14] != keyAndValue[10]^keyAndValue[11]^keyAndValue[12] ||
			keyAndValue[13] != keyAndValue[7]^keyAndValue[8]^keyAndValue[9] {
			return ErrBadBoundary
		}
	}

	br := bitio.NewReader(bytes.NewReader(keyAndValue[7:]))

	// byte 0
	ucp, _ := br.ReadBits(4)
	um, _ := br.ReadBits(1)
	ux, _ := br.ReadBits(1)
	ue, _ := br.ReadBits(1)
	uptr1, _ := br.ReadBits(1)

	if (ucp == 15 || ucp == 13) && um != 1 {
		return ErrAndroidBoundaryBotMarkCPM
	}
	if ucp == 15 || ucp == 13 {
		return ErrAndroidBoundaryBotMarkCP
	}
	if um != 1 {
		return ErrAndroidBoundaryBotMarkMotion
	}

	if VersionGreaterThanOrEqualTo(appVer, "2.5.1") {
		if ux != 0 {
			return ErrAndroidBoundaryXPosed
		}
		if ue != 0 {
			return ErrAndroidBoundaryInEmulator
		}

		hmtStr := strconv.FormatInt(hmacTime, 2)
		l := len(hmtStr)
		if l < 41 {
			return ErrBadBoundary
		}

		hmtStr = hmtStr[:l-10]
		hmtStr = hmtStr[len(hmtStr)-24:]

		hmt, err := strconv.ParseUint(hmtStr, 2, 64)
		if err != nil {
			return ErrBadBoundary
		}

		// byte 1 2 3
		hiddenTime, _ := br.ReadBits(24)

		if AbsInt64(int64(hiddenTime-hmt)) > 10 {
			return ErrBadBoundary
		}

		uptr2, _ := br.ReadBits(1)

		if VersionGreaterThanOrEqualTo(appVer, "2.5.3") {
			if uptr1 == 1 || uptr2 == 1 {
				return HiddenError(fmt.Sprintf("android boundary ts %d %d", uptr1, uptr2))
			}
		}
	}

	return nil
}

func ParseAuthBearerAccessToken(header http.Header) string {
	return parseAuthHeader(header, AuthTypeBearer)
}

func parseAuthHeader(header http.Header, name string) string {
	var val string
	str := strings.TrimSpace(header.Get("Authorization"))
	arr := strings.SplitN(str, " ", 2)
	for k, v := range arr {
		arr[k] = strings.TrimSpace(v)
	}
	if len(arr) == 2 && arr[0] == name {
		val = arr[1]
	}

	return val
}

func parseAuthMacHeader(request *http.Request) (*MacHeader, error) {
	uriPath := request.URL.Path
	if !strings.HasPrefix(request.URL.Path, "/v1") {
		uriPath = "/v1" + uriPath
	}
	macAuth := parseAuthHeader(request.Header, AuthTypeMac)
	if len(macAuth) == 0 {
		return nil, nil
	}

	var body []byte
	var err error
	if strings.HasPrefix(request.Header.Get("Content-Type"), "application/json") {
		var buf bytes.Buffer
		tr := io.TeeReader(request.Body, &buf)
		body, err = ioutil.ReadAll(tr)
		if err != nil {
			return nil, err
		}
		request.Body = ioutil.NopCloser(&buf)
	}

	// v1: version, sec, timestamp, access_token, mac
	macInfo := []string{}
	if err := json.Unmarshal([]byte(macAuth), &macInfo); err != nil {
		return nil, err
	}
	if len(macInfo) == 0 {
		return nil, fmt.Errorf("Version missing from MAC header")
	}
	switch macInfo[0] {
	case "1":
		if len(macInfo) != 5 {
			return nil, fmt.Errorf("MAC/1 should have %v fields but was %v", 5, len(macInfo))
		}
		ts, err := strconv.ParseInt(macInfo[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("MAC/1 could not parse timestamp %v", macInfo[2])
		}
		message := fmt.Sprintf("%s.%s", macInfo[2], macInfo[3])
		return &MacHeader{
			Version:       macInfo[0],
			SecretVersion: macInfo[1],
			Timestamp:     ts,
			AccessToken:   macInfo[3],
			Message:       message,
			Mac:           macInfo[4],
		}, nil
	case "2":
		if !(len(macInfo) == 6 && len(macInfo[4]) > 0) {
			// make sure build_info_non_empty
			return nil, fmt.Errorf("MAC/2 should have %v fields but was %v", 6, len(macInfo))
		}
		ts, err := strconv.ParseInt(macInfo[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("MAC/2 could not parse timestamp %v", macInfo[2])
		}
		message := fmt.Sprintf("%s.%s.%s", macInfo[2], macInfo[3], macInfo[4])
		return &MacHeader{
			Version:       macInfo[0],
			SecretVersion: macInfo[1],
			Timestamp:     ts,
			AccessToken:   macInfo[3],
			BuildInfoHash: macInfo[4],
			Message:       message,
			Mac:           macInfo[5],
		}, nil
	case "3":
		if len(macInfo) != 6 {
			return nil, fmt.Errorf("MAC/3 should have %v fields but was %v", 6, len(macInfo))
		}
		ts, err := strconv.ParseInt(macInfo[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("MAC/3 could not parse timestamp %v", macInfo[2])
		}

		message := fmt.Sprintf("%s.%s.%s.%s", macInfo[2], macInfo[3], macInfo[4], uriPath)
		return &MacHeader{

			Version:       macInfo[0],
			SecretVersion: macInfo[1],
			TimestampMs:   ts,
			AccessToken:   macInfo[3],
			BuildInfoHash: macInfo[4],
			UriPath:       uriPath,
			Message:       message,
			Mac:           macInfo[5],
		}, nil
	case "4":
		if len(macInfo) != 5 {
			return nil, fmt.Errorf("MAC/4 should have %v fields but was %v", 5, len(macInfo))
		}
		ts, err := strconv.ParseInt(macInfo[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("MAC/4 could not parse timestamp %v", macInfo[2])
		}

		message := fmt.Sprintf("%s.%s.%s", macInfo[2], macInfo[3], uriPath)
		return &MacHeader{
			Version:          macInfo[0],
			SecretVersion:    macInfo[1],
			TimestampMs:      ts,
			DeviceIdentifier: macInfo[3],
			UriPath:          uriPath,
			Message:          message,
			Mac:              macInfo[4],
		}, nil
	case "5":
		if len(macInfo) != 6 {
			return nil, fmt.Errorf("MAC/5 should have %v fields but was %v", 6, len(macInfo))
		}
		ts, err := strconv.ParseInt(macInfo[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("MAC/5 could not parse timestamp %v", macInfo[2])
		}

		message := fmt.Sprintf("%s%s%s%s", macInfo[2], macInfo[3], macInfo[4], uriPath)
		return &MacHeader{
			Version:       macInfo[0],
			SecretVersion: macInfo[1],
			TimestampMs:   ts,
			AccessToken:   macInfo[3],
			BuildInfoHash: macInfo[4],
			UriPath:       uriPath,
			Message:       message,
			Mac:           macInfo[5],
		}, nil
	case "6":
		if len(macInfo) != 5 {
			return nil, fmt.Errorf("MAC/6 should have %v fields but was %v", 5, len(macInfo))
		}
		ts, err := strconv.ParseInt(macInfo[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("MAC/6 could not parse timestamp %v", macInfo[2])
		}

		message := fmt.Sprintf("%s%s%s", macInfo[2], macInfo[3], uriPath)
		return &MacHeader{
			Version:          macInfo[0],
			SecretVersion:    macInfo[1],
			TimestampMs:      ts,
			DeviceIdentifier: macInfo[3],
			UriPath:          uriPath,
			Message:          message,
			Mac:              macInfo[4],
		}, nil
	case "7":
		if len(macInfo) != 6 {
			return nil, fmt.Errorf("MAC/7 should have %v fields but was %v", 6, len(macInfo))
		}
		ts, err := strconv.ParseInt(macInfo[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("MAC/7 could not parse timestamp %v", macInfo[2])
		}

		message := fmt.Sprintf("%s%s%s%s%s", macInfo[2], macInfo[3], macInfo[4], uriPath, body)

		var logMessage string
		logBody, err := MaskSensitiveFields(request.URL.Path, body)
		if err == nil {
			logMessage = fmt.Sprintf("%s%s%s%s%s", macInfo[2], macInfo[3], macInfo[4], uriPath, logBody)
		}
		return &MacHeader{
			Version:       macInfo[0],
			SecretVersion: macInfo[1],
			TimestampMs:   ts,
			AccessToken:   macInfo[3],
			BuildInfoHash: macInfo[4],
			UriPath:       uriPath,
			Message:       message,
			LogMessage:    logMessage,
			Mac:           macInfo[5],
		}, nil
	case "8":
		if len(macInfo) != 5 {
			return nil, fmt.Errorf("MAC/8 should have %v fields but was %v", 5, len(macInfo))
		}
		ts, err := strconv.ParseInt(macInfo[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("MAC/8 could not parse timestamp %v", macInfo[2])
		}

		message := fmt.Sprintf("%s%s%s%s", macInfo[2], macInfo[3], uriPath, body)

		var logMessage string
		logBody, err := MaskSensitiveFields(request.URL.Path, body)
		if err == nil {
			logMessage = fmt.Sprintf("%s%s%s%s%s", macInfo[2], macInfo[3], macInfo[4], uriPath, logBody)
		}
		return &MacHeader{
			Version:          macInfo[0],
			SecretVersion:    macInfo[1],
			TimestampMs:      ts,
			DeviceIdentifier: macInfo[3],
			UriPath:          uriPath,
			Message:          message,
			LogMessage:       logMessage,
			Mac:              macInfo[4],
		}, nil
	case "9":
		if len(macInfo) != 6 {
			return nil, fmt.Errorf("MAC/9 should have %v fields but was %v", 6, len(macInfo))
		}
		ts, err := strconv.ParseInt(macInfo[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("MAC/9 could not parse timestamp %v", macInfo[2])
		}

		message := fmt.Sprintf("%s%s%s%s%s", macInfo[2], macInfo[3], macInfo[4], uriPath, body)

		var logMessage string
		logBody, err := MaskSensitiveFields(request.URL.Path, body)
		if err == nil {
			logMessage = fmt.Sprintf("%s%s%s%s%s", macInfo[2], macInfo[3], macInfo[4], uriPath, logBody)
		}
		return &MacHeader{
			Version:       macInfo[0],
			SecretVersion: macInfo[1],
			TimestampNs:   ts,
			AccessToken:   macInfo[3],
			BuildInfoHash: macInfo[4],
			UriPath:       uriPath,
			Message:       message,
			LogMessage:    logMessage,
			Mac:           macInfo[5],
		}, nil
	case "10":
		if len(macInfo) != 5 {
			return nil, fmt.Errorf("MAC/10 should have %v fields but was %v", 5, len(macInfo))
		}
		ts, err := strconv.ParseInt(macInfo[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("MAC/10 could not parse timestamp %v", macInfo[2])
		}

		message := fmt.Sprintf("%s%s%s%s", macInfo[2], macInfo[3], uriPath, body)

		var logMessage string
		logBody, err := MaskSensitiveFields(request.URL.Path, body)
		if err == nil {
			logMessage = fmt.Sprintf("%s%s%s%s%s", macInfo[2], macInfo[3], macInfo[4], uriPath, logBody)
		}
		return &MacHeader{
			Version:          macInfo[0],
			SecretVersion:    macInfo[1],
			TimestampNs:      ts,
			DeviceIdentifier: macInfo[3],
			UriPath:          uriPath,
			Message:          message,
			LogMessage:       logMessage,
			Mac:              macInfo[4],
		}, nil
	}
	return nil, fmt.Errorf("Unsupported MAC version %v", macInfo[0])
}

func (self *MacHeader) RemoveSensitive() *MacHeader {
	macHeader := *self
	if macHeader.Message != "" && macHeader.LogMessage != "" {
		macHeader.Message = ""
	}
	return &macHeader
}

func (self *MacHeader) validate(allowedTimeDiff int, secrets HMacSecrets) error {
	if self.Version == "1" || self.Version == "2" {
		if AbsInt64(time.Now().Unix()-self.Timestamp) > int64(allowedTimeDiff) {
			return fmt.Errorf("Abs(Now - %v) > AllowedDiff(%v)", self.Timestamp, allowedTimeDiff)
		}
	} else if self.Version == "9" || self.Version == "10" {
		if AbsInt64(time.Now().UnixNano()-self.TimestampNs) > (int64(allowedTimeDiff) * int64(time.Second/time.Nanosecond)) {
			return fmt.Errorf("Abs(Now - %v) > AllowedDiff(%v)", self.TimestampNs, allowedTimeDiff)
		}
	} else {
		if AbsInt64(time.Now().Unix()-self.TimestampMs/1000) > int64(allowedTimeDiff) {
			return fmt.Errorf("Abs(NowMs - %v) > AllowedDiff(%v)", self.TimestampMs, allowedTimeDiff)
		}
	}

	if err := IsClientOSSecretVersionMismatch(self.SecretVersion, self.ClientOS); err != nil {
		return err
	}

	if err := IsAppVersionSecretVersionMismatch(self.SecretVersion, self.ClientOS, self.AppVersion, secrets); err != nil {
		return err
	}

	var secret []byte
	if self.ClientOS == OSAndroid {
		secret = []byte(secrets.Android[self.SecretVersion].Secret)
	} else if self.ClientOS == OSIOS {
		secret = []byte(secrets.IOS[self.SecretVersion].Secret)
	}
	if len(secret) == 0 {
		return fmt.Errorf("secret key not found by secret version: %s", self.SecretVersion)
	}

	switch self.Version {
	case "1", "2", "3", "4":
		mac := hmac.New(sha1.New, secret)
		mac.Write([]byte(self.Message))
		expectedMac := mac.Sum(nil)

		headerMac, err := base64.StdEncoding.DecodeString(self.Mac)
		if err != nil {
			return err
		}

		// android should be able to return here
		if hmac.Equal(expectedMac, headerMac) {
			return nil
		}

		if len(headerMac) != 20 {
			return ErrBadHMAC
		}

		if VersionGreaterThan("2.3.3", self.AppVersion) {
			return ErrBadHMAC
		}

		fieldsCount := 4

		if err := checkHMACHiddenInfoForiOS(expectedMac, headerMac, fieldsCount); err != nil {
			return err
		}
	case "5", "6", "7", "8":
		md5Sum := md5.Sum(append(secret, self.Message...))
		md5r := string(md5Sum[:])

		sha1Sum := sha1.Sum([]byte(reverseString(string(secret)) + md5r))

		headerMac, err := base64.StdEncoding.DecodeString(self.Mac)
		if err != nil {
			return err
		}

		expectedMac := sha1Sum[:]

		pass := hmac.Equal(expectedMac, headerMac)
		if pass {
			return nil
		}

		// pass should be true for android
		if !pass && self.ClientOS == OSAndroid {
			return ErrBadHMAC
		}

		if len(headerMac) != 20 {
			return ErrBadHMAC
		}

		fieldsCount := 4
		if err := checkHMACHiddenInfoForiOS(expectedMac, headerMac, fieldsCount); err != nil {
			return err
		}
	case "9", "10":
		if !validateTimestamp(self.ClientOS, self.TimestampNs) {
			return errors.New("Illegal timestamp")
		}
		sha256Sum := sha256.Sum256(append(secret, self.Message...))
		md5r := string(sha256Sum[:])
		sha1Sum := sha1.Sum([]byte(reverseString(string(secret)) + md5r))

		headerMac, err := base64.StdEncoding.DecodeString(self.Mac)
		if err != nil {
			return err
		}
		expectedMac := sha1Sum[:]
		pass := hmac.Equal(expectedMac, headerMac)
		if !pass {
			return ErrBadHMAC
		}
	}

	return nil
}

func validateTimestamp(os string, ts int64) bool {
	v := (ts % 1000000) % 177
	switch os {
	case OSIOS:
		if v == 37 {
			return true
		}
	case OSAndroid:
		if v == 137 {
			return true
		}
	}
	return false
}

func GenerateSHA1Hash(message []byte) string {
	h := sha1.New()
	h.Write(message)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func MaskSensitiveFields(urlPath string, body interface{}) (retBody []byte, err error) {
	retBodyJson := make(map[string]interface{})
	switch v := body.(type) {
	case []byte:
		json.Unmarshal(v, &retBodyJson)
	case url.Values:
		for k := range v {
			retBodyJson[k] = v.Get(k)
		}
	default:
		return nil, fmt.Errorf("Unsupport type: %T", v)
	}

	maskText := "[REMOVED]"
	filterMap := map[string][]string{
		"/verify-password": []string{"value"},
	}
	filters := []string{"password", "old", "new", "thirdparty_access_token"}
	if fields, ok := filterMap[urlPath]; ok {
		filters = append(filters, fields...)
	}
	for i := range filters {
		if v, find := retBodyJson[filters[i]]; find {
			if _, ok := v.(string); ok {
				retBodyJson[filters[i]] = maskText
			}
		}
	}
	retBody, err = json.Marshal(retBodyJson)
	if err != nil {
		return nil, err
	}
	return retBody, nil
}

func checkHMACHiddenInfoForiOS(expectedMac, headerMac []byte, fieldsCount int) error {
	if !hmac.Equal(expectedMac[:16], headerMac[:16]) {
		return ErrBadHMAC
	}

	result := extractHiddenInfo(headerMac, fieldsCount)
	if result[2] > 6 || result[1] == 15 || result[0] == 1 {
		return HiddenError(fmt.Sprintf("reinstall %d copy-paste %d motion %d are not valid in hmac", result[2], result[1], result[0]))
	}

	x := (32 - fieldsCount*3) / 8
	if !bytes.Equal(expectedMac[16:16+x], headerMac[16:16+x]) {
		return ErrBadHMAC
	}

	return nil
}

func extractHiddenInfo(data []byte, fieldsCount int) []uint64 {
	p1 := data[0:4]
	p2 := data[4:8]
	p3 := data[8:12]
	p4 := data[12:16]
	p5 := data[16:20]

	cp, err := rc4.NewCipher(data[:16])
	if err != nil {
		panic(err)
	}
	cp.XORKeyStream(p5, p5)
	p5 = xorBytes(p5, p4, p3, p2, p1)
	br := bitio.NewReader(bytes.NewReader(p5))

	br.ReadBits(byte(32 - fieldsCount*3))

	result := make([]uint64, 3)

	br.ReadBits(3)
	br.ReadBits(1)
	result[0], _ = br.ReadBits(1) // motion
	result[1], _ = br.ReadBits(4) // copy-paste
	result[2], _ = br.ReadBits(3) // reinstall

	for i := 16; i < 20; i++ {
		data[i] = p5[i-16]
	}

	return result
}

func obfuscation(data []byte, fieldsCount int, value map[int]uint64) []byte {
	tmp := make([]byte, len(data))
	copy(tmp, data)
	p1 := tmp[0:4]
	p2 := tmp[4:8]
	p3 := tmp[8:12]
	p4 := tmp[12:16]

	br := bitio.NewReader(bytes.NewReader(tmp[16:20]))
	wbuf := bytes.NewBuffer(make([]byte, 0, 4))
	bw := bitio.NewWriter(wbuf)

	keptBitsCount := byte(32 - 3*fieldsCount)
	keptBits, _ := br.ReadBits(keptBitsCount)

	bw.WriteBits(keptBits, keptBitsCount)
	for i := fieldsCount - 1; i >= 0; i-- {
		bw.WriteBits(value[i], 3)
	}
	bw.Close()

	p5 := xorBytes(wbuf.Bytes(), p1, p2, p3, p4)

	cp, err := rc4.NewCipher(tmp[:16])
	if err != nil {
		panic(err)
	}
	cp.XORKeyStream(p5, p5)

	for i := 16; i < 20; i++ {
		tmp[i] = p5[i-16]
	}

	return tmp
}

func reverseString(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func xorBytes(b []byte, bmore ...[]byte) []byte {
	rv := make([]byte, len(b))

	for i := range b {
		rv[i] = b[i]
		for _, m := range bmore {
			rv[i] = rv[i] ^ m[i]
		}
	}

	return rv
}
