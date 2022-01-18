package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	// S3 session
	S3session *s3.S3
	AWSRegion string = "us-east-1"
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

// Upload a file to AWS S3 bucket
func fileCreate(filename string) (resp *s3.PutObjectOutput) {
	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("There was an error generating %s. Error: %v", filename, err)
	}

	// Write the file
	data := []byte(`http {
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
	`)

	os.WriteFile(filename, data, 0644)

	params := &s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	}

	fmt.Println("Uploading:", filename)
	resp, err = S3session.PutObject(params)

	if err != nil {
		fmt.Printf("Wasn't able to upload '%s' file to the S3 bucket\n", filename)
		fmt.Printf("Error: %s", err)
		return
	}

	return
}

func main() {
	fileCreate("upstream.txt")
}
