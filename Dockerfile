FROM alpine:latest

COPY bin/api /api

ENV HELLOGRPC_HOST=:9090
EXPOSE 9090

ENTRYPOINT ["/api"]