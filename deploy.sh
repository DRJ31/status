#!/bin/bash

docker stop status
docker rm status
docker rmi dengrenjie31/status
docker compose up -d