package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	// S3 session
	S3session     *s3.S3
	AWSRegion     string = "us-east-1"
	donwloadsPath string = "downloads"
)

const bucketName = "ramiro-test-bucket"

func init() {
	// Initialize S3 session
	// METHODS:
	// s3.New()
	// Is a method that creates a new instance of a S3 client with a session
	//
	// session.Must()
	// Is a method that validates that we send all required values
	//
	// session.NewSession()
	// Returns a new Session created from SDK defaults, config files, environment, and user provided config files
	//
	// aws.String()
	// Returns a pointer to the string value passed in
	S3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(AWSRegion),
		// Credentials: ,
	})))
}

// Create a new bucket at AWS S3
func bucketCreate() (resp *s3.CreateBucketOutput) {
	params := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		// ACL: aws.String(s3.BucketCannedACLPublicRead),
		ACL: aws.String(s3.BucketCannedACLPrivate),
		// CreateBucketConfiguration: &s3.CreateBucketConfiguration{
		// 	// LocationConstraint: aws.String(AWSRegion),
		// 	// If we do not specify it, it will create it at Virginia (us-east-1)
		// },
	}

	resp, err := S3session.CreateBucket(params)
	if err != nil {
		fmt.Printf("Wasn't able to create '%s' bucket at S3, the reason might be that there already exist a bucket with provided name\n", bucketName)
		fmt.Printf("Error: %s", err)
		return
	}

	return
}

func main() {
	bucketCreate()
}
