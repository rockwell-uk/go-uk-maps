FROM golang:1.18-buster

RUN apt-get update
RUN apt-get install -y libgeos-dev

WORKDIR /app