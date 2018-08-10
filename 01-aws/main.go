  package main

  import (
  	"context"
  	"flag"
  	"fmt"
  	"os"
  	"time"

  	"github.com/aws/aws-sdk-go/aws"
  	"github.com/aws/aws-sdk-go/aws/awserr"
  	"github.com/aws/aws-sdk-go/aws/request"
  	"github.com/aws/aws-sdk-go/aws/session"
  	"github.com/aws/aws-sdk-go/service/s3"
  )

  // Uploads a file to S3 given a bucket and object key. 

  // The AWS Region needs to be provided in the AWS shared config or on the
  // environment variable as `AWS_REGION`. 

  // Usage: 
  //   go run main.go -b mybucket -k myKey -d 10m < myfile.txt

  func main() {
  	var bucket, key string
  	var timeout time.Duration

  	flag.StringVar(&bucket, "b", "", "Bucket name.")
  	flag.StringVar(&key, "k", "", "Object key name.")
  	flag.DurationVar(&timeout, "d", 0, "Upload timeout.")
  	flag.Parse()

  	sess := session.Must(session.NewSession())

  	svc := s3.New(sess)

  	ctx := context.Background()
  	var cancelFn func()
  	if timeout > 0 {
  		ctx, cancelFn = context.WithTimeout(ctx, timeout)
  	}

  	defer cancelFn()

  	_, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
  		Bucket: aws.String(bucket),
  		Key:    aws.String(key),
  		Body:   os.Stdin,
  	})

  	if err != nil {
  		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
  			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
  		} else {
  			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
  		}
  		os.Exit(1)
  	}

  	fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)
  }