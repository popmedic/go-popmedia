#!/usr/bin/env bash

set -e

function test {
    go test -cover github.com/popmedic/popmedia2/server/...
}

test