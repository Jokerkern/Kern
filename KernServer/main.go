package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.Default()
	// TLS 使用文件例子
	err := app.Run(iris.TLS("127.0.0.1:8080", "www.jokerkern.cn_bundle.crt", "www.jokerkern.cn.key"))
	if err != nil {
		return
	}
}
