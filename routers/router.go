package routers

import (
	"github.com/kataras/iris"
)

// RouterRegister web路由注册
func RouterRegister(app *iris.Application) {
	// 注册路由
	for _, router := range collect {
		app.Handle(router.Method, router.Pattern, router.HandlerFunc)
	}
}
