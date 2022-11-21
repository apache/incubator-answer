FROM amd64/node AS node-builder

LABEL maintainer="mingcheng<mc@sf.com>"

COPY . /answer
WORKDIR /answer
RUN make install-ui-packages ui && mv ui/build /tmp

# stage2 build the main binary within static resource
FROM golang:1.19-alpine AS golang-builder
LABEL maintainer="aichy@sf.com"

ARG GOPROXY
ENV GOPROXY ${GOPROXY:-direct}

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PACKAGE github.com/answerdev/answer
ENV BUILD_DIR ${GOPATH}/src/${PACKAGE}

ARG TAGS="sqlite sqlite_unlock_notify"
ENV TAGS "bindata timetzdata $TAGS"
ARG CGO_EXTRA_CFLAGS

COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}
COPY --from=node-builder /tmp/build ${BUILD_DIR}/ui/build
RUN apk --no-cache add build-base git \
    && make clean build \
    && cp answer /usr/bin/answer

RUN mkdir -p /data/uploads && chmod 777 /data/uploads \
    && mkdir -p /data/i18n && cp -r i18n/*.yaml /data/i18n

# stage3 copy the binary and resource files into fresh container
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
