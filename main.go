package main

import (
	"github.com/gates/gates"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("gates", map[string]interface{}{
		"compile": compile,
	})
}

func compile(x string) *js.Object {
	program, err := gates.Compile(x)
	if err != nil {
		panic(js.Global.Get("SyntaxError").New(err.Error()))
	}
	return js.MakeWrapper(program)
}
