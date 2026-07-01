#!/usr/bin/env bash
set -euo pipefail

REGISTRY="${REGISTRY:-192.168.1.242:5000}"
SOURCE_IMAGE="${SOURCE_IMAGE:-busybox:latest}"
TARGET_IMAGE="${TARGET_IMAGE:-${REGISTRY}/busybox:latest}"
PLATFORM="${PLATFORM:-linux/arm64}"
POD_NAME="${POD_NAME:-busybox-registry-test}"
NAMESPACE="${NAMESPACE:-default}"


docker buildx version >/dev/null 2>&1 || { echo "Missing Docker plugin: buildx" >&2; exit 1; }

echo "Checking registry endpoint: http://${REGISTRY}/v2/"
curl -fsS "http://${REGISTRY}/v2/" >/dev/null

echo "Pushing ${TARGET_IMAGE} for ${PLATFORM}"
printf '%s\n' \
  'ARG SOURCE_IMAGE=busybox:latest' \
  'FROM ${SOURCE_IMAGE}' \
  | docker buildx build \
      --platform "${PLATFORM}" \
      --build-arg "SOURCE_IMAGE=${SOURCE_IMAGE}" \
      -t "${TARGET_IMAGE}" \
      --push \
      -

echo "Tags:"
curl -fsS "http://${REGISTRY}/v2/busybox/tags/list"
echo

echo "Testing pull with pod ${NAMESPACE}/${POD_NAME}"
kubectl -n "${NAMESPACE}" delete pod "${POD_NAME}" --ignore-not-found --wait=true

cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: ${POD_NAME}
  namespace: ${NAMESPACE}
  labels:
    app: busybox-registry-test
spec:
  restartPolicy: Never
  containers:
    - name: busybox
      image: ${TARGET_IMAGE}
      imagePullPolicy: Always
      command:
        - sh
        - -c
        - echo "pulled ${TARGET_IMAGE}"; uname -m; sleep 3600
EOF

kubectl -n "${NAMESPACE}" wait --for=condition=Ready "pod/${POD_NAME}" --timeout=180s
kubectl -n "${NAMESPACE}" get pod "${POD_NAME}" -o wide
kubectl -n "${NAMESPACE}" logs "${POD_NAME}"
echo "Delete with: kubectl -n ${NAMESPACE} delete pod ${POD_NAME}"
