#!/bin/bash

rm -rf output

mkdir -p output/{bin,conf}

go build -o output/bin/tiktok *.go