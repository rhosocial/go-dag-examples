package main

import (
	"fmt"

	v8 "rogchap.com/v8go"
)

func main() {
	ctx := v8.NewContext()                           // new context with a default VM
	obj := ctx.Global()                              // get the global object from the context
	obj.Set("version", "v1.0.0")                     // set the property "version" on the object
	val, _ := ctx.RunScript("version", "version.js") // global object will have the property set within the JS VM
	fmt.Printf("version: %s\n", val)

	if obj.Has("version") { // check if a property exists on the object
		fmt.Printf("version deleted: %v\n", obj.Delete("version")) // remove the property from the object
	}
}
