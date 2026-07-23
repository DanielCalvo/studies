# Traffic generator

This utility creates steady, low-rate application traffic for exploring logs,
Prometheus metrics, and dashboards. It is not a capacity or load test.

Run it against the homelab deployment:

```bash
./traffic-gen/run.sh
```

It continuously cycles through ten generated JPEG fixtures, pausing briefly
between requests:

- `640x480`, resized to `320x240`
- `800x1200`, resized to `400x600`
- `1200x1200`, resized to `600x600`
- `1600x1200`, resized to `800x600`
- `2000x1500`, resized to `1000x750`
- `2400x1800`, resized to `1200x900`
- `2000x3000`, resized to `1000x1500`
- `3200x2400`, resized to `1600x1200`
- `4000x3000`, resized to `2000x1500`
- `5000x5000`, resized to `2500x2500`

The script generates the fixtures under `test-data/`, lists the JPEGs once,
shuffles that list once per run, and then repeatedly follows the shuffled order.
Stop it with `Ctrl+C`. The cluster endpoint and pause are kept directly in the
script so their values remain obvious.
