package routers

import "github.com/gin-gonic/gin"

type loadHandler func(engine *gin.Engine)

var loadHandlers []loadHandler


func IncludeHandlers(objs ...loadHandler){

	loadHandlers = append(loadHandlers, objs...)
}

func InitRouters() *gin.Engine{
	r:= gin.Default()
	for _, loadHandler := range loadHandlers{
		loadHandler(r)
	}
	return r
}