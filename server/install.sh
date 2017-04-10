#!/usr/bin/env bash

idir="/var/popmedia-server/bin"

mkdir -p "/var/popmedia-server/bin"

echo "copy config.json to ${idir}"
cp -n config.json ${idir}

echo "copy popmedia-server to ${idir}"
cp -f popmedia-server ${idir}

echo "copy templates to ${idir}"
cp -rf templates ${idir}

echo "copy services script to ${idir}"
cp -f popmedia-server-service.sh ${idir}

echo "installing services"
ln -sf "${idir}/popmedia-server-service.sh" "/etc/init.d/popmedia-server-service"