#!/usr/bin/env bash
set -e

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
test_data_dir="${script_dir}/test-data"
base_url="http://192.168.1.222"
sleep_time=0.1

echo "Generating traffic fixtures in ${test_data_dir}"
/usr/local/go/bin/go run "${script_dir}/generate-test-data.go" "${test_data_dir}"

# Build the JPEG list once, then shuffle it once for this run.
mapfile -t images < <(printf '%s\n' "${test_data_dir}"/*.jpg | shuf)

echo "Generating sequential traffic against ${base_url}; press Ctrl+C to stop"
# --fail stops on an HTTP error, --output discards the resized JPEG, and --form
# uploads the source JPEG in the multipart field named "image".
while true; do
  for image in "${images[@]}"; do
    filename="${image##*/}"
    dimensions="${filename%.jpg}"
    input_width="${dimensions%%x*}"
    input_height="${dimensions#*x}"
    output_width=$((input_width / 2))
    output_height=$((input_height / 2))

    echo "Resizing ${filename} from ${input_width}x${input_height} to ${output_width}x${output_height}"
    curl --fail --output /dev/null \
      --form "image=@${image}" \
      "${base_url}/v1/resize?width=${output_width}"

    echo "sleeping ${sleep_time}"
    sleep "${sleep_time}"
    echo
  done
done
