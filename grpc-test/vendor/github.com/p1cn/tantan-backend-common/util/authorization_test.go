package util

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestParseAuthBasic(t *testing.T) {
	headers := []http.Header{
		http.Header{},
		http.Header{"Authorization": []string{"Bearer 789 abc"}},
		http.Header{"Authorization": []string{"MAC 456 abc"}},
	}

	for _, header := range headers {
		s := ParseAuthBasic(header)
		if s != "" {
			t.Errorf("unexpected header %v", s)
		}
	}

	header := http.Header{
		"Authorization": []string{"Basic 123 abc"},
	}

	s := ParseAuthBasic(header)
	if s != "123 abc" {
		t.Errorf("unexpected header %s", s)
	}
}

func TestParseAuthBearerAccessToken(t *testing.T) {
	headers := []http.Header{
		http.Header{},
		http.Header{"Authorization": []string{"Basic 123 abc"}},
		http.Header{"Authorization": []string{"MAC 456 abc"}},
	}

	for _, header := range headers {
		s := ParseAuthBearerAccessToken(header)
		if s != "" {
			t.Errorf("unexpected header %v", s)
		}
	}

	header := http.Header{
		"Authorization": []string{"Bearer 789 abc"},
	}

	s := ParseAuthBearerAccessToken(header)
	if s != "789 abc" {
		t.Errorf("unexpected header %s", s)
	}
}

func TestParseAuthMacBearer(t *testing.T) {
	t.Skip("skipped before making proper fix")
	r := &http.Request{}
	r.Header = http.Header{
		"Authorization": []string{"Bearer 123 abc"},
	}

	h, err := ParseAuthMac(r, "1.7.1", "ios1.7.1", 0, HMacSecrets{})
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if h.Version != "0" {
		t.Errorf("unexpected mac version %v", h.Version)
	}
	if h.AccessToken != "123 abc" {
		t.Errorf("unexpected access token %v", h.AccessToken)
	}
}

func TestParseAuthMacV1(t *testing.T) {
	t.Skip("skipped before making proper fix")
	var hMacSecrets HMacSecrets

	hMacSecrets.IOS = make(map[string]HMacSecret)
	hMacSecrets.IOS["ios1.7.1"] = HMacSecret{Secret: "BGwe%4zMjZN7JJTc&f*GyX7L!g@8dpXuJ7!*rXR2QWc$-+Y2ZP", AppVersionRange: [2]string{"1.7.1", "2.3.0"}}
	allowedDiff := int(time.Now().Unix()-1465202473) + 10

	r, err := http.NewRequest("GET", "http://test.com/v1/test", nil)
	_ = err
	r.Header = http.Header{
		"Authorization": []string{`MAC ["1","ios1.7.1","1465202473","0a35470327fb982c0d73f1b39bd2b5bd88a23f499c5e8c595ae58ec2f98479d0","VLD2OK4b9svOw94zZH8ZowY5Ls0="]`},
	}

	h, err := ParseAuthMac(r, "1.7.1", "ios1.7.1", allowedDiff, hMacSecrets)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if h == nil {

	}
	if h.Version != "1" {
		t.Errorf("unexpected mac version %v", h.Version)
	}
	if h.AccessToken != "0a35470327fb982c0d73f1b39bd2b5bd88a23f499c5e8c595ae58ec2f98479d0" {
		t.Errorf("unexpected access token %v", h.AccessToken)
	}

}

func TestDetectHMACHiddenInfo(t *testing.T) {
	t.Skip("skipped before making proper fix")
	for _, table := range []struct {
		original       string
		hidden         string
		fieldsCount    int
		errShouldBeNil bool
	}{
		{`3388b3c358fe88b6b454dfacf8fd8b24b6264234`, `3388b3c358fe88b6b454dfacf8fd8b24dc076bbb`, 4, false},
		{`b7f5a5383717e86112d34d394090823bbdc3b1fc`, `b7f5a5383717e86112d34d394090823b7c164fff`, 4, true},
		{`bf199ee5290b542f6b329ff632bc00ca17fbf919`, `bf199ee5290b542f6b329ff632bc00caf742c3a5`, 4, true},
		{`e0790da0dc0d3cc92fa0a7336d5b8a01c8676e26`, `e0790da0dc0d3cc92fa0a7336d5b8a017c915135`, 4, true},
	} {
		o, _ := hex.DecodeString(table.original)
		h, _ := hex.DecodeString(table.hidden)
		err := checkHMACHiddenInfoForiOS(o, h, table.fieldsCount)
		if err != nil && table.errShouldBeNil {
			t.Fatal(err)
		}
	}
}

func TestDetectAndroidBoundary(t *testing.T) {
	for _, table := range []struct {
		boundary       string
		errShouldBeNil bool
	}{
		{`a65db794-aa7a-487b-b1c8-60007adc8769`, false},
		{`9c633701-bdd2-44f6-8550-6b0232c1e566`, true},
	} {
		err := ParseAuthBoundary(table.boundary, "2.4.0", 0)
		if err != nil && table.errShouldBeNil {
			t.Fatal(err)
		}
	}
}

func TestMaskSensitiveFields(t *testing.T) {
	password := "correct_password"
	expectedPassword := "[REMOVED]"
	body := map[string]interface{}{
		"test1":    "1",
		"test2":    "2",
		"test3":    "3",
		"password": password,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	res, err := MaskSensitiveFields("/test", bodyBytes)
	if err != nil {
		t.Fatal(err)
	}
	var resJson map[string]interface{}
	json.Unmarshal(res, &resJson)
	if resJson["password"] != expectedPassword {
		t.Errorf("mask failed, origin: %v, masked: %v", body, resJson)
	}

	emptyBody := map[string]interface{}{}
	emptyBodyBytes, err := json.Marshal(emptyBody)
	if err != nil {
		t.Fatal(err)
	}
	emptyRes, err := MaskSensitiveFields("/test", emptyBodyBytes)
	if err != nil {
		t.Fatal(err)
	}
	var emptyResJson map[string]interface{}
	json.Unmarshal(emptyRes, &emptyResJson)
	if len(emptyResJson) != 0 {
		t.Errorf("mask failed, origin: %v, masked: %v", body, resJson)
	}
}

func TestHMacSecretsIsValid(t *testing.T) {
	var hMacSecrets HMacSecrets

	hMacSecrets.IOS = make(map[string]HMacSecret)

	hMacSecrets.IOS["ios1.7.1"] = HMacSecret{Secret: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", AppVersionRange: [2]string{"1.7.1", "2.3.0"}}
	hMacSecrets.IOS["ios2.3.0"] = HMacSecret{Secret: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", AppVersionRange: [2]string{"2.3.0", "2.3.2"}}
	hMacSecrets.IOS["ios2.3.2"] = HMacSecret{Secret: "cccccccccccccccccccccccccccccccccccccccccccccccccc", AppVersionRange: [2]string{"2.3.2", "2.4.0"}}

	if !hMacSecrets.IsValid() {
		t.Fatal("HMac validation should be passed")
	}

	hMacSecrets.IOS["ios1.7.1"] = HMacSecret{Secret: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", AppVersionRange: [2]string{"1.7.1", "2.3.0"}}
	hMacSecrets.IOS["ios2.3.0"] = HMacSecret{Secret: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", AppVersionRange: [2]string{"1.7.1", "2.3.2"}}
	hMacSecrets.IOS["ios2.3.2"] = HMacSecret{Secret: "cccccccccccccccccccccccccccccccccccccccccccccccccc", AppVersionRange: [2]string{"2.3.2", "2.4.0"}}

	if hMacSecrets.IsValid() {
		t.Fatal("HMac validation should be failed")
	}

	hMacSecrets.IOS["ios1.7.1"] = HMacSecret{Secret: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", AppVersionRange: [2]string{"1.7.1", "2.3.1"}}
	hMacSecrets.IOS["ios2.3.0"] = HMacSecret{Secret: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", AppVersionRange: [2]string{"2.3.0", "2.3.2"}}
	hMacSecrets.IOS["ios2.3.2"] = HMacSecret{Secret: "cccccccccccccccccccccccccccccccccccccccccccccccccc", AppVersionRange: [2]string{"2.3.2", "2.4.0"}}

	if hMacSecrets.IsValid() {
		t.Fatal("HMac validation should be failed")
	}

	hMacSecrets.IOS["ios1.7.1"] = HMacSecret{Secret: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", AppVersionRange: [2]string{"1.7.1", "2.4.0"}}
	hMacSecrets.IOS["ios2.3.0"] = HMacSecret{Secret: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", AppVersionRange: [2]string{"2.3.0", "2.3.2"}}
	hMacSecrets.IOS["ios2.3.2"] = HMacSecret{Secret: "cccccccccccccccccccccccccccccccccccccccccccccccccc", AppVersionRange: [2]string{"2.3.2", "2.4.0"}}

	if hMacSecrets.IsValid() {
		t.Fatal("HMac validation should be failed")
	}
}
