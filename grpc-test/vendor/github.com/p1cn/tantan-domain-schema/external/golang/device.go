package external

const (
	DeviceType             = "device"
	AppUIVersionDefault    = "1.0.0"
	AppUIVersionSummer2017 = "2.0.0"
	AppUIVersionSummer2018 = "3.0.0"
)

type DevicePushInfo struct {
	Service string `json:"service"`
	Token   string `json:"token"`
}

type DevicePushNotification struct {
	Service string `json:"service"`
	Token   string `json:"token"`
}

type DeviceOs struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type DeviceIdentifier struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type Device struct {
	Id                string             `json:"id"`
	Owner             IdType             `json:"owner"`
	DeviceIdentifier  string             `json:"identifier"`
	DeviceIdentifiers []DeviceIdentifier `json:"identifiers"`
	AppVersion        string             `json:"appVersion"`
	AppBuild          string             `json:"appBuild"`
	OperatingSystem   DeviceOs           `json:"operatingSystem"`
	DeviceName        string             `json:"name"`
	DeviceMacAddress  string             `json:"macAddress"`
	IpAddress         string             `json:"-"`
	Language          string             `json:"language"`
	Locale            string             `json:"locale"`
	Ringtone          string             `json:"default"`
	CreatedTime       Iso8601Time        `json:"createdTime"`
	Type              string             `json:"type"`
	AppUIVersion      string             `json:"appUIVersion"`

	PushNotifications       DevicePushInfo           `json:"pushNotifications"`
	DevicePushNotifications []DevicePushNotification `json:"devicePushNotifications"`
}

func (self Device) IsValid() bool {
	for _, v := range []string{
		self.AppVersion,
		self.AppBuild,
		self.OperatingSystem.Name,
		self.OperatingSystem.Version,
		self.DeviceName,
		self.Language,
		self.Locale,
	} {
		if v == "" {
			return false
		}
	}
	return true
}

// func (self Device) ShouldTrimFieldLength() (trim bool, trimFunc func(*Device)) {
// 	lengthMap := map[int][]string{
// 		50: []string{
// 			self.PushNotifications.Service,
// 			self.AppVersion,
// 			self.AppBuild,
// 			self.OperatingSystem.Name,
// 			self.OperatingSystem.Version,
// 			self.Language,
// 			self.Locale,
// 		},
// 		100: []string{self.DeviceName},
// 		150: []string{self.PushNotifications.Token},
// 	}
// 	trim = false
// 	trimFunc = nil
// 	for length, fields := range lengthMap {
// 		for _, field := range fields {
// 			if len([]rune(field)) > length {
// 				trim = true
// 				trimFunc = func(d *Device) {
// 					d.PushNotifications.Service = util.MultiByteSubstr(d.PushNotifications.Service, 0, 50)
// 					d.AppVersion = util.MultiByteSubstr(d.AppVersion, 0, 50)
// 					d.AppBuild = util.MultiByteSubstr(d.AppBuild, 0, 50)
// 					d.OperatingSystem.Name = util.MultiByteSubstr(d.OperatingSystem.Name, 0, 50)
// 					d.OperatingSystem.Version = util.MultiByteSubstr(d.OperatingSystem.Version, 0, 50)
// 					d.Language = util.MultiByteSubstr(d.Language, 0, 50)
// 					d.Locale = util.MultiByteSubstr(d.Locale, 0, 50)
// 					d.DeviceName = util.MultiByteSubstr(d.DeviceName, 0, 100)
// 					d.PushNotifications.Token = util.MultiByteSubstr(d.PushNotifications.Token, 0, 150)
// 				}
// 				return
// 			}
// 		}
// 	}
// 	return
// }

func (self Device) Identifiers() []string {
	if len(self.DeviceIdentifiers) > 0 {
		ids := make([]string, len(self.DeviceIdentifiers))
		for i, id := range self.DeviceIdentifiers {
			ids[i] = id.Token
		}
		return ids
	}
	if self.DeviceIdentifier == "" {
		return nil
	}
	return []string{self.DeviceIdentifier}
}
