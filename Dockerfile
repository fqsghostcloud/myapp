FROM centos:latest
MAINTAINER fqsghostcloud

ENV OEM=myapp VER=0.6
ADD ./ /myapp
WORKDIR /myapp