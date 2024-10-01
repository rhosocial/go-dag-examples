package main

import (
	v8 "rogchap.com/v8go"
)

func main() {
	source := "const multiply = (a, b) => a * b"
	iso1 := v8.NewIsolate()                                                         // creates a new JavaScript VM
	ctx1 := v8.NewContext(iso1)                                                     // new context within the VM
	script1, _ := iso1.CompileUnboundScript(source, "math.js", v8.CompileOptions{}) // compile script to get cached data
	_, _ = script1.Run(ctx1)

	cachedData := script1.CreateCodeCache()

	iso2 := v8.NewIsolate()     // create a new JavaScript VM
	ctx2 := v8.NewContext(iso2) // new context within the VM

	script2, _ := iso2.CompileUnboundScript(source, "math.js", v8.CompileOptions{CachedData: cachedData}) // compile script in new isolate with cached data
	_, _ = script2.Run(ctx2)
}
