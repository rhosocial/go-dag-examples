package main

import (
	"fmt"
	"log"

	"rogchap.com/v8go"
)

func main() {
	// 创建一个 V8 Isolate 实例，V8 引擎的执行上下文
	iso := v8go.NewIsolate()
	global := v8go.NewObjectTemplate(iso)

	// 向 JavaScript 注入一个 Go 函数
	global.Set("printMessage", v8go.NewFunctionTemplate(iso, printMessage))

	// 创建 V8 执行上下文并加载全局对象
	ctx := v8go.NewContext(iso, global)

	// 执行 JavaScript 文件中的代码
	script := `
        printMessage('JavaScript running in V8');
        // 确保 getObject 已定义
        if (typeof getObject !== 'function') {
            throw new Error('getObject function is not defined');
        }
        getObject();
    `

	// 读取并执行 JavaScript 代码
	_, err := ctx.RunScript(script, "main.js")
	if err != nil {
		log.Fatal(err)
	}

	// 获取 JavaScript 返回的 getObject 函数
	val, err := ctx.Global().Get("getObject")
	if err != nil {
		log.Fatal(err)
	}

	// 确保 getObject 是一个函数
	if val.IsFunction() {
		fn, err := val.AsFunction() // 将 *v8go.Value 转换为 *v8go.Function
		if err != nil {
			log.Fatal(err)
		}

		// 调用 getObject 函数
		result, err := fn.Call(v8go.Undefined(iso)) // 使用 AsFunction().Call() 方法调用函数
		if err != nil {
			log.Fatal(err)
		}

		// result 应该是一个对象
		obj := result.Object()
		if obj != nil {
			// 获取 sayHello 方法
			sayHelloVal, err := obj.Get("sayHello")
			if err != nil {
				log.Fatal(err)
			}

			// 确保 sayHello 是一个函数
			if sayHelloVal.IsFunction() {
				sayHelloFn, err := sayHelloVal.AsFunction() // 将 sayHello 转换为 JavaScript 函数
				if err != nil {
					log.Fatal(err)
				}

				// 使用 v8go.NewValue 创建字符串
				name, err := v8go.NewValue(iso, "Go Developer")
				if err != nil {
					log.Fatal(err)
				}

				// 调用 sayHello 函数
				helloResult, err := sayHelloFn.Call(v8go.Undefined(iso), name)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(helloResult.String()) // 输出：Hello, Go Developer from JavaScript!
			}
		}
	}
}

// Go 代码注入的函数
func printMessage(info *v8go.FunctionCallbackInfo) *v8go.Value {
	fmt.Println(info.Args()[0].String())
	return nil
}
