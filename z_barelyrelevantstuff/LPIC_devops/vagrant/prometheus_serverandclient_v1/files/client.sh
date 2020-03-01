#!/usr/bin/env bash

apt-get update
apt-get install golang apache2 net-tools git -y

mkdir -p ~/go
export GOPATH=/root/go

go get github.com/neezgee/apache_exporter
ln -sf /root/go/bin/apache_exporter /usr/bin
cp -v /vagrant/files/apache_exporter.conf /etc/init/apache_exporter.conf

#service apache_exporter start

/usr/bin/apache_exporter &