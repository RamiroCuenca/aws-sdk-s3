
# Golang - Manipulating buckets from S3 using the AWS SDK

In this project I develop a set of services that allow us to manipulate AWS S3 buckets & files. I ahieve it by using the AWS SDK library for Go.

I developed functions for: 

1. Create a new S3 bucket.
2. List existing S3 buckets.
3. Create or insert a file inside an existing S3 bucket.
4. List existing files inside an existing S3 bucket.
5. Delete a file from an existing S3 bucket.
6. Get/fetch and download a file from an existing S3 bucket.

The objective of this project was to learn how to manipulate either S3 Buckets or S3 Files from Go using the AWS SDK.

## Run Locally

Clone the project.

```bash
  git clone https://github.com/RamiroCuenca/go-aws-s3.git
```

Go to the project directory.

```bash
  cd go-aws-s3
```

Before running the project... we need to create manually an S3 bucket with the name we set at main.go for the variable _bucketName_. 
Also, we need to modify the _main function_ so that it calls the function that we want.

Run the main.go file.

```bash
  go run main.go
```

## Tech Stack

**Language:**
- Go

**Libraries:**
- Amazon Web Services SDK

**External Services:**
- Amazon Web Services S3

  
## Author

- [@RamiroCuenca](https://www.linkedin.com/in/ramiro-cuenca-salinas-749a2020a/)
