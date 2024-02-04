#!/usr/bin/env sh



if [ "$#" -gt 0 ]; then
    BINARY=nxrmuploader
fi

go build -o /opt/bin/${BINARY} .
