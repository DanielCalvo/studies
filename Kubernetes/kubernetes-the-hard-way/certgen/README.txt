This is a golang script to automate the generation of certificates for k8s the hard way.

It's a lot of certificates, doing it by hand is not fun! Also very error prone...

USAGE:
cd $YOUR_FS_PATH/certgen
docker build . -t certgen
docker run -v "$(pwd)"/certs:/go/certs certgen

The whole thing:
docker build . -t certgen && docker run -v "$(pwd)"/certs:/go/certs certgen

Troubleshooting:
docker build . -t certgen && docker run -it -v "$(pwd)"/certs:/go/certs certgen bash

