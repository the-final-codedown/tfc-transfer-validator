#!/bin/bash

sh copy-cap.sh

[[ ! -d src/vendor ]] && sh vendor.sh
IMAGE=tfc/transfer-validator
VERSION=

docker build -t ${IMAGE} -f Dockerfile-build .
docker tag ${IMAGE} localhost:5000/${IMAGE}${VERSION}
docker push localhost:5000/${IMAGE}${VERSION}