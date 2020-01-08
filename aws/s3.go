package aws

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// CreateBucket - Create S3 bucket
func CreateBucket(awsRegion, bucketName string) {
	svc := s3.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(awsRegion),
		},
	}
	result, err := svc.CreateBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// UploadBucket - Upload data to S3 bucket
func UploadBucket(awsRegion, bucketName, fileUpload string) {
	svc := s3.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader(isFile(fileUpload))),
		Bucket: aws.String(bucketName),
		Key:    aws.String(getFilename(fileUpload)),
	}
	result, err := svc.PutObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// ListBuckets - List of all buckets
func ListBuckets(awsRegion string) {
	svc := s3.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// DeleteBucket - Delete bucket
func DeleteBucket(awsRegion, bucketName string) {
	svc := s3.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}
	result, err := svc.DeleteBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DeleteObjects - Delete object from s3 bucket
func DeleteObjects(awsRegion, bucketName, objectKey string) {
	svc := s3.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}
	result, err := svc.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

// ListObjects - List all objects in a bucket
func ListObjects(awsRegion, bucketName string) {
	svc := s3.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucketName),
		MaxKeys: aws.Int64(2),
	}
	result, err := svc.ListObjectsV2(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}
