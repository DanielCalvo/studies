#!/usr/bin/env bash

#terraform destroy can't destroy a bucket with something in it
#It would be cool if I could load the output of the terraform in this
#2020 and releasing software is still a hacky job

cd terraform
terraform init
terraform apply
cd ..

#hey why are you not using the output as a variable in here?
aws s3 cp ./assets s3://dcalvo-dev-bucket --recursive --grants Permission=read

#check if files are on the s3 bucket
#if files are not on the repo, upload them