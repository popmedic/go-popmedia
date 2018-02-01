#!/usr/bin/env bash

set -e

source build.sh
echo "running ${app_path} -config=\"server/config.json\""
${app_path} -config="server/config.json"
