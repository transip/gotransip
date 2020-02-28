#!/bin/sh

confidence_level="0.9"

golint -set_exit_status -min_confidence $confidence_level ./...