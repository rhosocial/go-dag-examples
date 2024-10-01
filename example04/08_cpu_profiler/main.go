package main

import (
	"fmt"
	"time"

	v8 "rogchap.com/v8go"
)

func main() {
	ctx := v8.NewContext() // new context with a default VM

	vals := make(chan *v8.Value, 1)
	errs := make(chan error, 1)
	script := `function stopForAWhile() {
  console.log("停止前");
  setTimeout(function() {
    console.log("停止后");
  }, 5000); // 停止 5 秒
  stopForAWhile();
}`

	go func() {
		val, err := ctx.RunScript(script, "forever.js") // exec a long running script
		if err != nil {
			errs <- err
			return
		}
		vals <- val
	}()

	select {
	case val := <-vals:
		fmt.Println(val)
		// success
	case err := <-errs:
		fmt.Println(err.Error())
		// javascript error
	case <-time.After(200 * time.Microsecond):
		vm := ctx.Isolate()     // get the Isolate from the context
		vm.TerminateExecution() // terminate the execution
		err := <-errs           // will get a termination error back from the running script
		fmt.Println(err.Error())
	}
}

func createProfile() {
	iso := v8.NewIsolate()
	ctx := v8.NewContext(iso)
	cpuProfiler := v8.NewCPUProfiler(iso)

	cpuProfiler.StartProfiling("my-profile")

	//ctx.RunScript(profileScript, "script.js") # this script is defined in cpuprofiler_test.go
	//val, _ := ctx.Global().Get("start")
	//fn, _ := val.AsFunction()
	//fn.Call(ctx.Global())
	//
	//cpuProfile := cpuProfiler.StopProfiling("my-profile")
	//
	//printTree("", cpuProfile.GetTopDownRoot()) # helper function to print the profile
}

func printTree(nest string, node *v8.CPUProfileNode) {
	fmt.Printf("%s%s %s:%d:%d\n", nest, node.GetFunctionName(), node.GetScriptResourceName(), node.GetLineNumber(), node.GetColumnNumber())
	count := node.GetChildrenCount()
	if count == 0 {
		return
	}
	nest = fmt.Sprintf("%s  ", nest)
	for i := 0; i < count; i++ {
		printTree(nest, node.GetChild(i))
	}
}
