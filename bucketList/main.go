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

// List existing S3 buckets.
// Returns a JSON object similar to the following:
/*
{
	Buckets: [{
		CreationDate: 2022-01-17 15:13:39 +0000 UTC,
		Name: "bucket-unique-name"
	  }],
	Owner: {
	  DisplayName: "my-account-name",
	  ID: "31easadasdasdadsdagjkukkuykadgsd124aa62cb93e8845e"
	}
}
*/
func bucketList() (resp *s3.ListBucketsOutput) {
	params := &s3.ListBucketsInput{}

	// S3session.ListBuckets() Returns a list of all buckets owned by the authenticated sender of the request
	resp, err := S3session.ListBuckets(params) // Leave it empty so that it retrieve all
	if err != nil {
		fmt.Println("Wasn't able to find any bucket at S3")
		fmt.Printf("Error: %s", err)
		return
	}

	return
}

func main() {
	buckets := bucketList()
	if buckets != nil {
		fmt.Println(buckets)
	}
}
