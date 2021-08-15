#!/bin/bash

namespace=kube-system

kubectl delete secret admission-webhook-example-certs -n ${namespace}
kubectl delete svc admission-webhook-example-svc -n ${namespace}
kubectl delete ServiceAccount admission-webhook-example-sa -n ${namespace}
kubectl delete ClusterRole admission-webhook-example-cr -n ${namespace}
kubectl delete ClusterRoleBinding admission-webhook-example-crb  -n ${namespace}
kubectl delete Deployment admission-webhook-example-deployment  -n ${namespace}
kubectl delete CertificateSigningRequest admission-webhook-example-svc  -n ${namespace}
kubectl delete ValidatingWebhookConfiguration validation-webhook-example-cfg  -n ${namespace}
kubectl delete MutatingWebhookConfiguration mutating-webhook-example-cfg  -n ${namespace}

