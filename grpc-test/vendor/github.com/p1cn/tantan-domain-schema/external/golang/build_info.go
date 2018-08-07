package external

const (
	BuildInfoType = "buildInfo"
)

type BuildInfo struct {
	UserId   string
	JsonData string
	Hash     string
	Status   string
}
