FROM debian:latest
RUN apt-get update && apt-get install dialog
RUN apt-get install curl php-symfony php-xml git php-zip php-sqlite3 apache2 libapache2-mod-php vim wget -y
COPY demovhost.conf /etc/apache2/sites-available/demovhost.conf
RUN a2ensite demovhost
RUN curl -sS https://getcomposer.org/installer | php && mv composer.phar /usr/local/bin/composer
RUN useradd -ms /bin/bash symfony
RUN echo "127.0.0.1 symfonydemo" >> /etc/hosts
USER symfony
WORKDIR /home/symfony
#You don't need to run as the symfony user, use guso!
RUN composer create-project symfony/symfony-demo && chmod -R 777 /home/symfony/symfony-demo/var/cache && chmod -R 777 /home/symfony/symfony-demo/var/logs && chmod -R 777 /home/symfony/symfony-demo/var/sessions

USER root
EXPOSE 80

CMD apachectl -D FOREGROUND
#TODO here: Install and configure apache as per symfony docs
#https://symfony.com/doc/current/setup/web_server_configuration.html
