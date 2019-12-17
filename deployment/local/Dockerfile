FROM golang:1.13-alpine AS builder

# TODO 目前路徑有錯誤
COPY . /go/src/guava 
WORKDIR /go/src/guava
RUN go build -mod vendor -o server .

FROM golang:1.13-alpine
COPY --from=builder /go/src/guava/server /go/src/guava/server
COPY ./configs /go/src/guava/configs

WORKDIR /go/src/guava
ENTRYPOINT [ "./server" ]