package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func main() {

	bucketName := "s3-cost-calc-test-bucket"
	AWSregion := "eu-west-1"

	startTime, err := time.Parse("2006-Jan-02", "2023-Jan-27")
	endTime, err := time.Parse("2006-Jan-02", "2023-Jan-28")

	s, err := session.NewSession(&aws.Config{Region: aws.String(AWSregion)})
	if err != nil {
		log.Fatal(err)
	}

	svc := cloudwatch.New(s, &aws.Config{Region: aws.String(AWSregion)})

	params := &cloudwatch.GetMetricStatisticsInput{
		EndTime:    aws.Time(endTime),
		MetricName: aws.String("NumberOfObjects"),
		Namespace:  aws.String("AWS/S3"),
		Period:     aws.Int64(86400),
		StartTime:  aws.Time(startTime),

		Statistics: []*string{
			aws.String("Average"),
		},
		Dimensions: []*cloudwatch.Dimension{
			{Name: aws.String("BucketName"), Value: aws.String(bucketName)},
			{Name: aws.String("StorageType"), Value: aws.String("AllStorageTypes")},
		},
	}

	resp, err := svc.GetMetricStatistics(params)
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.Datapoints) > 1 {
		log.Fatal("Found more than one bucket matching cloudfront metrics. The program doesn't know how to handle this, exiting")
	}
	//Hey what if you find no metrics? Maybe you just created the bucket...

	fmt.Println("Objects in bucket:", *resp.Datapoints[0].Average)

	//Now list objects!
	/*
		Data you want
		- Path on the bucket
		- Size
		- Age
		- Storage class
	*/

	//This is what I'm looking for, very noice!
	//aws s3api list-objects-v2 --bucket mdscanner-bucket

}
