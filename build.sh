#!/bin/bash
docker buildx create --use --name bot_adapter
docker buildx build --tag scjtqs/bot_adapter:latest --platform linux/amd64,linux/arm64,linux/386,linux/arm/v7 --push  .
docker buildx rm bot_adapter