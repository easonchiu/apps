# 使用 Ubuntu 作为基础镜像
FROM ubuntu:latest

# 将编译好的可执行文件复制到容器中
COPY run /app/

# 设置工作目录
WORKDIR /app

# 暴露程序运行的端口
EXPOSE 8000

# 指定容器启动时执行的命令
CMD ["./run"]