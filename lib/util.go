package ecs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ec2Client *ec2.EC2
)

func init() {
	initAWSClients("")
}

func initAWSClients(roleArn string) {

	var creds *credentials.Credentials
	if roleArn != "" {
		creds = stscreds.NewCredentials(sess, roleArn, func(arp *stscreds.AssumeRoleProvider) {})
	}

	awsConfig := &aws.Config{Credentials: creds}

	ec2Client = ec2.New(sess, awsConfig)
}
