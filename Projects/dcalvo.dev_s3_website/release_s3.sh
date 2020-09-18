#!/usr/bin/env bash

#terraform destroy can't destroy a bucket with something in it
#It would be cool if I could load the output of the terraform in this
#2020 and releasing software is still a hacky job

#IF NO ARGUMENTS
#echo "Please provide one of the following arguments: release, destroy

#IF RELEASE
cd terraform
terraform init
terraform apply
cd ..
#hey why are you not using the output as a variable in here?
aws s3 cp ./assets s3://dcalvo.dev --recursive --acl public-read

#IF DESTROY
aws s3 rb s3://dcalvo.dev --force
cd terraform
#Empty the bucket
terraform destroy -auto-approve



#check if files are on the s3 bucket
#if files are not on the repo, upload them