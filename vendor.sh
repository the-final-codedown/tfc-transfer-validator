#!/bin/bash

mkdir -p vendor
docker run --rm \
  -v "$(pwd)":/home/golang \
  -v "$(pwd)/tfc-cap-updater":/home/tfc-cap-updater \
  -w /home/golang \
  --entrypoint "sh" \
  golang:1.13.6-alpine3.11 \
  -c "go mod vendor"

