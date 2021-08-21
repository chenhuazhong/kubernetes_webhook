#!/bin/bash

version=${1:-v1}

if [ -e ./dist ]
then
  echo dist is exists
else
  echo dist not exists
  mkdir dist
fi

echo version----:$version
go build  -o ./dist/kubernetes_webhook ./main.go

docker build . -t registry.example.com:5000/kubernetes_webhook:${version}
docker push  registry.example.com:5000/kubernetes_webhook:${version}
