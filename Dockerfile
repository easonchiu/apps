# 使用 Ubuntu 作为基础镜像
FROM crpi-1q6uj003eu6yxxye.cn-shanghai.personal.cr.aliyuncs.com/yaoshang/ubuntu:25.10

# 将编译好的可执行文件复制到容器中
COPY run /app/

# 设置工作目录
WORKDIR /app

# 暴露程序运行的端口
EXPOSE 8000

# 数据库配置通过环境变量在运行时传入，例如：
#   MONGODB_URI=mongodb://user:pass@host:27017
#   MONGODB_DATABASE=ysgame
# 使用方式：docker run -e MONGODB_URI=... -e MONGODB_DATABASE=... 或 --env-file .env

# 指定容器启动时执行的命令
CMD ["./run"]