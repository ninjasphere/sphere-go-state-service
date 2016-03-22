FROM       ubuntu:14.04
MAINTAINER Ninja Blocks <developers@ninjablocks.com>

RUN        apt-get -qy update && apt-get -qy install vim-common gcc mercurial supervisor


COPY build/sphere-go-state-service /app/
ADD etc/supervisord.conf /app/etc/

WORKDIR /app

EXPOSE     6100
CMD ["/usr/bin/supervisord"]