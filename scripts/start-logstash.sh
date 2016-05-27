#!/bin/sh

: "${LOGSTASH_PATH:?Need to set LOGSTASH_PATH non-empty}"

${LOGSTASH_PATH}/bin/logstash -f ${LOGSTASH_PATH}/logstash.conf
