FROM ghcr.io/orange-opensource/hurl:latest
COPY _hurl/seed.hurl /seed.hurl
COPY scripts/seed-db/check.hurl /check.hurl
COPY scripts/seed-db/entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]
