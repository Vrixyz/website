#!/bin/sh

: "${WEBSITE_PATH:?Need to set WEBSITE_PATH non-empty}"

go run ${WEBSITE_PATH}/server.go
