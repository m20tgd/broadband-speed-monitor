#!/bin/bash

# Navigate to directory
file=${BASH_SOURCE[0]}
dir=$(dirname $file)
cd $dir

# Build docker image
docker compose -f ../docker/docker-compose.yaml build