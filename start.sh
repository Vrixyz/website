#!/bin/sh

: "${WEBSITE_PATH:?Need to set WEBSITE_ROOT non-empty}"

#FIXME: logstash needs conf file as parameter
( \
  ${WEBSITE_PATH}/scripts/start-custom.sh -d elasticsearch && \
  ${WEBSITE_PATH}/scripts/start-custom.sh -d logstash && \
  ${WEBSITE_PATH}/scripts/start-custom.sh -d kibana && \
  screen -d -m -S webserver ${WEBSITE_PATH}/scripts/start-webserver.sh && \
  echo "started screen webserver. All services launched.") \
|| echo "Something failed.";
