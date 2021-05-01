#!/usr/bin/env bash

#https://docs.nextcloud.com/server/13/admin_manual/installation/

apt-get update

apt-get install apache2 mariadb-server libapache2-mod-php7.0 php7.0-gd php7.0-json php7.0-mysql php7.0-curl php7.0-mbstring php7.0-intl php7.0-mcrypt php-imagick php7.0-xml php7.0-zip -y

cd ~
wget https://download.nextcloud.com/server/releases/nextcloud-13.0.5.tar.bz2
tar -xjf nextcloud-13.0.5.tar.bz2
cp -rf /root/nextcloud/ /var/www/nextcloud

chown -R www-data.www-data /var/www/nextcloud

cp /vagrant/files/nextcloud.conf /etc/apache2/sites-available/nextcloud.conf

a2dissite 000-default
a2ensite nextcloud
a2enmod headers
a2enmod env
a2enmod dir
a2enmod mime
a2enmod rewrite

service apache2 restart

mysql -e "CREATE USER 'nextcloud'@'%' IDENTIFIED BY 'nextcloud';"
mysql -e "GRANT ALL ON nextcloud.* TO 'nextcloud'@'%';"