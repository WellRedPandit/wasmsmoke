package main

import (
  "fmt"
  "log"
  "syscall/js"
)

type Digit struct {
  Deva string
  Arab string
}

var Digits = []Digit{
  {"०", "0"},
  {"१", "1"},
  {"२", "2"},
  {"३", "3"},
  {"४", "4"},
  {"५", "5"},
  {"६", "6"},
  {"७", "7"},
  {"८", "8"},
  {"९", "9"},
}

var digidx = 0;

func nextDigit() Digit {
  if digidx > 9 {
    digidx = 0;
  }
  dig := Digits[digidx]
  digidx++
  return dig
}

func main() {
  funRootOnClick := js.FuncOf(RootOnClick)
  draSearch := js.Global().Get("document").Call("getElementById", "root")
  draSearch.Set("onclick", funRootOnClick)

  forever := make(chan bool)
  js.Global().Get("document").
    Call("getElementById", "wasmReady").
    Set("innerHTML", "<b>Loaded</b>")
  js.Global().Get("document").
    Call("getElementById", "loadWasm").
    Set("hidden", true)
  RootOnClick(js.ValueOf(nil),nil)
  <-forever
}

func RootOnClick(this js.Value, args []js.Value) interface{} {
  dig := nextDigit()
  log.Printf("DIGIT: %s %s", dig.Arab, dig.Deva)
  js.Global().Get("document").
    Call("getElementById", "root").
    Set("innerHTML", fmt.Sprintf("<b>%s</b> %s", dig.Arab, dig.Deva))
  return nil
}
