package web

import "github.com/kataras/iris"

// 首页目录
func Index(ctx iris.Context) {
	ctx.HTML("<h1>hello word</h1>")
}

// 家目录
func Home(ctx iris.Context) {
	ctx.HTML("<h1>家目录啊</h1>")
}
