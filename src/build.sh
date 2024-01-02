#!/usr/bin/env sh



if [ "$#" -gt 0 ]; then
    BINARY=uploadNxRM
fi

go build -o /opt/bin/${BINARY} .
