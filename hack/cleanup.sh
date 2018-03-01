#!/usr/bin/env bash

# delete service for operator
kubectl delete service kubecar-operator-svc

# delete operator
kubectl delete deployment kubecar-operator
# delete kubecars
kubectl delete kubecar kubecar1

# delete crd
kubectl delete crd kubecars.kubecar.emruz.com