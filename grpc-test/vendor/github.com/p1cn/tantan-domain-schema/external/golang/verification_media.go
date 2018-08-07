package external

type VerificationMedia struct {
	Url       string  `json:"url"`
	MediaType string  `json:"mediaType"`
	Size      []int   `json:"size"`
	Duration  float64 `json:"duration"`
}

func (self VerificationMedia) isImage() bool {
	return self.MediaType == "image/jpeg"
}

func (self VerificationMedia) IsValid() bool {
	if self.Url == "" ||
		!self.isImage() {
		return false
	}
	return true
}
