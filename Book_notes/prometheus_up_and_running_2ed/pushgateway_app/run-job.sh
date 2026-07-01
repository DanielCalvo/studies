#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
compose_dir="${script_dir}/../prometheus_config"

cd "${compose_dir}"

docker compose run --rm --no-deps pushgateway_app "$@"
