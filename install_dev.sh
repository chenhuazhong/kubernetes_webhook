#!/bin/bash

bash deployments/create_cert.sh --ip 192.168.179.128
#cat deployments/dev/mutatingwebhook.yaml | deployments/webhook-patch-ca-bundle.sh | kubectl apply -f -
cat deployments/dev/mutatingwebhook.yaml | deployments/webhook-patch-ca-bundle.sh | kubectl apply -f -
