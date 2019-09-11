#!/bin/bash

GOOS=js GOARCH=wasm go build -o static/smoke/smoke.wasm cmd/smoke/smoke.go
