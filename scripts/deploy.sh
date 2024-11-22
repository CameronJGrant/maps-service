#!/bin/bash

echo "Deploying ${IMAGE}:${GO_PIPELINE_LABEL} as ${APP}"

helm upgrade --install "${APP}" --namespace ${NAMESPACE} strafesnet/service-generic \
  --set "image.repository=${IMAGE}" \
  --set "image.tag=${GO_PIPELINE_LABEL}" \
  --set "service.grpc=9000" \
  --set "config.postgres=pgo-cluster-pguser-game-data" \
  --set "config.cron=*/15 * * * *"