#!/bin/bash

# Ensure NUMA is enabled
for a in /sys/bus/pci/devices/*; do echo 0 | sudo tee -a $a/numa_node > /dev/null; done

docker run -it \
  -e "GOPATH=${GOPATH}" \
  -p 3000:3000 \
  -u $(id -u):$(id -g) \
  -v "${GOPATH}":"${GOPATH}" \
  -w "${PWD}/project" \
  --runtime=nvidia \
  --rm \
  --name machine-learning-golang machine-learning-golang:latest \
  bash
