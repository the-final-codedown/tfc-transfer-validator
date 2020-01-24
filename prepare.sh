#!/bin/bash

mkdir -p ./vendor/tfc-cap-updater
cp -r ../tfc-cap-updater/ ./vendor/tfc-cap-updater/
go mod vendor
docker build -t tfc/tfc-transfer-validator -f Dockerfile-build .