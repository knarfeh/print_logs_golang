FROM golang:1.7
MAINTAINER knarfeh@outlook.com

WORKDIR $GOPATH/src/github.com/knarfeh/print_logs_golang
ADD . $GOPATH/src/github.com/knarfeh/print_logs_golang
RUN go build .

CMD ["./start.sh"]
