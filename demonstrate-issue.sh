#!/bin/bash

echo "first some setup: we create a directory named github.com that even the current user is not allowed to write to"
echo "the github.com directory would be automatically created by go-snaps under -trimpath"
set -x
install -d -m 555 github.com
set +x

echo
echo
echo "now we prove, that the test work perfectly, if GOFLAGS is empty"
echo
set -x
export GOFLAGS=""
go test -count=1 -v ./...
set +x

echo
echo
echo "and now we can prove, that if we add -trimpath to GOFLAGS, the tests start to break"
echo
set -x
export GOFLAGS="-trimpath"
go test -count=1 -v ./...
set +x
