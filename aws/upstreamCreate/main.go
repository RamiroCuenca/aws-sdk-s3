package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	AWSRegion string = "us-east-1"
	ec2Client *ec2.EC2
)

type Vpce struct {
	VpcEndpoints []struct {
		CreationTimestamp time.Time `json:"CreationTimestamp"`
		DnsEntries        []struct {
			DnsName      string `json:"DnsName"`
			HostedZoneId string `json:"HostedZoneId"`
		} `json:"DnsEntries"`
		Groups []struct {
			GroupId   string `json:"GroupId"`
			GroupName string `json:"GroupName"`
		} `json:"Groups"`
		NetworkInterfaceIds []string `json:"NetworkInterfaceIds"`
		OwnerId             string   `json:"OwnerId"`
		PolicyDocument      string   `json:"PolicyDocument"`
		PrivateDnsEnabled   bool     `json:"PrivateDnsEnabled"`
		RequesterManaged    bool     `json:"RequesterManaged"`
		ServiceName         string   `json:"ServiceName"`
		State               string   `json:"State"`
		SubnetIds           []string `json:"SubnetIds"`
		VpcEndpointId       string   `json:"VpcEndpointId"`
		VpcEndpointType     string   `json:"VpcEndpointType"`
		VpcId               string   `json:"VpcId"`
	} `json:"VpcEndpoints"`
}

// Upload a file to AWS S3 bucket
func fileCreate(client *ec2.EC2) (resp *ec2.DescribeVpcEndpointsOutput) {
	params := &ec2.DescribeVpcEndpointsInput{}

	resp, err := client.DescribeVpcEndpoints(params)

	if err != nil {
		fmt.Printf("Wasn't able to fetch any VPC endpoint from VPC\n")
		fmt.Printf("Error: %s", err)
		return
	}

	return
}

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	ec2Client := ec2.New(sess)

	vpcEndpoints := fileCreate(ec2Client)
	fmt.Println(vpcEndpoints)
	// fmt.Println([]byte(vpcEndpoints.String()))

	var endpoint *Vpce
	// json.NewDecoder([]byte(vpcEndpoints.String())).Decode(&endpoint)
	err = json.Unmarshal([]byte(vpcEndpoints.String()), &endpoint)
	if err != nil {
		fmt.Println("error unmarshaling")
		fmt.Println("error: ", err)
	}
	if vpcEndpoints != nil {
		fmt.Println(endpoint)
	}
}
