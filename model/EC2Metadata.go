package model

type EC2Metadata struct {
	AccountID               string
	Architecture            string
	AvailabilityZone        string
	BillingProducts         []string
	DevpayProductCodes      []string
	MarketplaceProductCodes []string
	ImageID                 string
	InstanceID              string
	InstanceType            string
	KernelID                string
	PendingTime             string
	PrivateIP               string
	RamdiskID               string
	Region                  string
	Version                 string
}
