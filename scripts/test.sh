#!/bin/bash

# Run Go tests, filter out the noise, but preserve the original exit status
set -o pipefail

go test ./... | grep -v "no test files"