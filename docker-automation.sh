#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0
cd accountservice;go get;go build -o accountservice-linux-amd64;echo built `pwd`;cd ..
export GOOS=darwin
docker build -t samratjha96/accountservice .
docker service rm accountservice
docker service create --name=accountservice --replicas=1 --network=service-mesh -p=6767:6767 samratjha96/accountservice