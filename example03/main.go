package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/robertkrimen/otto"
)

// 应用提供的登录接口
func login(username string, password string) bool {
	// 模拟登录逻辑
	if username == "admin" && password == "123456" {
		return true
	}
	return false
}

// 应用提供的退出接口
func logout() {
	// 模拟退出逻辑
	fmt.Println("退出成功")
}

// 应用提供的获取用户信息接口
func getUserInfo() map[string]string {
	// 模拟获取用户信息逻辑
	userInfo := map[string]string{
		"username": "admin",
		"email":    "admin@example.com",
	}
	return userInfo
}

func main() {
	// 解析命令行参数
	jsFile := flag.String("js", "", "JavaScript文件路径")
	flag.Parse()

	// 检查是否提供了JavaScript文件路径
	if *jsFile == "" {
		fmt.Println("请提供JavaScript文件路径")
		return
	}

	// 读取JavaScript文件的内容
	userJsContent, err := os.ReadFile(*jsFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 执行JavaScript代码
	vm := otto.New()
	// 注册应用提供的接口
	vm.Set("app", map[string]interface{}{
		"login":       login,
		"logout":      logout,
		"getUserInfo": getUserInfo,
	})

	// 执行JavaScript代码
	_, err = vm.Run(string(userJsContent))
	if err != nil {
		fmt.Println(err)
		return
	}
}
