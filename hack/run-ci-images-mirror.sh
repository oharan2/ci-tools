#!/usr/bin/env bash

set -euo pipefail

TMP_DIR="$(mktemp -d)"

trap 'rm -rf ${TMP_DIR}' EXIT
oc --context app.ci  -n ci extract secret/ci-images-mirror --to="${TMP_DIR}"

release="${RELEASE:-"$(go env GOPATH)/src/github.com/openshift/release"}"

set -x
KUBECONFIG="${TMP_DIR}/sa.ci-images-mirror.app.ci.config" go run  ./cmd/ci-images-mirror \
  --leader-election-namespace=ci \
  --leader-election-suffix="-${USER}" \
  --release-repo-git-sync-path="${release}"  \
  --quayIOCIImagesDistributorOptions.additional-image-stream-namespace=ci \
  --dry-run=true
