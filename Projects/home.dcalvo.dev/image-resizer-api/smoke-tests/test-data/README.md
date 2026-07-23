# Smoke-test data

`post-deployment-smoke-test.sh` generates all input fixtures in this directory so
the suite does not depend on personal images or external downloads.

Generated inputs include:

- A valid `120x80` JPEG
- A valid PNG
- A truncated, corrupt JPEG
- A JPEG header declaring more than 40 million pixels
- A JPEG upload larger than 10 MiB
- A plain-text file
- A minimal MP4 container header

Generated binary and text fixtures are ignored by Git and can be recreated by
running the smoke-test script.
