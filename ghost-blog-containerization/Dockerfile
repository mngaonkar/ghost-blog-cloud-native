FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y sudo
RUN apt-get install -y nginx
RUN apt-get install -y curl
RUN curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash 
RUN apt-get install -y nodejs
RUN apt-get install -y python

RUN npm install -g ghost-cli

RUN mkdir -p /var/www/ghost
RUN useradd -ms /bin/bash admin
RUN echo "admin:admin" | chpasswd
RUN passwd -d admin
RUN adduser admin sudo
RUN chown admin:admin /var/www/ghost
RUN chmod 775 /var/www/ghost

USER admin
WORKDIR /var/www/ghost

EXPOSE 2368
RUN sudo apt-get install -y nano
RUN sudo apt-get install -y sqlite3

ENV GHOST_ENV production
ENV GHOST_INSTALL /var/www/ghost
ENV GHOST_CONTENT /var/www/ghost/content

#RUN ghost install --db sqlite3 --no-prompt --no-stack --no-setup
#RUN ghost config --ip 0.0.0.0 --url http://localhost:2368 --db sqlite3 --dbpath $GHOST_CONTENT/data/ghost.db \
#--process local

RUN ghost install --db sqlite3 --process local --url http://localhost:2368 --ip 0.0.0.0 --no-setup-nginx \
--no-prompt

ENTRYPOINT ["ghost", "run"]





