FROM alpine:latest
RUN apk add --no-cache curl jq
ARG TARGETARCH
RUN wget -qO /tmp/fga.tar.gz "https://github.com/openfga/cli/releases/latest/download/fga_linux_${TARGETARCH:-amd64}.tar.gz" \
  && tar -xzf /tmp/fga.tar.gz -C /usr/local/bin fga
COPY _openfga/openhealth.fga /openhealth.fga
COPY scripts/openfga-init/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
CMD ["/entrypoint.sh"]
