#!/usr/bin/env bash

set -eax

DOCS_IMAGE_TAG="aws-sdk-go-v2-docs"
DOCS_BUILD_CONTAINER_NAME="aws-sdk-go-v2-docs-build"
SITE_ENV=${1:-production}

rm -rf ./docs

docker build -t $DOCS_IMAGE_TAG --build-arg SITE_ENV="${SITE_ENV}" .
docker container create --name $DOCS_BUILD_CONTAINER_NAME $DOCS_IMAGE_TAG
docker container cp $DOCS_BUILD_CONTAINER_NAME:/aws-sdk-go-v2-docs/docs .
docker container rm -f $DOCS_BUILD_CONTAINER_NAME
