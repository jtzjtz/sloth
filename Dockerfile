FROM jtzjtz/ubuntu:sloth



COPY ./output/app /data/web-app
COPY ./*.sh /data/

COPY ./app/html/  /data/app/html/
COPY ./app/resource/layui /data/app/resource/layui
COPY ./app/resource/ /data/app/resource/

RUN chmod +x /data/web-app
WORKDIR /data/
expose 8000
ENTRYPOINT ["./web-app"]
