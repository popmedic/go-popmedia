#!/usr/bin/env bash

set -e

source build.sh
echo "running ${app_path}"
${app_path}
