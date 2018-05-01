FROM debian:latest
RUN apt-get update && apt-get install dialog
RUN apt-get install curl php-symfony php-xml git php-zip php-sqlite3 apache2 libapache2-mod-php -y
RUN curl -sS https://getcomposer.org/installer | php && mv composer.phar /usr/local/bin/composer
RUN useradd -ms /bin/bash symfony
USER symfony
WORKDIR /home/symfony
#You don't need to run as the symfony user, use guso!
RUN composer create-project symfony/symfony-demo

USER root

EXPOSE 80

#TODO here: Install and configure apache as per symfony docs
#https://symfony.com/doc/current/setup/web_server_configuration.html
