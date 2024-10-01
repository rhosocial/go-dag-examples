package main

import (
	"fmt"

	v8 "rogchap.com/v8go"
)

func main() {
	iso := v8.NewIsolate() // create a new VM
	// a template that represents a JS function
	printfn := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		fmt.Printf("%v", info.Args()) // when the JS function is called this Go callback will execute
		return nil                    // you can return a value back to the JS caller if required
	})
	global := v8.NewObjectTemplate(iso)       // a template that represents a JS Object
	global.Set("print", printfn)              // sets the "print" property of the Object to our function
	ctx := v8.NewContext(iso, global)         // new Context with the global Object set to our object template
	ctx.RunScript("print('foo')", "print.js") // will execute the Go callback with a single argunent 'foo'
}
