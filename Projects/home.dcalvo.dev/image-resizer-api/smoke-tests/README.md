# Post-deployment smoke tests

This suite performs a small set of functional checks against a deployed image
resizer. It covers the essential success path, representative input rejection
paths, operational endpoints, response dimensions, and request-ID behavior.

Run it against the homelab deployment with:

```bash
./smoke-tests/post-deployment-smoke-test.sh
```

Pass another base URL when testing a different environment:

```bash
./smoke-tests/post-deployment-smoke-test.sh http://image-resizer.example.test
```

The script generates all fixtures under `test-data/`, compiles a temporary helper
with the local Go toolchain, and removes only its temporary response files and
helper binary when it exits.

The suite intentionally remains small. It is a post-deployment functional smoke
test, not an exhaustive API regression suite.
