package main

import (
	"fmt"
	"time"

	v8 "rogchap.com/v8go"
)

func main() {
	src := `function factorial(n) {
    return n === 1 ? n : n * factorial(--n);
}

let i = 0;
let result = 0;

while (i++ < 1e8) {
    result = factorial(10);
}
result;
`
	script := `main.js`

	ctx := v8.NewContext()
	startedAt := time.Now()
	runScript, err := ctx.RunScript(src, script)
	if err == nil {
		fmt.Println(runScript)
	} else {
		fmt.Println(err.Error())
	} // executes a script on the global context
	fmt.Printf("elapsed: %v\n", time.Since(startedAt))
}
