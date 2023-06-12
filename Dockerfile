FROM golang:1.19-alpine AS golang-builder
LABEL maintainer="aichy@sf.com"

ARG GOPROXY
# ENV GOPROXY ${GOPROXY:-direct}
ENV GOPROXY=https://goproxy.io,direct

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PACKAGE github.com/answerdev/answer
ENV BUILD_DIR ${GOPATH}/src/${PACKAGE}
ENV ANSWER_MODULE ${BUILD_DIR}

ARG TAGS="sqlite sqlite_unlock_notify"
ENV TAGS "bindata timetzdata $TAGS"
ARG CGO_EXTRA_CFLAGS

COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}
RUN apk --no-cache add build-base git bash nodejs npm && npm install -g pnpm corepack \
    && make install-ui-packages clean build

RUN chmod 755 answer
RUN ["/bin/bash","-c","script/build_plugin.sh"]
RUN cp answer /usr/bin/answer

RUN mkdir -p /data/uploads && chmod 777 /data/uploads \
    && mkdir -p /data/i18n && cp -r i18n/*.yaml /data/i18n

FROM alpine
LABEL maintainer="maintainers@sf.com"

ENV TZ "Asia/Shanghai"
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
    && echo "Asia/Shanghai" > /etc/timezone

COPY --from=golang-builder /usr/bin/answer /usr/bin/answer
COPY --from=golang-builder /data /data
COPY /script/entrypoint.sh /entrypoint.sh
RUN chmod 755 /entrypoint.sh

VOLUME /data
EXPOSE 80
ENTRYPOINT ["/entrypoint.sh"]
