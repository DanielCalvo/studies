#!/usr/bin/env bash
set -euo pipefail

readonly SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
readonly CONTEXT="opi"
readonly NAMESPACE="image-resizer"
readonly DEPLOYMENT="image-resizer-traffic-gen"
readonly REGISTRY="192.168.1.225:5000"
readonly IMAGE_NAME="image-resizer-traffic-gen"
readonly DEPLOYMENT_TEMPLATE="${SCRIPT_DIR}/k8s/deployment.yaml"

for command in date docker kubectl mktemp sed; do
  command -v "${command}" >/dev/null 2>&1 || { echo "Missing required command: ${command}" >&2; exit 1; }
done
docker buildx version >/dev/null 2>&1 || { echo "Missing Docker plugin: buildx" >&2; exit 1; }

readonly IMAGE_TAG="$(date '+v%Y-%m-%d-%H-%M-%S')"
readonly IMAGE="${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
readonly RENDERED_DEPLOYMENT="$(mktemp /tmp/image-resizer-traffic-gen-deployment.XXXXXX.yaml)"
cleanup() { rm -f -- "${RENDERED_DEPLOYMENT}"; }
trap cleanup EXIT

sed "s|REPLACE_WITH_TAG|${IMAGE_TAG}|" "${DEPLOYMENT_TEMPLATE}" >"${RENDERED_DEPLOYMENT}"

echo "Building and pushing ${IMAGE} for linux/arm64"
docker buildx build --platform linux/arm64 --tag "${IMAGE}" --progress plain --push "${SCRIPT_DIR}"

echo "Verifying the selected OPI cluster"
kubectl --context "${CONTEXT}" config view --minify -o jsonpath='{.clusters[0].cluster.server}{"\n"}'
kubectl --context "${CONTEXT}" get nodes -o wide

echo "Applying Kubernetes manifests"
kubectl --context "${CONTEXT}" apply -f "${SCRIPT_DIR}/k8s/namespace.yaml"
kubectl --context "${CONTEXT}" apply --dry-run=client -f "${RENDERED_DEPLOYMENT}" >/dev/null
kubectl --context "${CONTEXT}" apply -f "${RENDERED_DEPLOYMENT}"

echo "Waiting for traffic generator rollout"
kubectl --context "${CONTEXT}" -n "${NAMESPACE}" rollout status "deployment/${DEPLOYMENT}" --timeout=180s
kubectl --context "${CONTEXT}" -n "${NAMESPACE}" get deployment,pods -l app.kubernetes.io/name="${DEPLOYMENT}" -o wide
