- Following: https://docs.aws.amazon.com/AmazonECR/latest/userguide/getting-started-cli.html

```shell script
docker build -t hello-world .
#docker run -t -i -p 80:80 hello-world #To test
```

```shell script
aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.eu-west-1.amazonaws.com
```

```shell script
aws ecr create-repository \
    --repository-name danitest1 \
    --image-scanning-configuration scanOnPush=true \
    --region eu-west-1
```
```shell script
docker tag hello-world:latest ${AWS_ACCOUNT_ID}.dkr.ecr.eu-west-1.amazonaws.com/hello-world:latest
docker push ${AWS_ACCOUNT_ID}.dkr.ecr.eu-west-1.amazonaws.com/hello-world:latest

```