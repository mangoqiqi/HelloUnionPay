FROM registry.cn-hangzhou.aliyuncs.com/choerodon-tools/golang:1.9.4-alpine3.7 as builder
WORKDIR /go/src/github.com/unionpay
COPY . /go/src/github.com/unionpay
RUN go build .

FROM registry.saas.hand-china.com/tools/alpine:latest
WORKDIR /
COPY --from=builder /go/src/github.com/unionpay .
CMD ["/unionpay", "run"]
