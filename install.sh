#!/bin/bash
bash uninstall.sh
namespace=kube-system

bash deployments/create_cert.sh --namespace ${namespace}
#cat deployments/mutatingwebhook.yaml | deployments/webhook-patch-ca-bundle.sh | kubectl apply -f -
cat deployments/mutatingwebhook.yaml | deployments/webhook-patch-ca-bundle.sh | kubectl apply  -n ${namespace} -f -
cat deployments/validatingwebhook.yaml | deployments/webhook-patch-ca-bundle.sh | kubectl apply  -n ${namespace} -f -
kubectl apply -f deployments/rbac.yaml -n ${namespace}
kubectl apply -f deployments/service.yaml -n ${namespace}
kubectl apply -f deployments/deployment.yaml -n ${namespace}


