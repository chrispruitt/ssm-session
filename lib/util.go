package ecs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var (
	sess      *session.Session
	ec2Client *ec2.EC2
	ssmClient *ssm.SSM
)

func init() {
	Init("")
}

func Init(profile string) {
	opts := session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}
	if profile != "" {
		opts.Profile = profile
	}
	sess = session.Must(session.NewSessionWithOptions(opts))
	initAWSClients("")
}

func initAWSClients(roleArn string) {
	var creds *credentials.Credentials
	if roleArn != "" {
		creds = stscreds.NewCredentials(sess, roleArn, func(arp *stscreds.AssumeRoleProvider) {})
	}

	awsConfig := &aws.Config{Credentials: creds}

	ec2Client = ec2.New(sess, awsConfig)
	ssmClient = ssm.New(sess, awsConfig)
}
