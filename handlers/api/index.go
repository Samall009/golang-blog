package api

import (
	"fmt"
	"github.com/kataras/iris"
	"lucky1998/blog/handlers"
)

func Index(ctx iris.Context) {
	// 初始化结构体
	type format struct {
		Message string `json:"message"`
		Name    string `json:"name"`
	}

	// 获取结构体数据
	data := handlers.ReadJson(&format{}, ctx)

	if data == (format{}) {
		fmt.Println(111)
	}
}
