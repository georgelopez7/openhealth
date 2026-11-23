#!/bin/sh
# EXIT ON ANY ERROR OR UNSET VARIABLE
set -eu

# CONFIGURATION - READ FROM ENVIRONMENT WITH SENSIBLE DEFAULTS
API_URL="http://openhealth-api:8000"
API_AUTH_TOKEN="${API_AUTH_TOKEN:-}"

# SEED ONLY ONCE - IF ACCOUNTS ALREADY EXIST, EXIT WITHOUT DUPLICATING DATA
if hurl --variable "API_AUTH_TOKEN=$API_AUTH_TOKEN" /check.hurl >/tmp/check.out 2>&1; then
  echo "no accounts found - seeding"
else
  echo "seed check failed - not seeding (accounts may already exist, or the API check errored)"
  echo "--- check output ---"
  cat /tmp/check.out
  echo "--------------------"
  exit 1
fi

# RUN SEED - REWRITE THE LOCALHOST URL TO THE IN-STACK API AND EXECUTE
sed "s|http://localhost:8000|$API_URL|g" /seed.hurl > /tmp/seed.hurl
hurl --variable "API_AUTH_TOKEN=$API_AUTH_TOKEN" /tmp/seed.hurl

echo "✨ Seed Completed Successfully"
