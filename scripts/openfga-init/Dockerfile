FROM openfga/cli:v0.7.20 AS fga

FROM alpine:latest
RUN apk add --no-cache curl jq
COPY --from=fga /fga /usr/local/bin/fga
COPY _openfga/openhealth.fga /openhealth.fga
COPY scripts/openfga-init/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
CMD ["/entrypoint.sh"]
