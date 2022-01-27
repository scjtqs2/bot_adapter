FROM golang:1.17-alpine AS builder
RUN  sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache git \
  && go env -w GO111MODULE=auto \
  && go env -w CGO_ENABLED=1 \
  && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY ./ .

RUN set -ex \
    && BUILD=`date +%FT%T%z` \
    && COMMIT_SHA1=`git rev-parse HEAD` \
    && go build -ldflags "-s -w -extldflags '-static' -X main.Version=${COMMIT_SHA1}|${BUILD}" -v -o bot_adapter


FROM alpine AS production

RUN  sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

ENV ADAPTER_SECRET "secret"
ENV HTTP_LISTEN "0.0.0.0:5800"
ENV HTTP_TOKEN ""
ENV HTTP_SECRET ""
ENV HTTP_CQHTTP_ADDR ""


COPY ./init.sh /
COPY --from=builder /build/bot_adapter /usr/bin/bot_adapter
RUN chmod +x /usr/bin/bot_adapter && chmod +x /init.sh

WORKDIR /data

ENTRYPOINT [ "/init.sh" ]