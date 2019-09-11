#!/bin/bash

if [[ -z "$1" ]] ; then
  echo "copy where?"
  exit 13
fi

cp -v "$(go env GOROOT)/misc/wasm/wasm_exec.js" "$1"
