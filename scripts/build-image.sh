#!/bin/bash

docker build \
  -f build/Dockerfile \
  -t machine-learning-golang:latest \
  .
