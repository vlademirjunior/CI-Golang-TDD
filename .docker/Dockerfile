FROM golang:alpine as builder

WORKDIR $GOPATH/src/

COPY ./src .

RUN go build -ldflags="-s -w" -o /go/app

FROM scratch

COPY --from=builder /go/app /go/app

CMD ["/go/app"]
