package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	// S3 session
	S3session *s3.S3
	// EC2 session
	EC2session *ec2.EC2
	AWSRegion  string = "us-east-1"
)

// Must exist before executing this file.
// Should have:
// - Versioning enabled.
// - KMS Encryption enabled.
const bucketName = "rsalinas-nginx-upstream"

func init() {
	// Start a new session
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})
	if err != nil {
		fmt.Printf("Couldn't initialize new session. Error: %v", err)
		return
	}

	// Initialize S3 session - It will use our default credencials from .aws/credentials
	S3session = s3.New(sess)

	// Initialize EC2 session - It will use our default credencials from .aws/credentials
	EC2session = ec2.New(sess)
}

// Generate upstream config file & upload it to AWS S3 bucket
func fileCreate(filename string) (s3output *s3.PutObjectOutput, err error) {

	ec2input := &ec2.DescribeVpcEndpointsInput{}

	ec2output, err := EC2session.DescribeVpcEndpoints(ec2input)

	if err != nil {
		fmt.Printf("Couldn't fetch any VPC endpoint from VPC\n - Error: %s", err)
		return
	}

	// Create upstream config file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Couldn't generate %s. - Error: %v\n", filename, err)
		return
	}

	// Write/update the file
	fileContent := []byte(fmt.Sprintf(
		`http {
			upstream ourproject {
				server 127.0.0.1:8000;
				server 127.0.0.1:8001;
				server 127.0.0.1:8002;
				server 127.0.0.1:8003;
			}
		
			server {
				listen 80;
				server_name www.domain.com;
				location / {
					proxy_pass http://ourproject;
				}
			}
		}

		VPCe: %v
	`, ec2output.String(),
	))

	os.WriteFile(filename, fileContent, 0644)

	// Upload the file to AWS S3
	params := &s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	}

	fmt.Println("Uploading:", filename)
	s3output, err = S3session.PutObject(params)

	if err != nil {
		fmt.Printf("Couldn't upload '%s' file to the S3 bucket - Error: %s\n", filename, err)
		return
	}

	return
}

func main() {
	fileCreate("upstream.txt")
}
