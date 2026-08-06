#!/usr/bin/env bash

set -euo pipefail

dashboard_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
dashboard_file="$dashboard_dir/image-resizer-overview.json"
grafana_url="${GRAFANA_URL:-http://192.168.1.231}"
grafana_user="${GRAFANA_USER:-admin}"
grafana_password="${GRAFANA_PASSWORD:-admin}"
dashboard_uid="image-resizer-overview"

usage() {
  echo "Usage: $0 pull|push" >&2
  exit 2
}

case "${1:-}" in
  pull)
    curl --fail-with-body --silent --show-error \
      --user "$grafana_user:$grafana_password" \
      "$grafana_url/api/dashboards/uid/$dashboard_uid" \
      | jq '{dashboard: .dashboard, overwrite: true}' > "$dashboard_file"
    echo "Pulled $dashboard_uid into $dashboard_file"
    ;;
  push)
    curl --fail-with-body --silent --show-error \
      --user "$grafana_user:$grafana_password" \
      --header 'Content-Type: application/json' \
      --data-binary "@$dashboard_file" \
      "$grafana_url/api/dashboards/db"
    ;;
  *)
    usage
    ;;
esac
