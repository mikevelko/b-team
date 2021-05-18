#!/bin/sh

PROJECT_ROOT=$(git rev-parse --show-toplevel)
REPO_NAME=$(basename ${PROJECT_ROOT})
COMMIT_TAG=$(git rev-parse --short HEAD)


for service in auth-service hotel-service offer-service room-service user-service
do
  IMAGE_ID="${REGISTRY_URL}/${REPO_NAME}-${service}:${COMMIT_TAG}"
  echo "building $service with image ID=${IMAGE_ID}"
  docker build --build-arg SERVICE_NAME=${service} -t ${IMAGE_ID}  "${PROJECT_ROOT}/server"
  docker push "${IMAGE_ID}"
done