#!/bin/sh

if [ "$WRITE_TO_FILES" != "" ]; then
    print_logs_golang > log_files1.log &
    print_logs_golang > log_files2.log &
fi

print_logs_golang
