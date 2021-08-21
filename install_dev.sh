#!/bin/bash
bash uninstall.sh
ip=${1:-192.168.179.128}
port=${2:-8088}
export addr=$ip:$port

bash deployments/create_cert.sh --ip $ip
#cat deployments/dev/mutatingwebhook.yaml | deployments/webhook-patch-ca-bundle.sh | kubectl apply -f -
cat deployments/dev/mutatingwebhook.yaml | deployments/dev/webhook-patch-ca-bundle.sh | kubectl apply -f -
cat deployments/dev/validatingwebhook.yaml | deployments/dev/webhook-patch-ca-bundle.sh | kubectl apply -f -

