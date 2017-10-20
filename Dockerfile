FROM golang:1.9.0 AS build-env
ADD . /go/src/github.com/surajssd/tweetbase
RUN export GOBIN=/go/bin/ && \
    cd /go/src/github.com/surajssd/tweetbase && \
    go install tweetbase.go

# final stage
FROM alpine

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=build-env /go/bin/tweetbase /usr/local/bin/
