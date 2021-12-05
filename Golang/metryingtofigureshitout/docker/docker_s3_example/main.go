package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

type Fileyaml struct {
	One   string `yaml:"one"`
	Two   string `yaml:"two"`
	Three string `yaml:"three"`
}

func main() {
	fmt.Println("hello world")

	//downloader := s3manager.NewDownloader(sess)
	//fmt.Println(downloader)
	fmt.Println(os.Getenv("AWS_ACCESS_KEY_ID"))
	fmt.Println(os.Getenv("AWS_SECRET_ACCESS_KEY"))
	fmt.Println(os.Getenv("AWS_DEFAULT_REGION"))
	//s3 logic goes here
	//error out to the webserver if you can't list the bucket contents (just printf to the webserver)
	S3session, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-1")})
	fmt.Println(S3session, err)

	myS3Client := s3.New(S3session)

	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(os.Getenv("MY_SAMPLE_BUCKET")),
		MaxKeys: aws.Int64(3),
	}

	result, err := myS3Client.ListObjectsV2(input)
	fmt.Println(result, err)

}
