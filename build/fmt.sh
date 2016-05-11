#!/usr/bin/env sh

total=`git ls-files | grep '.go$' | xargs gofmt -l 2>&1 | wc -l`

if [ $total -gt 0 ]; then
    echo "run 'go fmt ./...' to format your source code."
    exit 1
fi