FROM gcr.io/distroless/static-debian11
COPY bin/server_linux_amd64 /app/server
ENTRYPOINT ["/app/server"]