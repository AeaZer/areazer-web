# 设置基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /web

# 拷贝项目文件到容器中
COPY . .

RUN go mod download golang.org/x/net

# 构建应用程序
RUN go build -o main .

# 暴露端口
EXPOSE 8080

# 运行应用程序
CMD ["./main"]