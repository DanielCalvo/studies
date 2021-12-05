```shell
docker build . -t docker_s3_example

docker run \
--env AWS_ACCESS_KEY_ID=KEY \
--env AWS_SECRET_ACCESS_KEY=SECRET \
--env AWS_DEFAULT_REGION=eu-west-1 \
--env MY_SAMPLE_BUCKET=mysamplebucket9342348374 \
docker_s3_example
```


