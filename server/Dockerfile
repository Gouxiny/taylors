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
# 时区设置成当前时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

EXPOSE 9999
ENTRYPOINT /taylors/app

# docker build -t taylors .
# docker run --name taylors -p 80:9999 -d taylors
#