#!/bin/sh

echo DAEMON    = "${DAEMON}"

: "${ELASTICESEARCH_PATH:?Need to set ELASTICESEARCH_PATH non-empty}"

${ELASTICESEARCH_PATH}/bin/elasticsearch
