#!/usr/bin/env bash

set -e

source build.sh
cd "${bin_dir}"
echo "running ${app_name}"
./${app_name}
