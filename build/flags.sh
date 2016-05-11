#!/usr/bin/env sh

set -eu

cd "$(dirname "${0}")/.."

echo '-X "github.com/abcum/caboodle/util/build.rev='$(git rev-parse HEAD)'"' \
     '-X "github.com/abcum/caboodle/util/build.time='$(date -u '+%Y/%m/%d %H:%M:%S')'"'
