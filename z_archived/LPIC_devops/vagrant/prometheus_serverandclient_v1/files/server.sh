#!/usr/bin/env bash

echo "Hello! You have ran the prometheus script"
apt-get update
apt-get install prometheus net-tools -y

cp -v /vagrant/files/prometheus.yml /etc/prometheus/prometheus.yml

echo "Restarting Prometheus"
service prometheus restart
echo "Prometheus restarted"