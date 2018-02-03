#!/usr/bin/env bash

source test.sh

app_name="popmedia-server"
main_path="server/cmd/main.go"

project_root="."
artifact_dir="${project_root}/artifact"
bin_dir="${artifact_dir}/bin"
app_path="${bin_dir}/${app_name}"

function build {
    mkdir -p "$1"
    app_path="$1/${app_name}"
    if [ -z ${2+x} ] 
    then
        goos=""
    else
        goos="${2}"
    fi
    if [ -z ${3+x} ] 
    then
        goarch=""
    else
        goarch="${3}"
    fi
    echo "building ${app_name} for ${goos}/${goarch}..."
    GOOS="${goos}" GOARCH="${goarch}" go build -o "${app_path}" "${main_path}"
}
mkdir -p "${bin_dir}"
cp -r server/templates "${bin_dir}"
cp server/config.json "${bin_dir}"
build "${bin_dir}"