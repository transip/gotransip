#!/bin/bash

function gofmt() {
  gofmt_output=$(go fmt $1)
  if [ ! -z "${gofmt_output}" ]; then
    echo $gofmt_output
  fi
  # test if output of go fmt > empty
  test -z "${gofmt_output}"
}

echo "# Running go fmt"
gofmt ./...
