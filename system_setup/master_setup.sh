#!/bin/bash

#Assumes you found a way to get this script into the server


apt-get update
apt-get upgrade -y

apt-get install vim git

git config --global user.email "vioseven@gmail.com"
git config --global user.name "Daniel Calvo"
