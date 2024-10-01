package main

import "github.com/rhosocial/go-dag/workflow/simple"

func main() {
	_, _ = simple.NewWorkflow[struct{}, struct{}]()
}
