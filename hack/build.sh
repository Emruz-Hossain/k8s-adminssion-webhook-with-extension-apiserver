#!/usr/bin/env bash

# build binary
go build -o kbc ./main.go

# build docker image
docker build -t emruzhossain/kbc .

# load docker image to minikube
docker save emruzhossain/kbc | pv | (eval $(minikube docker-env) && docker load)

# create crd
kubectl apply -f ./hack/yamls/crd-kc.yaml

# create kubecar
kubectl apply -f ./hack/yamls/kubecar.yaml

# deploy operator
kubectl apply -f ./hack/yamls/operator.yaml

