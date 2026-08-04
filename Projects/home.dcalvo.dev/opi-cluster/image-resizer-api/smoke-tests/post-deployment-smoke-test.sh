#!/usr/bin/env bash
set -euo pipefail

readonly SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
readonly TEST_DATA_DIR="${SCRIPT_DIR}/test-data"
readonly BASE_URL="${1:-${BASE_URL:-http://192.168.1.222}}"
readonly WORK_DIR="$(mktemp -d /tmp/image-resizer-smoke.XXXXXX)"

cleanup() {
  rm -rf -- "${WORK_DIR}"
}
trap cleanup EXIT

for command in awk curl mktemp; do
  command -v "${command}" >/dev/null 2>&1 || {
    echo "Missing required command: ${command}" >&2
    exit 1
  }
done

GO_BIN="${GO_BIN:-/usr/local/go/bin/go}"
if [[ ! -x "${GO_BIN}" ]]; then
  GO_BIN="$(command -v go || true)"
fi
if [[ -z "${GO_BIN}" ]]; then
  echo "Missing required command: go" >&2
  exit 1
fi
readonly GO_BIN

pass_count=0

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

pass() {
  pass_count=$((pass_count + 1))
  echo "PASS: $*"
}

http_status() {
  local output_file="$1"
  shift
  curl \
    --silent \
    --show-error \
    --connect-timeout 3 \
    --max-time 20 \
    --output "${output_file}" \
    --write-out '%{http_code}' \
    "$@"
}

assert_get() {
  local path="$1"
  local expected_status="$2"
  local output_file="${WORK_DIR}/get-${path#/}.body"
  local status

  if ! status="$(http_status "${output_file}" "${BASE_URL}${path}")"; then
    fail "GET ${path} could not reach ${BASE_URL}"
  fi
  if [[ "${status}" != "${expected_status}" ]]; then
    fail "GET ${path} returned ${status}, expected ${expected_status}; body: $(<"${output_file}")"
  fi
  pass "GET ${path} returned ${expected_status}"
}

assert_rejected_upload() {
  local name="$1"
  local fixture="$2"
  local query="$3"
  local expected_status="$4"
  local expected_error="$5"
  local output_file="${WORK_DIR}/${name}.body"
  local status

  if ! status="$(http_status \
    "${output_file}" \
    --form "image=@${TEST_DATA_DIR}/${fixture}" \
    "${BASE_URL}/v1/resize${query}")"; then
    fail "${name} request could not reach ${BASE_URL}"
  fi
  if [[ "${status}" != "${expected_status}" ]]; then
    fail "${name} returned ${status}, expected ${expected_status}; body: $(<"${output_file}")"
  fi

  local body
  body="$(<"${output_file}")"
  if [[ "${body}" != *"\"error\":\"${expected_error}\""* ]]; then
    fail "${name} did not return error code ${expected_error}; body: ${body}"
  fi
  pass "${name} returned ${expected_status} and ${expected_error}"
}

echo "Generating smoke-test data in ${TEST_DATA_DIR}"
"${GO_BIN}" build -o "${WORK_DIR}/fixture-tool" "${SCRIPT_DIR}/fixture-tool.go"
"${WORK_DIR}/fixture-tool" generate "${TEST_DATA_DIR}"

echo "Running post-deployment smoke tests against ${BASE_URL}"
assert_get "/livez" "200"
assert_get "/readyz" "200"
assert_get "/metrics" "200"

success_body="${WORK_DIR}/valid-resize.jpg"
success_headers="${WORK_DIR}/valid-resize.headers"
if ! success_status="$(curl \
  --silent \
  --show-error \
  --connect-timeout 3 \
  --max-time 20 \
  --dump-header "${success_headers}" \
  --output "${success_body}" \
  --write-out '%{http_code}' \
  --form "image=@${TEST_DATA_DIR}/valid.jpg" \
  "${BASE_URL}/v1/resize?width=60")"; then
  fail "valid JPEG request could not reach ${BASE_URL}"
fi
if [[ "${success_status}" != "200" ]]; then
  fail "valid JPEG returned ${success_status}, expected 200; body: $(<"${success_body}")"
fi
if [[ "$("${WORK_DIR}/fixture-tool" dimensions "${success_body}")" != "60x40" ]]; then
  fail "valid JPEG response did not have dimensions 60x40"
fi
if ! awk 'tolower($1) == "x-request-id:" && length($2) > 1 { found = 1 } END { exit !found }' "${success_headers}"; then
  fail "valid JPEG response did not include X-Request-ID"
fi
pass "valid JPEG was resized from 120x80 to 60x40"

assert_rejected_upload "PNG upload" "valid.png" "?width=60" "415" "unsupported_image_format"
assert_rejected_upload "corrupt JPEG" "corrupt.jpg" "?width=60" "415" "unsupported_image_format"
assert_rejected_upload "text upload" "plain-text.txt" "?width=60" "415" "unsupported_image_format"
assert_rejected_upload "MP4 upload" "video.mp4" "?width=60" "415" "unsupported_image_format"
assert_rejected_upload "oversized upload" "oversized.jpg" "?width=60" "413" "image_too_large"
assert_rejected_upload "excessive pixel count" "too-many-pixels.jpg" "?width=60" "413" "image_dimensions_too_large"
assert_rejected_upload "upscaling request" "valid.jpg" "?width=121" "422" "upscaling_not_supported"
assert_rejected_upload "zero-width request" "valid.jpg" "?width=0" "400" "invalid_width"

assert_get "/livez" "200"
assert_get "/readyz" "200"

echo
echo "Smoke tests passed: ${pass_count}"
