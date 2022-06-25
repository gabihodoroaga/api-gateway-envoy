#!/bin/sh

set -e

if [ -n "$ENVOY_LOG_LEVEL" ]; then
  ENVOY_ARGS="$ENVOY_ARGS -l $ENVOY_LOG_LEVEL"
fi

echo 'generate envoy config...'

/app/envtemplate < /app/envoy-template.yaml > /app/envoy.yaml
echo 'done'

echo 'starting envoy...'
/usr/local/bin/envoy -c /app/envoy.yaml $ENVOY_ARGS
