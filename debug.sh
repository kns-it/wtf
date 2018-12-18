#!/usr/bin/env bash

dlv_path=$(which dlv)

if [[ $? -ne 0 ]]; then
    echo "Installing delve because it's not in PATH"
    go get -u github.com/derekparker/delve/cmd/dlv
fi

dlv debug --headless --listen=:2345 --api-version=2 ./cmd/wtf/main.go