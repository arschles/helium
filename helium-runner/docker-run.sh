#!/bin/bash

docker run --rm -v $(pwd)/../.helium:/helium -e HELIUM_SCRIPTS_DIR=/helium -e HELIUM_TARGET_DIR=/home/src/dist/src helium-runner:v0.0.1
docker run -it --rm -v $(pwd)/../.helium:/helium -e HELIUM_SCRIPTS_DIR=/helium -e HELIUM_TARGET_DIR=/home/src/dist/src helium-runner:v0.0.1 ash
