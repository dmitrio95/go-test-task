#!/bin/bash
set -e

go test
go build -o docker/go-test-task

cd docker/
docker build -t go-test-task-image .
