#!/usr/bin/env bash
set -euo pipefail

# gepeto built this too open ended, i think you dont need half of the things in here, refine later if you need to, it works for now

readonly SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
readonly NAMESPACE="image-resizer"
readonly DEPLOYMENT="image-resizer-api"
readonly SERVICE="image-resizer-api"
readonly PLATFORM="linux/arm64"
readonly REGISTRY="192.168.1.242:5000"
readonly LOAD_BALANCER_IP="192.168.1.222"
readonly DEPLOYMENT_TEMPLATE="${SCRIPT_DIR}/k8s/deployment.yaml"

for command in awk curl date docker kubectl mktemp sed; do
  command -v "${command}" >/dev/null 2>&1 || {
    echo "Missing required command: ${command}" >&2
    exit 1
  }
done

docker buildx version >/dev/null 2>&1 || {
  echo "Missing Docker plugin: buildx" >&2
  exit 1
}

readonly IMAGE_TAG="$(date '+v%Y-%m-%d-%H-%M-%S')"
readonly IMAGE="${REGISTRY}/image-resizer-api:${IMAGE_TAG}"
readonly RENDERED_DEPLOYMENT="$(mktemp /tmp/image-resizer-deployment.XXXXXX.yaml)"

cleanup() {
  rm -f -- "${RENDERED_DEPLOYMENT}"
}
trap cleanup EXIT

readonly TEMPLATE_IMAGE="$(awk '$1 == "image:" { print $2; exit }' "${DEPLOYMENT_TEMPLATE}")"
if [[ "${TEMPLATE_IMAGE}" != "${REGISTRY}/image-resizer-api:REPLACE_WITH_TAG" ]]; then
  echo "Expected REPLACE_WITH_TAG in ${DEPLOYMENT_TEMPLATE}" >&2
  exit 1
fi

sed "s|REPLACE_WITH_TAG|${IMAGE_TAG}|" \
  "${DEPLOYMENT_TEMPLATE}" >"${RENDERED_DEPLOYMENT}"

readonly RENDERED_IMAGE="$(awk '$1 == "image:" { print $2; exit }' "${RENDERED_DEPLOYMENT}")"
if [[ "${RENDERED_IMAGE}" != "${IMAGE}" ]]; then
  echo "Could not render ${IMAGE} into the Deployment manifest" >&2
  exit 1
fi

echo "Checking Docker and registry access"
docker info >/dev/null
curl -fsS "http://${REGISTRY}/v2/" >/dev/null

echo "Building and pushing ${IMAGE} for ${PLATFORM}"
docker buildx build \
  --platform "${PLATFORM}" \
  --tag "${IMAGE}" \
  --progress plain \
  --push \
  "${SCRIPT_DIR}"

echo "Applying Kubernetes manifests"
kubectl apply -f "${SCRIPT_DIR}/k8s/namespace.yaml"
kubectl apply --dry-run=client -f "${RENDERED_DEPLOYMENT}" >/dev/null
kubectl apply -f "${RENDERED_DEPLOYMENT}"
kubectl apply -f "${SCRIPT_DIR}/k8s/service.yaml"
kubectl apply -f "${SCRIPT_DIR}/k8s/service-monitor.yaml"

echo "Waiting for deployment rollout"
kubectl -n "${NAMESPACE}" rollout status \
  "deployment/${DEPLOYMENT}" \
  --timeout=180s

echo "Waiting for MetalLB address ${LOAD_BALANCER_IP}"
kubectl -n "${NAMESPACE}" wait \
  --for="jsonpath={.status.loadBalancer.ingress[0].ip}=${LOAD_BALANCER_IP}" \
  "service/${SERVICE}" \
  --timeout=120s

echo "Checking externally exposed endpoints"
curl -fsS "http://${LOAD_BALANCER_IP}/livez"
curl -fsS "http://${LOAD_BALANCER_IP}/readyz"
curl -fsS "http://${LOAD_BALANCER_IP}/metrics" >/dev/null

kubectl -n "${NAMESPACE}" get deployment,pods,service -o wide

echo
echo "Deployment complete: http://${LOAD_BALANCER_IP}"
echo "Deployed image: ${IMAGE}"
