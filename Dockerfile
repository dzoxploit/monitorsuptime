FROM alpine:edge

ENV GOPATH /go
ENV PATH /go/src/github.com/dzoxploit/monitorsuptime/bin:$PATH

ADD . /go/src/github.com/dzoxploit/monitorsuptime

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories \
    && apk add --no-cache --update bash ca-certificates \
    && apk add --no-cache --virtual .build-deps go gcc git libc-dev \
    && mkdir -p /configs /usr/local/bin /var/log/monitorsuptime \
    && go get github.com/gregdel/pushover \
    && cd /go/src/github.com/dzoxploit/monitorsuptime \
    && go build -v -o /usr/local/bin/monitorsuptime cmd/monitorsuptime/main.go \
    && apk del --purge .build-deps \
    && rm -rf /var/cache/apk*

ADD configs /configs

CMD ["monitorsuptime", "-config", "/configs/default1.json", "-http", ":8000", "-log", "/var/log/gossm/gossm.log"]

EXPOSE 8000