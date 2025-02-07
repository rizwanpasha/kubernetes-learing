#!/bin/bash

# kubectl apply -f gin-server-clusterip-service.yaml
# kubectl apply -f gin-server-pod.yaml
# kubectl apply -f ping-pod.yaml
kubectl apply -f gin-server-loadbalancer-service.yaml

# kubectl delete pod gin-server-pod
# kubectl delete pod ping-pod
# kubectl delete service gin-server-clusterip-service
kubectl delete service gin-server-loadbalancer-service