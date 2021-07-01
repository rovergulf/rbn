#!/usr/bin/env bash

set -e

source ./scripts/util.sh

log_callout log_info "[$(date)] generating proto files"
protoc "-I.:.." --proto_path=rpc --go_out=plugins=grpc:. \
    --go_opt=paths=source_relative ./rpc/*.proto || exit 1

log_success "[$(date)] successfully generated pb files"