#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"

test "$(kubectl --context hp config view --minify \
  -o jsonpath='{.clusters[0].cluster.server}')" = "https://192.168.1.211:6443"
kubectl --context hp get nodes hp1 hp2 -o wide

cleanup() {
  kubectl --context hp delete -f smoke-test.yaml --ignore-not-found
}
trap cleanup EXIT

kubectl --context hp apply -f smoke-test.yaml
kubectl --context hp -n local-path-storage wait \
  --for=condition=Ready pod/local-path-smoke-test --timeout=120s

kubectl --context hp -n local-path-storage exec local-path-smoke-test -- \
  sh -c 'echo local-path-persistence-test > /data/marker'
first_node="$(kubectl --context hp -n local-path-storage get pod \
  local-path-smoke-test -o jsonpath='{.spec.nodeName}')"

kubectl --context hp -n local-path-storage delete pod local-path-smoke-test --wait
kubectl --context hp apply -f smoke-test.yaml
kubectl --context hp -n local-path-storage wait \
  --for=condition=Ready pod/local-path-smoke-test --timeout=120s

test "$(kubectl --context hp -n local-path-storage exec \
  local-path-smoke-test -- cat /data/marker)" = "local-path-persistence-test"
test "$(kubectl --context hp -n local-path-storage get pod \
  local-path-smoke-test -o jsonpath='{.spec.nodeName}')" = "$first_node"

echo "Persistence test passed on $first_node."
