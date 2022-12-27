# 使用 golang 官方镜像提供 Go 运行环境，并且命名为 buidler 以便后续引用
FROM golang:1.17-alpine as builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 设置 GOPROXY 环境变量
RUN go env -w GOPROXY=https://goproxy.io,direct

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装 make
RUN apk add gcc g++ make libffi-dev openssl-dev libtool

# 设置工作目录
WORKDIR /app/go-rest-api

# 将项目代码拷贝到镜像中
COPY . .

# 构建二进制文件，添加来一些额外参数以便可以在 Alpine 中运行它
RUN make

# 下面是第二阶段的镜像构建
FROM alpine:latest

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk --update add --no-cache tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

# 和上个阶段一样设置工作目录
RUN mkdir /app
WORKDIR /app

# 这一步不再从宿主机拷贝二进制文件，而是从上一个阶段构建的 builder 容器中拉取
COPY --from=builder /app/go-rest-api/go-rest-api .
COPY --from=builder /app/go-rest-api/conf/*.* ./conf/

# 启动 grpc server
CMD ["./go-rest-api"]
EXPOSE 9000 