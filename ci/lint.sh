#!/bin/bash

set -e

confidence_level="0.9"

echo "# Running golint"
golint -set_exit_status -min_confidence $confidence_level ./...

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
