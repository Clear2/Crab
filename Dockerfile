FROM golang:alpine as builder

ENV GOPROXY https://goproxy.io,https://goproxy.cn,direct
ENV GO111MODULE on
ENV LANG=en_US.UTF-8
ENV LC_ALL=en_US.UTF-8

WORKDIR /app
COPY  . .
RUN ls -al
