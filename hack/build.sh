#!/usr/bin/env bash

set -xe

# build binary
go build -o kbc ./main.go

# build docker image
docker build -t emruzhossain/kbc .

# load docker image to minikube
docker save emruzhossain/kbc | pv | (eval $(minikube docker-env) && docker load)

pushd ./hack

# create necessary TLS certificates
./onessl create ca-cert
./onessl create server-cert server --domains=kubecar-operator.default.svc
export SERVICE_SERVING_CERT_CA=$(cat ca.crt | ./onessl base64)
export TLS_SERVING_CERT=$(cat server.crt | ./onessl base64)
export TLS_SERVING_KEY=$(cat server.key | ./onessl base64)
export KUBE_CA=$(./onessl get kube-ca | ./onessl base64)

# create crd
cat ./yamls/crd-kc.yaml | ./onessl envsubst | kubectl apply -f -

# deploy operator
cat ./yamls/operator.yaml  | ./onessl envsubst | kubectl apply -f -

# create Webhook configurations
cat ./yamls/admission.yaml | ./onessl envsubst | kubectl apply -f -

# create API service
cat ./yamls/apiserver.yaml | ./onessl envsubst | kubectl apply -f -


popd