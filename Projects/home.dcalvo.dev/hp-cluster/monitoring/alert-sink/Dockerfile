FROM scratch

COPY alert-sink /alert-sink

USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/alert-sink"]
