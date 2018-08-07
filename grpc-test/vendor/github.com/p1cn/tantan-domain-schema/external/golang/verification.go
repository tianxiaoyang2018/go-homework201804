package external

const (
	VerificationStatusPending     = "pending"
	VerificationStatusVerified    = "verified"
	VerificationStatusRejected    = "rejected"
	VerificationStatusDefault     = "default"
	VerificationStatusPrePending  = "prePending"
	VerificationStatusPreVerified = "preVerified"
)

const (
	VerifiedTypeStudent = "student"
	VerifiedTypeDefault = "default"
)

const (
	SearchPriorityStudent = "student"
)

type VerificationStudies struct {
	School           IdType              `json:"school"`
	SchoolName       string              `json:"-"`
	Pictures         []VerificationMedia `json:"pictures"`
	StartTime        *Date               `json:"startTime,omitempty"`
	EndTime          *Date               `json:"endTime,omitempty"`
	Status           string              `json:"status"`
	RejectionReasons []string            `json:"rejectionReasons"`
	HideSchoolName   bool                `json:"hideSchoolName"`
}

type Verification struct {
	UserId       string              `json:"-"`
	VerifiedType string              `json:"-"`
	Studies      VerificationStudies `json:"studies"`
}

const (
	VerificationWithPicturesType = iota
	VerificationWithoutPicturesType
	PreVerificationType
	InvalidType
)

func (self Verification) isStartTimeValid() bool {
	return self.Studies.StartTime != nil && !self.Studies.StartTime.IsZero()
}

func (self Verification) GetRequestType() int {
	if self.Studies.School.Id == "" ||
		self.Studies.School.Type != "school" {
		return InvalidType
	}

	if self.Studies.Pictures != nil && len(self.Studies.Pictures) == 2 {
		for _, picture := range self.Studies.Pictures {
			if !picture.IsValid() {
				return InvalidType
			}
		}

		if self.isStartTimeValid() {
			return VerificationWithPicturesType
		}
		return InvalidType
	}

	if self.Studies.Pictures == nil || len(self.Studies.Pictures) == 0 {
		if self.isStartTimeValid() {
			return VerificationWithoutPicturesType
		}
		return PreVerificationType
	}

	return InvalidType
}
