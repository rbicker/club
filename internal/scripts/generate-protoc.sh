#!/bin/sh
if ! [ -x "$(command -v protoc)" ]; then
  echo 'Error: protoc is not installed.' >&2
  exit 1
fi

# GO
protoc --proto_path=api/proto/v1 --go_out=plugins=grpc:api/proto/v1 club_service.proto

# TS
protoc --proto_path=api/proto/v1 --js_out=import_style=commonjs:web \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:web club_service.proto
