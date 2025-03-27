package service

import (
	"context"
	"lseg-exam/model"

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
		InstanceID:       doc.InstanceID,
		InstanceType:     doc.InstanceType,
		PrivateIP:        doc.PrivateIP,
		AvailabilityZone: doc.AvailabilityZone,
		Region:           doc.Region,
		ImageID:          doc.ImageID,
		AccountID:        doc.AccountID,
	}, nil
}
