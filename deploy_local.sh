#!/usr/bin/env bash

source deploy.sh

pushd ${artifact_dir} > /dev/null
unzip "${zip}" -d ${zip%.zip}
pushd "${zip%.zip}" > /dev/null
sudo ./install.sh
popd > /dev/null
popd > /dev/null
