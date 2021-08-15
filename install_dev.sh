#!/bin/bash
bash uninstall.sh
ip=${1:-192.168.179.128}
bash deployments/create_cert.sh --ip $ip
#cat deployments/dev/mutatingwebhook.yaml | deployments/webhook-patch-ca-bundle.sh | kubectl apply -f -
cat deployments/dev/mutatingwebhook.yaml | deployments/webhook-patch-ca-bundle.sh | kubectl apply -f -
