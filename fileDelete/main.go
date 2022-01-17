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

// Deletes an existing file from a S3 bucket
func deleteFile(filename string) (resp *s3.DeleteObjectOutput) {
	fmt.Println("Deleting: ", filename)

	params := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	}

	resp, err := S3session.DeleteObject(params)

	if err != nil {
		if err != nil {
			if err != nil {
				fmt.Println("Wasn't able to delete the file from the S3 bucket")
				fmt.Printf("Error: %s", err)
				return
			}
		}
	}

	return resp
}

func main() {
	deleteFile("angular.jpg")
}
