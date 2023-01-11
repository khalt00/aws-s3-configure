package main

import (
	"fmt"
	"os"
	"plugin"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/plugincreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Uploads a file to S3 given a bucket and object key. Also takes a duration
// value to terminate the update if it doesn't complete within that time.
//
// The AWS Region needs to be provided in the AWS shared config or on the
// environment variable as `AWS_REGION`. Credentials also must be provided
// Will default to shared config file, but can load from environment if provided.
//
// Usage:
//   # Upload myfile.txt to myBucket/myKey. Must complete within 10 minutes or will fail
//   go run withContext.go -b mybucket -k myKey -d 10m < myfile.txt

func main() {
	// var bucket, key string
	// var timeout time.Duration

	// flag.StringVar(&bucket, "b", "", "Bucket name.")
	// flag.StringVar(&key, "k", "", "Object key name.")
	// flag.DurationVar(&timeout, "d", 0, "Upload timeout.")
	// flag.Parse()

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.

	creds, err := plugin.Open("something.so")
	fmt.Print(creds)
	if err != nil {
		fmt.Printf("Oh wow, sai rui. plugin.Open")
	}

	mycreds, err := plugincreds.NewCredentials(creds)
	if err != nil {
		fmt.Printf("Oh wow, sai tiep roi. plugincreds.NewCredentials")
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: mycreds,
	}))

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	//svc := s3.New(sess)

	uploader := s3manager.NewUploader(sess)
	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.

	// Ensure the context is canceled to prevent leaking.
	// See context package for more information, https://golang.org/pkg/context/

	filename := "New Text Document (3).txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("mybucket"),
		Key:    aws.String("mykey"),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))

}
