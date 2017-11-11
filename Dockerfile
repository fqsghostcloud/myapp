FROM centos:latest
MAINTAINER fqsghostcloud

ENV OEM=myapp VER=0.4 
ADD ./ /myapp
WORKDIR /myapp