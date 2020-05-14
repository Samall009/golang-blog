package routers

import (
	"lucky1998/blog/handlers/api"
)

func init() {
	collect.AddRouter(Route{
		"ApiIndex",
		"ALL",
		"api/",
		api.Index,
	})
}
