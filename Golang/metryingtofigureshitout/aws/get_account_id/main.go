package main

// https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.RunInstances
// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ec2-example-create-images.html

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		panic(err)
	}

	input := &sts.GetCallerIdentityInput{}
	svc := sts.New(sess)
	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(*result.Account)

}
