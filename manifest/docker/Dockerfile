FROM loads/alpine:3.8

ENV TZ=Asia/Shanghai
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# 复制主文件
COPY gztv /gztv/gztv
# 复制配置文件
COPY conf /gztv/conf
# 复制public文件
COPY public /gztv/public
# 进入工作目录
WORKDIR /gztv
# 设置时区
ENV TZ=Asia/Shanghai
# 启动锚点
ENTRYPOINT ["./gztv"]
# 导出端口 6020 6021 6022 1936 4433 22
EXPOSE 6020 6021 6022 1936 4433 22