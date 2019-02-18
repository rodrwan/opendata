#/bin/sh

kubectl create -f ./namespace.yml
kubectl create -f ./frontend/frontend.yml
kubectl create -f ./graphql/backend.yml



