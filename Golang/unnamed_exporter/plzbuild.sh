#!/usr/bin/env bash

#I think you'll need to handle the absence of vendors here first with dep ensure or something along those lines

docker build . -t vioseven/unnamed-exporter:latest

sleep 1

docker push vioseven/unnamed-exporter:latest

sleep 1

kubectl delete -f manifests/031_unnamed-exporter/deployment.yaml

sleep 1

kubectl apply -f manifests/031_unnamed-exporter/deployment.yaml