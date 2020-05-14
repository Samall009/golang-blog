package handlers

import (
	"github.com/kataras/iris"
	"reflect"
)

// 读取json数据
func ReadJson(construction interface{}, ctx iris.Context) interface{} {
	// 读取json数据
	err := ctx.ReadJSON(construction)

	// 判断错误
	if err != nil {
		ctx.JSON(resultError(200, err.Error()))
	}

	// 返回数据
	return construction
}

// 基础类型
type result struct {
	code    int
	message string
	error   bool
	data    interface{}
}

// 返回错误信息
func resultError(code int, str string) result {
	return result{
		code,
		str,
		true,
		nil,
	}
}

// 返回字符串信息
func resultMessage(str string) result {
	return result{
		0,
		str,
		false,
		nil,
	}
}

// 返回data信息
func resultData(data interface{}) result {
	return result{
		0,
		"",
		false,
		data,
	}
}

// 返回数据
func Result(reload interface{}) result {
	switch reload.(type) {
	case string:
		return resultMessage(reflect.TypeOf(reload).String())
	case map[interface{}]interface{}:
		return resultData(reload)
	default:
		return result{
			100,
			"无效的返回值类型",
			true,
			nil,
		}
	}
}
