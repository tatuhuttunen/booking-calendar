#!/usr/bin/env bash

p=`pwd`
for d in $(ls ./cmd); do
  echo "verifying cmd/$d"
  cd $p/cmd/$d
  go fmt
  gometalinter --vendor --disable=gotype .
done
cd $p
