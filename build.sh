#!/bin/bash

set -Eeuo pipefail

go build test/simple.go
go build test/simple2.go
go build -tags debug test/simple.go
go build -tags debug test/simple2.go
rm -f simple simple2
