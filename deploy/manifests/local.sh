#!/bin/bash

export NAMESPACE=b
export DOMAIN=localhost

kubectl kustomize . | envsubst > output/out.yaml
kubectl apply -f output/out.yaml