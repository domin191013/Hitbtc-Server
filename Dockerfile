FROM golang:1.16-alpine

WORKDIR /go/src

COPY *.go go.mod go.sum Makefile ./

RUN apk add alpine-sdk

RUN make get-gotools && \
    make build && \
    mv hitbit-api / && \
    rm -fr /go/src/*

EXPOSE 8080

CMD [ "/hitbit-api" ]
