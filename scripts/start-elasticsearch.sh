#!/bin/sh

: "${ELASTICSEARCH_PATH:?Need to set ELASTICSEARCH_PATH non-empty}"

${ELASTICSEARCH_PATH}/bin/elasticsearch
