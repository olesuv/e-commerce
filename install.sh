#!/bin/sh
cd ./client; npm i --save; cd ..; go build -o server; rm server

