#!/bin/sh
hostip=`docker-machine ip dev`

wrk -t12 -c400 -d30s http://${hostip}:1323/
wrk -t12 -c400 -d30s http://${hostip}:1324/
wrk -t12 -c400 -d30s http://${hostip}:1325/
