package external

type AutoUpdate struct {
	HasUpdate  bool   `json:"hasUpdate"`
	AppVersion string `json:"appVersion"`
	URL        string `json:"url"`
	MD5        string `json:"md5"`
	ChangeLog  string `json:"changelog"`
	GotoMarket bool   `json:"goToMarket"`
	ForceDLAPK bool   `json:"forceDownloadApk"`
}
