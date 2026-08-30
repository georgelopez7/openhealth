#!/bin/sh
# EXIT ON ANY ERROR OR UNSET VARIABLE
set -eu

# CONFIGURATION - READ FROM ENVIRONMENT WITH SENSIBLE DEFAULTS
FGA_API_URL="${FGA_API_URL:-http://openfga:8080}"
STORE_NAME="${FGA_STORE_NAME:-openhealth}"

# RESOLVE STORE - PREFER FGA_STORE_ID IF PROVIDED, OTHERWISE LOOK UP BY NAME
STORE_ID="${FGA_STORE_ID:-}"
if [ -z "$STORE_ID" ]; then
  STORE_ID=$(curl -sf "$FGA_API_URL/stores?pageSize=100" \
    | jq -r --arg n "$STORE_NAME" '.stores[]? | select(.name == $n) | .id' \
    | head -n1)
fi

# CREATE STORE IF NONE EXISTS - OTHERWISE REUSE THE EXISTING ONE
if [ -z "$STORE_ID" ]; then
  STORE_ID=$(curl -sf -X POST "$FGA_API_URL/stores" \
    -H "Content-Type: application/json" \
    -d "{\"name\":\"$STORE_NAME\"}" | jq -r '.id')
  echo "created store '$STORE_NAME' ($STORE_ID)"
else
  echo "using existing store '$STORE_NAME' ($STORE_ID)"
fi

# CHECK FOR EXISTING MODEL - FETCH THE LATEST AUTHORIZATION MODEL IN THE STORE
MODEL_ID=$(curl -sf "$FGA_API_URL/stores/$STORE_ID/authorization-models?pageSize=1" \
  | jq -r '.authorization_models[0].id // empty')

# WRITE MODEL ONLY IF THE STORE HAS NONE - KEEPS REDEPLOYS IDEMPOTENT
if [ -n "$MODEL_ID" ]; then
  echo "store already has a model ($MODEL_ID) - skipping model write"
else
  MODEL_ID=$(FGA_API_URL="$FGA_API_URL" FGA_STORE_ID="$STORE_ID" \
    fga model write --file=/openhealth.fga \
    | jq -r '.authorization_model.id // .id')
  echo "wrote model ($MODEL_ID)"
fi

echo ""
echo "=============================================================="
echo " OpenFGA Configuration"
echo ""
echo "  FGA_STORE_ID=$STORE_ID"
echo "  FGA_MODEL_ID=$MODEL_ID"
echo "=============================================================="
