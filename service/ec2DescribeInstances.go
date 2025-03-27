package service

import (
	"context"
	"lseg-exam/model"
	"lseg-exam/util"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2 struct {
	client *ec2.Client
}

func NewEC2() (*EC2, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)

	return &EC2{client: client}, nil
}

func (e *EC2) DescribeInstances() ([]model.EC2Info, error) {
	out, err := e.client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		return nil, err
	}

	var result []model.EC2Info

	for _, res := range out.Reservations {
		for _, inst := range res.Instances {

			info := model.EC2Info{

				InstanceID:       util.PointerToString(inst.InstanceId),
				InstanceType:     string(inst.InstanceType),
				PrivateIP:        util.PointerToString(inst.PrivateIpAddress),
				PublicIP:         util.PointerToString(inst.PublicIpAddress),
				AvailabilityZone: util.PointerToString(inst.Placement.AvailabilityZone),
			}

			for _, tag := range inst.Tags {
				if util.PointerToString(tag.Key) == "Name" {
					info.Name = util.PointerToString(tag.Value)
					break
				}
			}
			result = append(result, info)
		}
	}
	return result, nil
}
