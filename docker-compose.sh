#!/bin/bash

set -e
cd "$(dirname "$0")"

go env -w GOPROXY=https://goproxy.cn,direct

go test -v ./...