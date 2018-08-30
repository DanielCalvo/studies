#!/usr/bin/env bash

echo "Hello! You have ran the server script"

#Java as a dependency
apt-get update
apt-get install openjdk-8-jdk vim curl unzip -y

#Installing elasticsearch
wget -qO - https://packages.elastic.co/GPG-KEY-elasticsearch | apt-key add -
echo "deb http://packages.elastic.co/elasticsearch/2.x/debian stable main" | tee -a /etc/apt/sources.list.d/elasticsearch-2.x.list
apt-get update
apt-get install elasticsearch -y
cp -v /vagrant/files/elasticsearch.yml /etc/elasticsearch/elasticsearch.yml
service elasticsearch restart
systemctl enable elasticsearch

#Installing Kibana
echo "deb http://packages.elastic.co/kibana/4.5/debian stable main" | tee -a /etc/apt/sources.list.d/kibana-4.5.x.list
apt-get update
apt-get install kibana -y --allow-unauthenticated
systemctl enable kibana
cp -v /vagrant/files/kibana.yml /opt/kibana/config/kibana.yml
service kibana restart

#Nginx
apt-get install nginx apache2-utils -y
cp -v /vagrant/files/nginx_default /etc/nginx/sites-available/default
service nginx restart

#Logstash
echo 'deb http://packages.elastic.co/logstash/2.2/debian stable main' | tee /etc/apt/sources.list.d/logstash-2.2.x.list
apt-get update
apt-get install logstash -y --allow-unauthenticated
cp -v /vagrant/files/02-beats-input.conf /etc/logstash/conf.d/02-beats-input.conf
cp -v /vagrant/files/10-syslog-filter.conf /etc/logstash/conf.d/10-syslog-filter.conf
cp -v /vagrant/files/30-elasticsearch-output.conf /etc/logstash/conf.d/30-elasticsearch-output.conf
systemctl enable logstash
service logstash restart

#Kibana dashboards
curl -L -O https://download.elastic.co/beats/dashboards/beats-dashboards-1.1.0.zip
unzip beats-dashboards-1.1.0.zip
cd beats-dashboards-1.1.0
./load.sh
cd ~

#Filebeat index template
curl -O https://gist.githubusercontent.com/thisismitch/3429023e8438cc25b86c/raw/d8c479e2a1adcea8b1fe86570e42abab0f10f364/filebeat-index-template.json
curl -XPUT 'http://localhost:9200/_template/filebeat?pretty' -d@filebeat-index-template.json

#Filebeat
curl -L -O https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-6.4.0-amd64.deb
sudo dpkg -i filebeat-6.4.0-amd64.deb
cp -v /vagrant/files/filebeat.yml /etc/filebeat/filebeat.yml
service filebeat restart
systemctl enable filebeat








