FROM golang

#Install the required tools as mentioned by kelsey


#You need to do the go get thing!


RUN curl -o cfssl https://pkg.cfssl.org/R1.2/cfssl_darwin-amd64 && \
    curl -o cfssljson https://pkg.cfssl.org/R1.2/cfssljson_darwin-amd64 && \
    chmod +x cfssl cfssljson && \
    mv cfssl cfssljson /usr/local/bin/

RUN wget -q --show-progress --https-only --timestamping \
    https://pkg.cfssl.org/R1.2/cfssl_linux-amd64 \
    https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64 && \
    chmod +x cfssl_linux-amd64 cfssljson_linux-amd64 && \
    mv cfssl_linux-amd64 /usr/local/bin/cfssl && \
    mv cfssljson_linux-amd64 /usr/local/bin/cfssljson

RUN wget https://storage.googleapis.com/kubernetes-release/release/v1.12.0/bin/linux/amd64/kubectl && \
    chmod +x kubectl && \
    mv kubectl /usr/local/bin/

COPY . .

CMD ["echo", "Hello", "World!"]
#CMD ["go","run","certgen.go"]