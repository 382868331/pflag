#!/usr/bin/env bash
set -euo pipefail
IMAGE_NAME=${1:-pflag-task}
DOCKER_PLATFORM=${2:-linux/amd64}
DOCKER_BUILDKIT=1 docker build --platform "$DOCKER_PLATFORM" -f goletalab.Dockerfile -t "$IMAGE_NAME" .
