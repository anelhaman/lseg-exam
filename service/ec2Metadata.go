package service

import (
	"context"
	"lseg-exam/model"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
)

type EC2MetaFetcher struct {
	client *imds.Client
}

func NewEC2MetaFetcher() (*EC2MetaFetcher, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	client := imds.NewFromConfig(cfg)

	return &EC2MetaFetcher{client: client}, nil
}

func (f *EC2MetaFetcher) GetMetadata() (*model.EC2Metadata, error) {
	doc, err := f.client.GetInstanceIdentityDocument(context.TODO(), &imds.GetInstanceIdentityDocumentInput{})
	if err != nil {
		return nil, err
	}

	return &model.EC2Metadata{
		AccountID:               toPtr(doc.AccountID),
		Architecture:            toPtr(doc.Architecture),
		AvailabilityZone:        toPtr(doc.AvailabilityZone),
		BillingProducts:         slicePtr(doc.BillingProducts),
		DevpayProductCodes:      slicePtr(doc.DevpayProductCodes),
		MarketplaceProductCodes: slicePtr(doc.MarketplaceProductCodes),
		ImageID:                 toPtr(doc.ImageID),
		InstanceID:              toPtr(doc.InstanceID),
		InstanceType:            toPtr(doc.InstanceType),
		KernelID:                toPtr(doc.KernelID),
		PendingTime:             doc.PendingTime.Format(time.RFC3339),
		PrivateIP:               toPtr(doc.PrivateIP),
		RamdiskID:               toPtr(doc.RamdiskID),
		Region:                  toPtr(doc.Region),
		Version:                 toPtr(doc.Version),
	}, nil
}

func toPtr(v string) string {
	if v == "" {
		return "null"
	}
	return v
}

func slicePtr(s []string) []string {
	if len(s) == 0 {
		return []string{}
	}
	return s
}
