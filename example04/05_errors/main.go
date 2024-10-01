package main

import (
	"errors"
	"fmt"

	v8 "rogchap.com/v8go"
)

func main() {
	ctx := v8.NewContext() // new context with a default VM
	src := `multiply(3, 4)`
	filename := "main.js"
	_, err := ctx.RunScript(src, filename)
	if err != nil {
		var e *v8.JSError
		errors.As(err, &e)        // JavaScript errors will be returned as the JSError struct
		fmt.Println(e.Message)    // the message of the exception thrown
		fmt.Println(e.Location)   // the filename, line number and the column where the error occured
		fmt.Println(e.StackTrace) // the full stack trace of the error, if available

		fmt.Printf("javascript error: %v\n", e)        // will format the standard error message
		fmt.Printf("javascript stack trace: %+v\n", e) // will format the full error stack trace
	}
}
