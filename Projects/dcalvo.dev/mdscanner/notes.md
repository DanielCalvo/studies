### Actionable items

### Coding notes
- Prefer i to index
- Prefer r to reader
- Prefer b to buffer
- define and parse your flags in func main
- The last component of a package path should be the same as the package name.
  


#### Terraform
- Create the following:
    - S3 bucket (not public)
    - User to access the bucket
    - A policy for the user granting access only to that bucket

#### Golang
- Create a program that will upload files with S3 using user credentials


### md scanner ideas
- Create a cloudwatch event that launches an instance, does something and then exits?

- How do I stop and start Amazon EC2 instances at regular intervals using Lambda?
    - https://aws.amazon.com/premiumsupport/knowledge-center/start-stop-lambda-cloudwatch/

- How to connect to S3?
    - https://docs.aws.amazon.com/AmazonS3/latest/dev/security-best-practices.html

- Can you use Lambda pre signed URLs to upload files securely?
    - https://medium.com/@lakshmanLD/upload-file-to-s3-using-lambda-the-pre-signed-url-way-158f074cda6c

- How is Lambda local development done again? Can you launch a function that uploads a mock file to s3?

#### Maybe
- Just run things on an ec2 instance for now, worry about migrating to t2 micro later

#### Possible feature ideas
- Can you take the feature toggle approach as defined in the cloud ops book?
    - Application could save something either to S3 or local filesystem
    - "How to feature toggle golang"
    - https://featureflags.io/go-feature-flags/