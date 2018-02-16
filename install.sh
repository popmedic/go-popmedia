#!/usr/bin/env bash

idir="/var/popmedia-server"

mkdir -p "${idir}"

echo "copy config.json to ${idir}"
cp -n config.json ${idir}

echo "copy popmedia-server to ${idir}"
cp -f popmedia-server ${idir}

echo "copy templates to ${idir}"
cp -rf templates ${idir}

echo "copy images to ${idir}"
cp -rf images ${idir}

echo "copy services script to ${idir}"
cp -f popmedia-server-service.sh ${idir}

echo "installing services"
if [ $(uname) == "Darwin" ]; then
    if [[ ! -z $(launchctl list | grep popmedic) ]]; then
        launchctl unload /Library/LaunchDaemons/com.popmedic.go-popmedia.plist
    fi
    cp com.popmedic.go-popmedia.plist /Library/LaunchDaemons
    launchctl load /Library/LaunchDaemons/com.popmedic.go-popmedia.plist
else
    ln -sf "${idir}/popmedia-server-service.sh" "/etc/init.d/popmedia-server-service"
    services popmedia-server-service start
fi