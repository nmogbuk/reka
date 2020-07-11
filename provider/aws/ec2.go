package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// returns only instance IDs of unprotected ec2 instances
func getInstanceDetails(svc *ec2.Client, output *ec2.DescribeInstancesResponse, region string) ([]*EC2, error) {
	var ec2Instances []*EC2
	for _, reservation := range output.Reservations {
		for _, instance := range reservation.Instances {
			tags := parseTags(instance.Tags)
			// We need the creation-date when parsing Tags for relative defintions
			tags["creation-date"] = (*instance.LaunchTime).String()
			ec2 := NewEC2(*instance.InstanceId)
			ec2.Region = region
			// Get CreationDate by getting LaunchTime of attached Volume
			ec2.CreationDate = *instance.LaunchTime
			ec2.Tags = tags
			ec2.State = getState(*instance.State.Code)
			ec2Instances = append(ec2Instances, &ec2)
		}
	}

	return ec2Instances, nil
}

// GetAllEC2Instances Get all instances
func GetAllEC2Instances(cfg aws.Config, region string) ([]*EC2, error) {
	svc := ec2.New(cfg)
	params := &ec2.DescribeInstancesInput{}

	// Build the request with its input parameters
	req := svc.DescribeInstancesRequest(params)

	// Send the request, and get the response or error back
	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}
	instances, err := getInstanceDetails(svc, resp, region)
	if err != nil {
		return nil, err
	}
	return instances, nil
}

// StopEC2Instances Stop Running Instances
func StopEC2Instances(cfg aws.Config, instances []*EC2) error {
	svc := ec2.New(cfg)
	var instanceIds []string

	for _, instance := range instances {
		if instance.IsActive() {
			instanceIds = append(instanceIds, instance.ID)
		}
	}

	if len(instanceIds) <= 0 {
		return nil
	}

	params := &ec2.StopInstancesInput{
		InstanceIds: instanceIds,
	}

	req := svc.StopInstancesRequest(params)
	resp, err := req.Send(context.Background())
	// TODO Attach error to specific instance where the error occurred if possible
	if err != nil {
		fmt.Println(resp, err)
	}
	return err
}

// StartEC2Instances Start Stopped instances
func StartEC2Instances(cfg aws.Config, instances []*EC2) error {
	svc := ec2.New(cfg)
	var instanceIds []string

	for _, instance := range instances {
		if instance.IsStopped() {
			instanceIds = append(instanceIds, instance.ID)
		}
	}

	if len(instanceIds) <= 0 {
		return nil
	}

	params := &ec2.StartInstancesInput{
		InstanceIds: instanceIds,
	}

	req := svc.StartInstancesRequest(params)
	resp, err := req.Send(context.Background())
	// TODO Attach error to specific instance where the error occurred if possible
	if err != nil {
		fmt.Println(resp, err)
	}
	return err
}

// StartEC2Instances Start Stopped instances
func TerminateEC2Instances(cfg aws.Config, instances []*EC2) error {
	svc := ec2.New(cfg)
	var instanceIds []string

	for _, instance := range instances {
		if instance.IsStopped() || instance.IsActive() {
			instanceIds = append(instanceIds, instance.ID)
		}
	}

	if len(instanceIds) <= 0 {
		return nil
	}

	params := &ec2.TerminateInstancesInput{
		InstanceIds: instanceIds,
	}

	req := svc.TerminateInstancesRequest(params)
	resp, err := req.Send(context.Background())
	// TODO Attach error to specific instance where the error occurred if possible
	if err != nil {
		fmt.Println(resp, err)
	}
	return err
}
