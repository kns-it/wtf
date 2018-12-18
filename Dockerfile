FROM golang:1.11-alpine as build

RUN apk add --no-cache git dep

WORKDIR /go/src/github.com/kns-it/wtf

ADD ./ ./

RUN dep ensure && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o wtf cmd/wtf/main.go

FROM alpine:3.8

LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="WTF" \
      org.label-schema.description="Small and fast utility to run network diagnostics" \
      org.label-schema.url="https://github.com/kns-it/wtf" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/kns-it/wtf" \
      org.label-schema.vendor="" \
      org.label-schema.version="0.0.1" \
      org.label-schema.schema-version="1.0" \
      maintainer="peter.kurfer@gmail.com"

COPY --from=BUILD /go/src/github.com/kns-it/wtf/wtf /usr/local/bin/wtf