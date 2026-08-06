#!/usr/bin/env bash
set -euo pipefail

readonly SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
readonly CONTEXT="hp"
readonly NAMESPACE="monitoring"
readonly DEPLOYMENT="alert-sink"
readonly REGISTRY="192.168.1.225:5000"
readonly IMAGE_NAME="alert-sink"
readonly DEPLOYMENT_TEMPLATE="${SCRIPT_DIR}/k8s/deployment.yaml"
readonly BINARY="${SCRIPT_DIR}/alert-sink"

for command in date docker kubectl mktemp sed; do
  command -v "${command}" >/dev/null 2>&1 || { echo "Missing required command: ${command}" >&2; exit 1; }
done

readonly IMAGE_TAG="$(date '+v%Y-%m-%d-%H-%M-%S')"
readonly IMAGE="${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
readonly RENDERED_DEPLOYMENT="$(mktemp /tmp/alert-sink-deployment.XXXXXX.yaml)"

cleanup() {
  rm -f -- "${BINARY}" "${RENDERED_DEPLOYMENT}"
}
trap cleanup EXIT

echo "Building ${IMAGE} for linux/amd64"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -trimpath -ldflags="-s -w" -o "${BINARY}" "${SCRIPT_DIR}"
docker build --platform linux/amd64 --tag "${IMAGE}" "${SCRIPT_DIR}"
docker push "${IMAGE}"

sed "s|REPLACE_WITH_TAG|${IMAGE_TAG}|" "${DEPLOYMENT_TEMPLATE}" >"${RENDERED_DEPLOYMENT}"

echo "Verifying the selected HP cluster"
kubectl --context "${CONTEXT}" config view --minify -o jsonpath='{.clusters[0].cluster.server}{"\n"}'
kubectl --context "${CONTEXT}" get nodes -o wide

echo "Applying alert sink manifests"
kubectl --context "${CONTEXT}" apply --dry-run=client -f "${RENDERED_DEPLOYMENT}" >/dev/null
kubectl --context "${CONTEXT}" apply -f "${RENDERED_DEPLOYMENT}"
kubectl --context "${CONTEXT}" apply -f "${SCRIPT_DIR}/k8s/service.yaml"
kubectl --context "${CONTEXT}" -n "${NAMESPACE}" rollout status "deployment/${DEPLOYMENT}" --timeout=180s
kubectl --context "${CONTEXT}" -n "${NAMESPACE}" get deployment,pods,service -l app.kubernetes.io/name="${DEPLOYMENT}" -o wide

echo "Deployed image: ${IMAGE}"
