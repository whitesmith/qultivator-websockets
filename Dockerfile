FROM golang:1.6

COPY server/server /go/bin/server

CMD ["/go/bin/server"]