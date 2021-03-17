package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/awserr"
    "fmt"
    "testing"
    "os"
)

const ExpectedNumberOfDeployedVPCs int = 1

func TestMain(m *testing.M) {
    setUp()
    retCode := m.Run()
    tearDown()
    os.Exit(retCode)
}

func setUp() {
    
    fmt.Println("Setting-Up Test Fixture")
}


func tearDown() {
    fmt.Println("Tearing-Down Test Fixture")
}


func TestTotalNumberOfDeployedVPCs(t *testing.T) {

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-1")},
    )

	svc := ec2.New(sess)
	
	input := &ec2.DescribeVpcsInput{}

	vpcOutput, err := svc.DescribeVpcs(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
    }

    if len(vpcOutput.Vpcs) != ExpectedNumberOfDeployedVPCs {
       t.Errorf("Number of deployed VPCs was incorrect, returned: %d, expected: %d.", len(vpcOutput.Vpcs), ExpectedNumberOfDeployedVPCs)
    }
}

