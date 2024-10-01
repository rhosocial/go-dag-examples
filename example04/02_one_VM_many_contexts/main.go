package main

import (
	"errors"

	v8 "rogchap.com/v8go"
)

func main() {
	iso := v8.NewIsolate()     // creates a new JavaScript VM
	ctx1 := v8.NewContext(iso) // new context within the VM
	ctx1.RunScript("const multiply = (a, b) => a * b", "math.js")

	ctx2 := v8.NewContext(iso) // another context on the same VM
	if _, err := ctx2.RunScript("multiply(3, 4)", "main.js"); err != nil {
		// this will error as multiply is not defined in this context
		var jsError *v8.JSError
		if errors.As(err, &jsError) {
			print(jsError.StackTrace)
		} else {
			print(err.Error())
		}
	}
}
