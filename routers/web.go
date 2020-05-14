package routers

import "lucky1998/blog/handlers/web"

var webRouter = routeCollect{
	{
		"home",
		"GET",
		"/home",
		web.Home,
	},
}

func init() {
	// 单注册方式
	collect.AddRouter(Route{
		"index",
		"GET",
		"/",
		web.Index,
	})

	// 多种路由添加方式
	// 循环集合方式
	for _, router := range webRouter {
		collect.AddRouter(router)
	}
}
