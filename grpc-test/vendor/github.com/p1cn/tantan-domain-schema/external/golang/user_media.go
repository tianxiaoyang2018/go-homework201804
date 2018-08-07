package external

type Parameter struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type UserMedia struct {
	Url         string                `json:"url"`
	MediaType   string                `json:"mediaType"`
	Size        []int                 `json:"size"`
	Duration    float64               `json:"duration"`
	Parameters  []Parameter           `json:"parameters"`
	Attachments []UserMediaAttachment `json:"attachments"`
	Status      string                `json:"-"`
}

type UserMediaAttachment struct {
	Url        string      `json:"url"`
	MediaType  string      `json:"mediaType"`
	Size       []int       `json:"size"`
	Duration   float64     `json:"duration"`
	Parameters []Parameter `json:"parameters"`
}

func (self UserMedia) IsImage() bool {
	return self.MediaType == "image/jpeg" &&
		len(self.Attachments) == 0
}

func (self UserMedia) IsAudio() bool {
	return self.MediaType == "audio/mp4" &&
		len(self.Attachments) == 0
}

func (self UserMedia) IsVideo() bool {
	return self.MediaType == "image/jpeg" &&
		len(self.Attachments) == 1 &&
		self.Attachments[0].MediaType == "video/mp4"
}
