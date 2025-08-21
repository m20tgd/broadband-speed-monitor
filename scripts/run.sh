#!/bin/bash

# Navigate to directory
file=${BASH_SOURCE[0]}
dir=$(dirname $file)
cd $dir

# Build docker image
docker compose -p broadband-speed-monitor -f ../docker/docker-compose.yaml up -d