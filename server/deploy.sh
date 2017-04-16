#!/usr/bin/env bash

source build.sh

goos=(linux)
goarch=(amd64)
# goos=(darwin \
#     darwin \
#     darwin \
#     dragonfly \
#     freebsd \
#     freebsd \
#     freebsd \
#     linux \
#     linux \
#     linux \
#     linux \
#     linux \
#     linux \
#     netbsd \
#     netbsd \
#     netbsd \
#     openbsd \
#     openbsd \
#     openbsd \
#     plan9 \
#     plan9 \
#     solaris \
#     windows \
#     windows)
# goarch=(386 \
#     amd64 \
#     arm \
#     amd64 \
#     386 \
#     amd64 \
#     arm \
#     386 \
#     amd64 \
#     arm \
#     arm64 \
#     ppc64 \
#     ppc64le \
#     386 \
#     amd64 \
#     arm \
#     386 \
#     amd64 \
#     arm \
#     386 \
#     amd64 \
#     amd64 \
#     386 \
#     amd64)

function deploy {
    os=$1
    arch=$2
    bd="${artifact_dir}/${os}/${arch}/bin"
    GOOS=os
    GOARCH=arch

    build "${bd}" "${os}" "${arch}"

    config_path="config.json"
    tmpl="templates"
    zip="${app_name}-${os}-${arch}.zip"
    svc="cmd/popmedia-server-service.sh"
    install="install.sh"

    echo "copy ${config_path} to ${bd}"
    cp -f "${config_path}" "${bd}"

    echo "copy ${tmpl} to ${bd}"
    cp -rf "${tmpl}" "${bd}"

    echo "copy ${svc} to ${bd}"
    cp -f "${svc}" "${bd}"

    echo "copy ${install} to ${bd}"
    cp -f "${install}" "${bd}"

    # TODO: Add the images and any other resources/assets
    
    pushd ${bd} > /dev/null
    echo "creating zip ${zip}"
    find "." -path '*/.*' -prune -o -type f -print | zip "${zip}" -@
    popd > /dev/null
    mv "${bd}/${zip}" "${artifact_dir}"
    rm -rf "${artifact_dir}/${os}"
}

i=0
for os in "${goos[@]}"
do
    arch="${goarch[${i}]}"
    deploy "${os}" "${arch}"
    i=${i}+1
done
