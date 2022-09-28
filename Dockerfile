FROM node:16 AS node-builder

LABEL maintainer="mingcheng<mc@sf.com>"

COPY . /answer
WORKDIR /answer

RUN make install-ui-packages ui
RUN mv ui/build /tmp
CMD ls -al /tmp
RUN du -sh /tmp/build


FROM golang:1.18 AS golang-builder
LABEL maintainer="aichy"

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PACKAGE github.com/segmentfault/answer
ENV GOPROXY https://goproxy.cn,direct
ENV BUILD_DIR ${GOPATH}/src/${PACKAGE}
ENV GOPRIVATE git.backyard.segmentfault.com
# Build
COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}
COPY --from=node-builder /tmp/build ${BUILD_DIR}/web/html
RUN make clean build && \
	cp answer /usr/bin/answer && \
    mkdir -p /tmp/cache && chmod 777 /tmp/cache && \
    mkdir /data && chmod 777 /data && cp configs/config.yaml /data/config.yaml && \
    mkdir -p /data/upfiles && chmod 777 /data/upfiles && cp -r i18n /data



FROM debian:bullseye
ENV TZ "Asia/Shanghai"
RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list \
        && sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list \
        && echo "Asia/Shanghai" > /etc/timezone \
        && apt -y update \
        && apt -y upgrade \
        && apt -y install ca-certificates openssl tzdata curl netcat dumb-init \
        && apt -y autoremove

COPY --from=golang-builder /data /data
VOLUME /data

COPY --from=golang-builder /usr/bin/answer /usr/bin/answer
COPY /script/entrypoint.sh /entrypoint.sh
RUN chmod 755 /entrypoint.sh

EXPOSE 80
ENTRYPOINT ["/entrypoint.sh"]
