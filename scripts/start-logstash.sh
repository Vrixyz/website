#!/bin/sh

: "${LOGSTASH_PATH:?Need to set LOGSTASH_PATH non-empty}"
: "${WEBSITE_PATH:?Need to set WEBSITE_PATH non-empty}"

${LOGSTASH_PATH}/bin/logstash -f ${WEBSITE_PATH}/logstash.conf
