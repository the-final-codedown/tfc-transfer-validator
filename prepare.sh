#!/bin/bash

mkdir -p ./tfc-cap-updater
cp -r ../tfc-cap-updater/*.go ./tfc-cap-updater/
cp -r ../tfc-cap-updater/go.* ./tfc-cap-updater/
cp -r ../tfc-cap-updater/proto ./tfc-cap-updater/proto

[[ ! -d ./vendor ]] && sh vendor.sh

docker build -t tfc/tfc-transfer-validator -f Dockerfile-build .