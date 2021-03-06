package aws

import (
	log "github.com/sirupsen/logrus"

	"github.com/mensaah/reka/config"
	"github.com/mensaah/reka/resource"
)

// Manages EC2 instances on the AWS.
// EC2 resources support stopping/resuming and terminating instances.

var ec2Manager resource.Manager

const (
	// Name of resource
	ec2Name = "ec2"
	// LongName descriptive name for resource
	ec2LongName = "Elastic Compute Cloud"
)

var ec2Logger *log.Entry

func newEC2Manager(cfg *config.Config, logPath string) resource.Manager {
	ec2Logger = config.GetLogger(ec2Name, logPath)

	ec2Manager = resource.Manager{
		Name:     ec2Name,
		LongName: ec2LongName,
		Config:   cfg,
		Logger:   logger,
		GetAll: func() ([]*resource.Resource, error) {
			return GetAllEC2Instances(*cfg.Aws)
		},
		Destroy: func(resources []*resource.Resource) error {
			return TerminateEC2Instances(*cfg.Aws, resources)
		},
		Stop: func(resources []*resource.Resource) error {
			return StopEC2Instances(*cfg.Aws, resources)
		},
		Resume: func(resources []*resource.Resource) error {
			return ResumeEC2Instances(*cfg.Aws, resources)
		},
	}
	return ec2Manager
}
