package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type Banana struct {
	Name     string
	Weight   int
	Forscale bool
}

func main() {

	S3_REGION := "eu-west-1"
	S3_BUCKET := "dcalvo-testing-bucket"

	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}

	a := Banana{
		Name:     "McBanana",
		Weight:   2,
		Forscale: true,
	}

	myJson, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		panic(err)
	}
	//
	//buffer := make([]byte, 128)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String("banana1.json"),
		//ACL:                  aws.String("private"),
		Body: bytes.NewReader(myJson),
		//ContentLength:        aws.Int64(size),
		//ContentType:          aws.String(http.DetectContentType(buffer)),
		//ContentDisposition:   aws.String("attachment"),
		//ServerSideEncryption: aws.String("AES256"),
	})

}
