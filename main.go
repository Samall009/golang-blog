package main

import (
	"github.com/kataras/iris"
	"lucky1998/blog/routers"
)

// 定义一个服务指针
var app *iris.Application

// 初始化执行方法?
func init() {
	app = iris.New()
	// 注册路由
	routers.RouterRegister(app)
}

func main() {
	app.Run(iris.Addr(":8080"))
}
