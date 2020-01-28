#!/bin/bash

[[ -d ./tfc-cap-updater ]] && rm -r ./tfc-cap-updater
mkdir -p ./tfc-cap-updater
cp -r ../tfc-cap-updater/src/*.go ./tfc-cap-updater/
cp -r ../tfc-cap-updater/src/go.* ./tfc-cap-updater/
cp -r ../tfc-cap-updater/src/proto ./tfc-cap-updater/proto