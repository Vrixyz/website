#!/bin/sh

: "${KIBANA_PATH:?Need to set KIBANA_PATH non-empty}"

${KIBANA_PATH}/bin/kibana
