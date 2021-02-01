FROM golang:alpine as builder

ENV GOPROXY https://goproxy.io,https://goproxy.cn,direct
ENV GO111MODULE on
ENV LANG=en_US.UTF-8
ENV LC_ALL=en_US.UTF-8

WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o bee_gin_docker main.go

# RUN go build -o main .
RUN ls -al
EXPOSE 8080
ENTRYPOINT ["./main"]

