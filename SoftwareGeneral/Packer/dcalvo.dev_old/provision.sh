#!/usr/bin/env bash
sudo apt-get update
sudo apt-get install nginx -y
sudo systemctl enable --now nginx
sudo mv /home/ubuntu/assets/* /var/www/html
