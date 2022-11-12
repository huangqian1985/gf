#!/usr/bin/env bash

GOARCH=${{ matrix.goarch }}
for file in `find . -name go.mod`; do
    dirpath=$(dirname $file)
    echo $dirpath

    # package oracle needs golang >= v1.17
    if [ "oracle" = $(basename $dirpath) ]; then
        if ! go version|grep -q "1.17"; then
          echo "ignore oracle as go version: $(go version)"
          continue 1
        fi
    fi

    # package kuhecm needs golang >= v1.18
    if [ "kubecm" = $(basename $dirpath) ]; then
        if ! go version|grep -q "1.18"; then
          echo "ignore kubecm as go version: $(go version)"
          continue 1
        fi
    fi

    # package example needs golang >= v1.18
    if [ "example" = $(basename $dirpath) ]; then
        if ! go version|grep -q "1.18"; then
          echo "ignore example as go version: $(go version)"
          continue 1
        fi
    fi

    # package cmd/gf needs golang >= v1.18
    if [ "gf" = $(basename $dirpath) ]; then
        if ! go version|grep -q "1.18"; then
          echo "ignore example as go version: $(go version)"
          continue 1
        fi
    fi

    cd $dirpath
    go mod tidy
    go build ./...
    go test ./... -race -coverprofile=coverage.out -covermode=atomic -coverpkg=./...,github.com/gogf/gf/... || exit 1
    if grep -q "/gogf/gf/.*/v2" go.mod; then
        sed -i "s/gogf\/gf\(\/.*\)\/v2/gogf\/gf\/v2\1/g" coverage.out
    fi
    cd -
done