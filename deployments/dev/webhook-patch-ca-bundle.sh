#!/bin/bash

ROOT=$(cd $(dirname $0)/../../; pwd)

set -o errexit
set -o nounset
set -o pipefail



#echo "addr---:$addr"
#export CA_BUNDLE=$(kubectl config view --raw --flatten -o json | jq -r '.clusters[] | select(.name == "'$(kubectl config current-context)'") | .cluster."certificate-authority-data"')
export CA_BUNDLE=$(kubectl config view --raw --flatten -o json | jq -r '.clusters[] | select(.name == "cluster.local") | .cluster."certificate-authority-data"')
if command -v envsubst >/dev/null 2>&1; then
    envsubst
else

    sed -e "s|\${CA_BUNDLE}|${CA_BUNDLE}|g"
    sed -e "s|\${addr}|${addr}|g"
fi
