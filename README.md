# WebAssembly Smoke Test

## Steps

if upgrading go, run

`./bootstrap_wasm_exec.sh`

then to build

```
./build_server.sh
./build_smoke.sh

```

then to run

`./server`

then navigate to http://localhost:3131/static/smoke/ in your browser

## Notes

Mac:
```
/usr/local/opt/go/libexec/misc/wasm/go_js_wasm_exec
/usr/local/opt/go/libexec/misc/wasm/wasm_exec.html
/usr/local/opt/go/libexec/misc/wasm/wasm_exec.js
```
