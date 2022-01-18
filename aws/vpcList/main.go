package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	AWSRegion string = "us-east-1"
	ec2Client *ec2.EC2
)

// Upload a file to AWS S3 bucket
func fileCreate(client *ec2.EC2) (resp *ec2.DescribeVpcEndpointsOutput) {

	// params := &s3.PutObjectInput{
	// 	Body:   file,
	// 	Bucket: aws.String(bucketName),
	// 	Key:    aws.String(filename),
	// 	ACL:    aws.String(s3.BucketCannedACLPublicRead),
	// }

	// fmt.Println("Uploading:", filename)
	// resp, err = S3session.PutObject(params)

	// params := &ec2.DescribeVpcEndpointsInput{
	// 	VpcEndpointIds: aws.StringSlice([]string{vpcId}),
	// }
	// var filters = []*ec2.Filter{
	// 	Name:   *aws.String("vpc-id"),
	// 	Values: []*aws.String{"vpc-0e077c6e7930992df"},
	// }

	// params := &ec2.DescribeVpcEndpointsInput{
	// 	Filters: filters,
	// }

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

	if vpcEndpoints != nil {
		fmt.Println(vpcEndpoints)
	}
}
