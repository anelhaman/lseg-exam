package model

type EC2Info struct {
	Name             string
	InstanceID       string
	InstanceType     string
	PrivateIP        string
	PublicIP         string
	AvailabilityZone string
}
