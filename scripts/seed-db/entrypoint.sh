#!/bin/sh
# EXIT ON ANY ERROR OR UNSET VARIABLE
set -eu

# CONFIGURATION - READ FROM ENVIRONMENT WITH SENSIBLE DEFAULTS
API_URL="http://openhealth-api:8000"
API_AUTH_TOKEN="${API_AUTH_TOKEN:-}"

# GATE - SKIP SEEDING UNTIL THE FGA STORE AND MODEL ARE DEFINED FOR THE API
if [ -z "${FGA_STORE_ID:-}" ] || [ -z "${FGA_MODEL_ID:-}" ]; then
  echo "FGA_STORE_ID/FGA_MODEL_ID not set - skipping seed"
  exit 0
fi

# SEED ONLY ONCE - IF ACCOUNTS ALREADY EXIST, EXIT WITHOUT DUPLICATING DATA
if hurl --variable "API_AUTH_TOKEN=$API_AUTH_TOKEN" /check.hurl >/dev/null 2>&1; then
  echo "no accounts found - seeding"
else
  echo "accounts already exist - skipping seed"
  exit 0
fi

# RUN SEED - REWRITE THE LOCALHOST URL TO THE IN-STACK API AND EXECUTE
sed "s|http://localhost:8000|$API_URL|g" /seed.hurl > /tmp/seed.hurl
hurl --variable "API_AUTH_TOKEN=$API_AUTH_TOKEN" \
  --header "Authorization: Bearer $API_AUTH_TOKEN" \
  /tmp/seed.hurl

echo "✨ Seed Completed Successfully"
