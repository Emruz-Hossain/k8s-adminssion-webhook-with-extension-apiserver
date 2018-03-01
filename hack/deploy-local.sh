#!/usr/bin/env bash

# build binary
go build -o kbc ./main.go

# create crd
kubectl apply -f ./hack/yamls/crd-kc.yaml

# create kubecar
kubectl apply -f ./hack/yamls/kubecar.yaml

# run controller
./kbc --kubeconfig=/home/ac/.kube/config

