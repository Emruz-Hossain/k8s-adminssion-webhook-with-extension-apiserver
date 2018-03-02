#!/usr/bin/env bash

set -x

# delete service for operator
kubectl delete service kubecar-operator

# delete operator
kubectl delete deployment kubecar-operator

# delete kubecars
kubectl delete kubecar kubecar1

# delete crd
kubectl delete crd kubecars.kubecar.emruz.com

# delete webhook configuration
kubectl delete validatingwebhookconfiguration validation.kubecar.emruz.com

# delete APIservice
kubectl delete apiservice v1alpha1.admission.kuebecar.emruz.com

# remove generated files
rm kbc
pushd ./hack/
rm ca.crt ca.key server.crt server.key
popd