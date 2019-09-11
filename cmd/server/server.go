package main

import (
  "flag"
  "log"
  "net/http"
  "strings"
)

func main() {
  port := flag.String("p", "3131", "listen port")
  flag.Parse()
  listen := "localhost:" + *port
  log.Printf("http://%s", listen)

  http.HandleFunc("/", wasmHandler)

  log.Fatal(http.ListenAndServe(listen, nil))
}
func wasmHandler(resp http.ResponseWriter, req *http.Request) {
  enableCors(&resp)
  if strings.HasSuffix(req.URL.Path, ".wasm") {
    resp.Header().Set("content-type", "application/wasm")
  }
  http.FileServer(http.Dir(".")).ServeHTTP(resp, req)
}

func enableCors(w *http.ResponseWriter) {
  (*w).Header().Set("Access-Control-Allow-Origin", "*")
}
