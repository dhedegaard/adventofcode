#!/bin/sh

for f in day*.go; do
    echo "--- RUNNING $f ---"
    go run $f
done
