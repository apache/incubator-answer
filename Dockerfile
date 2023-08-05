FROM golang:1.19-alpine AS golang-builder
LABEL maintainer="mui"

ARG GOPROXY
# ENV GOPROXY ${GOPROXY:-direct}
ENV GOPROXY=https://proxy.golang.com.cn,direct

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PACKAGE github.com/khamui/amapaw
ENV BUILD_DIR ${GOPATH}/src/${PACKAGE}
ENV ANSWER_MODULE ${BUILD_DIR}

ARG TAGS="sqlite sqlite_unlock_notify"
ENV TAGS "bindata timetzdata $TAGS"
ARG CGO_EXTRA_CFLAGS

COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}
RUN apk --no-cache add build-base git bash nodejs npm
RUN npm install -g pnpm corepack
RUN make install-ui-packages
RUN make clean
RUN make build

RUN chmod 755 answer
RUN ["/bin/bash","-c","script/build_plugin.sh"]
RUN cp answer /usr/bin/answer

RUN mkdir -p /data/uploads && chmod 777 /data/uploads \
    && mkdir -p /data/i18n && cp -r i18n/*.yaml /data/i18n

FROM alpine

ENV TZ "Europe/Berlin"
RUN apk update \
    && apk --no-cache add \
        bash \
        ca-certificates \
        curl \
        dumb-init \
        gettext \
        openssh \
        sqlite \
        gnupg \
    && echo "Europe/Berlin" > /etc/timezone

COPY --from=golang-builder /usr/bin/answer /usr/bin/answer
COPY --from=golang-builder /data /data
COPY /script/entrypoint.sh /entrypoint.sh
RUN chmod 755 /entrypoint.sh

VOLUME /data
EXPOSE 80
ENTRYPOINT ["/entrypoint.sh"]
