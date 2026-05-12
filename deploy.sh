#!/bin/bash

docker pull dengrenjie31/status
docker compose up -d
docker restart nginx
