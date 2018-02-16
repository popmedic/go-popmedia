#!/usr/bin/env bash
source clean.sh
source build.sh

function deploy {
    os=$1
    arch=$2
    bd="${artifact_dir}/${os}/${arch}/bin"
    GOOS=os
    GOARCH=arch

    build "${bd}" "${os}" "${arch}"

    config_path="server/config.json"
    tmpl="server/templates"
    imgs="server/images"
    zip="${app_name}-${os}-${arch}.zip"
    svc="server/cmd/popmedia-server-service.sh"
    plist="server/cmd/com.popmedic.go-popmedia.plist"
    install="install.sh"

    echo "copy ${config_path} to ${bd}"
    cp -f "${config_path}" "${bd}"

    echo "copy ${tmpl} to ${bd}"
    cp -rf "${tmpl}" "${bd}"

    echo "copy ${imgs} to ${bd}"
    cp -rf "${imgs}" "${bd}"

    echo "copy ${svc} to ${bd}"
    cp -f "${svc}" "${bd}"

    echo "copy ${plist} to ${bd}"
    cp -f "${plist}" "${bd}"

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
