The idea here was to list all the objects in a bucket, but I eventually found out that using S3 inventory was significantly cheaper for doing this on large buckets, so I dropped the idea 
- S3 inventory: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-inventory.html

The reason for doing this was to calculate per object and per folder costs!