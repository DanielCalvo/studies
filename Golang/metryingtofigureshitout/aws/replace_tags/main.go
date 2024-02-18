package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

func main() {
	// Specify your AWS region and S3 bucket name
	region := "eu-west-1"
	bucket := "dcalvo-testing-bucket"

	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create a new S3 client
	svc := s3.New(sess)

	start := time.Now()

	// Call GetBucketTagging to retrieve tags for the bucket
	resp, err := svc.GetBucketTaggingWithContext(context.TODO(), &s3.GetBucketTaggingInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		// If the bucket does not have any tags or if GetBucketTagging fails for other reasons,
		// you may receive an error such as NoSuchTagSet. You can handle it accordingly.
		fmt.Println("Error getting tags for bucket", bucket, ":", err)
		return
	}

	// Modify the existing tags or add a new one
	var newTags []*s3.Tag
	for _, tag := range resp.TagSet {
		// Replace the value of tag1 with "banana=123"
		if *tag.Key == "tag1" {
			newTags = append(newTags, &s3.Tag{
				Key:   aws.String("banana"),
				Value: aws.String("123"),
			})
		} else {
			// Keep other tags intact
			newTags = append(newTags, tag)
		}
	}

	// Update the tags for the S3 bucket
	_, err = svc.PutBucketTaggingWithContext(context.TODO(), &s3.PutBucketTaggingInput{
		Bucket: aws.String(bucket),
		Tagging: &s3.Tagging{
			TagSet: newTags,
		},
	})
	if err != nil {
		fmt.Println("Error updating tags for bucket", bucket, ":", err)
		return
	}
	duration := time.Since(start)

	fmt.Println("Tags updated successfully for bucket:", bucket)
	fmt.Println(duration)

}
