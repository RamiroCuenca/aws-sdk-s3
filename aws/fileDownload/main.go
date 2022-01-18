package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	// S3 session
	S3session     *s3.S3
	AWSRegion     string = "us-east-1"
	donwloadsPath string = "script"
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

// Download a file from S3
func getFile(filename string) {
	fmt.Println("Downloading: ", filename)

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	}

	resp, err := S3session.GetObject(params)

	if err != nil {
		fmt.Println("Wasn't able to fetch the file from the S3 bucket")
		fmt.Printf("Error: %s", err)
		return
	}

	body, err := io.ReadAll(resp.Body)

	// If it doesn't exist, create a directory where to store the downloaded file
	os.Mkdir(donwloadsPath, 0700)

	// Save or write the downloaded file into the desired directory
	err = os.WriteFile(donwloadsPath+"/"+filename, body, 0644)
	if err != nil {
		if err != nil {
			fmt.Println("Wasn't able to write the file at desired directory")
			fmt.Printf("Error: %s", err)
			return
		}
	}
}

func main() {
	getFile("upstream.txt")
}
