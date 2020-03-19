#!/bin/sh

set -e

# create the diff directory
mkdir -pv /tmp/diff

# create a backup from the go.* files to diff, before and after
cp -v go.* /tmp/diff/

go mod tidy

for f in go.*; do
  diff -u $f /tmp/diff/$f
done
