#!/bin/bash

#This script sets up a few packages and services
#This is not managed by salt, as this script sets up the salt master among other things
#It is advisable to run this only an a fresh install of Debian 9 (stretch)

apt-get update
apt-get upgrade -y

apt-get install vim git #what else?
