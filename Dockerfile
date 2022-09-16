FROM --platform=linux/amd64 amd64/alpine:3.12.0

ENV PATH /usr/local/bin:$PATH
ENV LANG C.UTF-8

ENV TZ=Asia/Shanghai

RUN apk update && apk upgrade \
    && apk add ca-certificates\
    && update-ca-certificates \
    && apk --no-cache add openssl wget \
	&& apk add --no-cache bash tzdata curl \
	&& set -ex \
    && mkdir -p /usr/bin \
    && mkdir -p /usr/sbin \
    && mkdir -p /data/wechat-webhook/

ADD wechatrobot /usr/bin/

WORKDIR /data/wechat-webhook/

CMD ["./wechatrobot"]
