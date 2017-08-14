#!/bin/sh

exec /sbin/setuser gouser go-test-task --address :8080
