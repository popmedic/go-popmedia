#!/usr/bin/env bash

source deploy.sh

goos=(darwin \
    darwin \
    darwin \
    dragonfly \
    freebsd \
    freebsd \
    freebsd \
    linux \
    linux \
    linux \
    linux \
    linux \
    linux \
    netbsd \
    netbsd \
    netbsd \
    openbsd \
    openbsd \
    openbsd \
    plan9 \
    plan9 \
    solaris \
    windows \
    windows)

goarch=(386 \
    amd64 \
    arm \
    amd64 \
    386 \
    amd64 \
    arm \
    386 \
    amd64 \
    arm \
    arm64 \
    ppc64 \
    ppc64le \
    386 \
    amd64 \
    arm \
    386 \
    amd64 \
    arm \
    386 \
    amd64 \
    amd64 \
    386 \
    amd64)

i=0
for os in "${goos[@]}"
do
    arch="${goarch[${i}]}"
    deploy "${os}" "${arch}"
    i=${i}+1
done
