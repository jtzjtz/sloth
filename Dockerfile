FROM ubuntu:latest
RUN apt-get update -y  \
    &&  apt-get upgrade -y \
    &&  apt install git -y \
    && apt-get install tar -y \
    && apt-get install wget -y \
    && apt-get install autoconf automake libtool curl make g++ unzip -y \
    && mkdir -p /data/src \
    && wget -c https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -O - |  tar -xz -C /usr/local
ENV GOROOT /usr/local/go
ENV GOPATH /data
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH
RUN echo $PATH \
    && go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.io,direct
RUN  cd /data/src \
    && git clone  https://github.com/google/protobuf \
    && cd protobuf \
    && git submodule update --init --recursive \
    && ./autogen.sh \
    && ./configure \
    && make \
#    && make check \
    && make install \
    && ldconfig \
    && go get -u github.com/golang/protobuf/proto \
    && go get -u github.com/golang/protobuf/protoc-gen-go



COPY ./output/app /data/web-app
COPY ./*.sh /data/

COPY ./app/html/  /data/app/html/
COPY ./app/resource/layui /data/app/resource/layui
COPY ./app/resource/ /data/app/resource/

RUN chmod +x /data/web-app
WORKDIR /data/
ENTRYPOINT ["./web-app"]
