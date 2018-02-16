#!/usr/bin/env bash

set -e

function test {
    go test -cover github.com/popmedic/go-popmedia/server/...
}

test