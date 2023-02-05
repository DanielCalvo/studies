package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	bucketName      string
	objectPrefix    string
	objectDelimiter string
	maxKeys         int
)

/*
Dani says: I copied and pasted this from an example somewhere but I can't find the source :o
*/

func init() {
	flag.StringVar(&bucketName, "bucket", "s3-cost-calc-test-bucket", "The `name` of the S3 bucket to list objects from.")
	flag.StringVar(&objectPrefix, "prefix", "", "The optional `object prefix` of the S3 Object keys to list.")
	flag.StringVar(&objectDelimiter, "delimiter", "",
		"The optional `object key delimiter` used by S3 List objects to group object keys.")
	flag.IntVar(&maxKeys, "max-keys", 0,
		"The maximum number of `keys per page` to retrieve at once.")
}

// Lists all objects in a bucket using pagination
func main() {
	flag.Parse()
	if len(bucketName) == 0 {
		flag.PrintDefaults()
		log.Fatalf("invalid parameters, bucket name required")
	}

	// Load the SDK's configuration from environment and shared config, and
	// create the client with this.
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	// Set the parameters based on the CLI flag inputs.
	params := &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	}
	if len(objectPrefix) != 0 {
		params.Prefix = &objectPrefix
	}
	if len(objectDelimiter) != 0 {
		params.Delimiter = &objectDelimiter
	}

	p := s3.NewListObjectsV2Paginator(client, params, func(o *s3.ListObjectsV2PaginatorOptions) {
		if v := int32(maxKeys); v != 0 {
			o.Limit = v
		}
	})

	// Iterate through the S3 object pages, printing each object returned.
	var i int
	log.Println("Objects:")
	for p.HasMorePages() { //This is gonnna take a while on a large bucket -- maybe create a bucket with 23498 objects or something
		i++

		// Next Page takes a new context for each page retrieval. This is where you could add timeouts or deadlines.
		page, err := p.NextPage(context.TODO())
		if err != nil {
			log.Fatalf("failed to get page %v, %v", i, err)
		}

		for _, obj := range page.Contents {
			fmt.Println(*obj.Key, obj.Size, *obj.LastModified, obj.StorageClass)
		}
	}
}

/*
Storage class can be one of:
"STANDARD"
"REDUCED_REDUNDANCY"
"GLACIER"
"STANDARD_IA"
"ONEZONE_IA"
"INTELLIGENT_TIERING"
"DEEP_ARCHIVE"
"OUTPOSTS"
"GLACIER_IR"
)
*/
