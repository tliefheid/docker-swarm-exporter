#!/usr/bin/env bash
docker service create --limit-cpu 0.5 --limit-memory 10m --reserve-cpu 0.2 --reserve-memory 5m --mode replicated --replicas 1 --name nginx1 nginx:latest 
docker service create --limit-cpu 0.4 --limit-memory 20m --reserve-cpu 0.1 --reserve-memory 4m --mode global --name nginx2 nginx:latest 