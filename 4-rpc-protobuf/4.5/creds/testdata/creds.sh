#!/bin/bash

# 证书认证
# 服务端私钥 & 证书
openssl genrsa -out server.key 2048
openssl req -new -x509 -days 3650 \
	-subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" \
	-key server.key -out server.crt
# 客户端私钥 & 证书
openssl genrsa -out client.key 2048
openssl req -new -x509 -days 3650 \
	-subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" \
	-key client.key -out client.crt