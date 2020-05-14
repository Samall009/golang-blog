package routers

import (
	"github.com/kataras/iris"
)

// 路由基础结构
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc iris.Handler
}

// 定义一个路由集合类型
type routeCollect []Route

// 定义一个路由容器
var collect = routeCollect{}

// 添加一个路由
func (collect *routeCollect) AddRouter(router Route) {
	*collect = append(*collect, router)
}
