package model

type EC2Metadata struct {
	InstanceID       string
	InstanceType     string
	PrivateIP        string
	AvailabilityZone string
	Region           string
	ImageID          string
	AccountID        string
	Hostname         string
	PublicHostname   string
	MacAddress       string
	IAMRoleName      string
	UserData         string
}
