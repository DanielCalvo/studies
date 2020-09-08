#!/usr/bin/env bash

#Don't forget to have your AWS credentials configured on your local

#Step 1: Build packer image
packer build packer/dcalvo_dev.json

#Step 2: Get ID of the packer image?
#Nah this phase is in terraform I think

#Launch instance with this image

#Step 3: Point DNS to new instance
#also in tf