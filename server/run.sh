#!/usr/bin/env bash

set -e

source build.sh
echo "running ${app_path}"
open -a Google\ Chrome "http://localhost:8080/"
${app_path}
