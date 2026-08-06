#!/usr/bin/env bash

set -euo pipefail

BRUNO_DIR="./docs/_bruno"

rm -rf "$BRUNO_DIR/_tmp"

bru import openapi \
  --source ./docs/openapi.yaml \
  --output "$BRUNO_DIR/_tmp" \
  --collection-name "openhealth-api" \
  > /dev/null

rm -rf "$BRUNO_DIR/_tmp/environments"
rsync -a --delete --exclude 'environments/' "$BRUNO_DIR/_tmp/" "$BRUNO_DIR/openhealth-api/"
rm -rf "$BRUNO_DIR/_tmp"

find "$BRUNO_DIR/openhealth-api" -name "*.yml" \
  -exec sed -i '' \
    -e 's/{{baseUrl}}/{{BASE_URL}}/g' \
    -e 's/name: baseUrl/name: BASE_URL/g' {} +

printf "\033[0;32m> Bruno Collection Updated\033[0m\n"
