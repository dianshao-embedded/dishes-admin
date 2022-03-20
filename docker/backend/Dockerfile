FROM golang:1.17.6-alpine3.15 AS builder

ENV GOPROXY=https://proxy.golang.com.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk add --no-cache make bash git gcc g++ && mkdir -p /dishes-admin

COPY . /dishes-admin

WORKDIR /dishes-admin

RUN make build

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk --no-cache add ca-certificates tzdata && mkdir -p /dishes-admin

COPY --from=builder /dishes-admin/cmd/dishes-admin/dishes-admin /dishes-admin

COPY ./configs /dishes-admin/configs

WORKDIR /dishes-admin
