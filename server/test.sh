#!/usr/bin/env bash

set -e

function test {
    go test -cover
}

test