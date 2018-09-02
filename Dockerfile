FROM golang:stretch

LABEL maintainer="mas9612 <mas9612@gmail.com>"

RUN go get -u github.com/mas9612/slackbase
WORKDIR /go/src/github.com/mas9612/slackbase
RUN go build ./cmd/slackbot

CMD ["./slackbot"]
