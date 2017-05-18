#!/bin/sh

if [ "$WRITE_TO_FILES" != "" ]; then
    /go/src/github.com/knarfeh/print_logs_golang/print_logs_golang > log_files1.log &
    /go/src/github.com/knarfeh/print_logs_golang/print_logs_golang > log_files2.log &
fi

/go/src/github.com/knarfeh/print_logs_golang/print_logs_golang
