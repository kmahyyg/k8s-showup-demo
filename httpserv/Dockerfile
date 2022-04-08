FROM gcr.io/distroless/static-debian11
LABEL org.opencontainers.image.source=https://github.com/kmahyyg/k8s-showup-demo
COPY bin/server_linux_amd64 /app/server
EXPOSE 8999
ENTRYPOINT ["/app/server"]
