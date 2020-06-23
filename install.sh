#!/bin/bash

set -e

export GOPROXY=https://goproxy.cn,direct
export GOPATH=/usr/local/go/packages
export GOROOT=/usr/local/go
export GO111MODULE=on
export GOSUMDB=off

make install_all