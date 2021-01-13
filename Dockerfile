FROM golang:alpine as builder

ENV GOPROXY https://goproxy.io,https://goproxy.cn,direct
ENV GO111MODULE on
ENV LANG=en_US.UTF-8
ENV LC_ALL=en_US.UTF-8

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
RUN ls -al
EXPOSE 8888
ENTRYPOINT ["./main"]

