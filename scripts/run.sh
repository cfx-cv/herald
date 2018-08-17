#!/usr/bin/env bash
set -e

go install -v ./...
$(basename $(pwd))
