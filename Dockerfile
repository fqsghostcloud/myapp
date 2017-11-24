FROM centos:latest
MAINTAINER fqsghostcloud

EXPOSE 8080 50051
ENV OEM=myapp VER=0.6
ADD ./ /myapp
WORKDIR /myapp