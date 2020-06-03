FROM golang:alpine as builder
# 设置go mod proxy 国内代理
# 设置golang path
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1
WORKDIR /taylors
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
ADD . /taylors
RUN go env && go list && go build -o app main.go

EXPOSE 9999
ENTRYPOINT /taylors/app

# docker build -t taylors .
# docker run --name taylors -p 80:9999 -d taylors
#