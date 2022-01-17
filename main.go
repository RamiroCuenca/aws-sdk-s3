package main

import (
	"fmt"
	"io"
	"os"
	"strings"

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
func listBuckets() (resp *s3.ListBucketsOutput) {
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

// Create a new bucket at AWS S3
func createBucket() (resp *s3.CreateBucketOutput) {
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

// Upload a file to AWS S3 bucket
func uploadFile(filename string) (resp *s3.PutObjectOutput) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Wasn't able to open/read '%s' file\n", filename)
		fmt.Printf("Error: %s", err)
		return
	}

	params := &s3.PutObjectInput{
		Body:   f,
		Bucket: aws.String(bucketName),
		Key:    aws.String(strings.Split(filename, "/")[1]),
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

// Get or fetch all files from bucket
func listFiles() (resp *s3.ListObjectsV2Output) {
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}

	resp, err := S3session.ListObjectsV2(params)

	if err != nil {
		fmt.Println("Wasn't able to fetch any file from the S3 bucket")
		fmt.Printf("Error: %s", err)
		return
	}

	return resp
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
	// 1. listBuckets()
	// buckets := listBuckets()
	// if buckets != nil {
	// 	fmt.Println(buckets)
	// }

	// 2. createBucket()
	// fmt.Println(createBucket())

	// 3. uploadFile()
	// folder := "images"

	// files, _ := ioutil.ReadDir(folder)
	// fmt.Println(files)
	// for _, file := range files {
	// 	if file.IsDir() {
	// 		continue
	// 	} else {
	// 		uploadFile(folder + "/" + file.Name())
	// 	}
	// }

	// 4. listFiles()
	// fmt.Println(listFiles())

	// 5. getFile()
	// getFile("aws.jpg")
	// getFile("golang.jpg")
	// getFile("nextjs.jpg")
	// getFile("angular.jpg")

	// 6. deleteFile()
	deleteFile("angular.jpg")
}
