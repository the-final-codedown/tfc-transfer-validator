#!/bin/bash

sh copy-cap.sh

[[ ! -d src/vendor ]] && sh vendor.sh

docker build -t tfc/tfc-transfer-validator -f Dockerfile-build .