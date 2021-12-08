package main

import (
	"github.com/ewenliu/gin-example/8.RouterSplit/cat"
	"github.com/ewenliu/gin-example/8.RouterSplit/dog"
	"github.com/ewenliu/gin-example/8.RouterSplit/routers"
)

func main() {
	routers.IncludeHandlers(cat.LoadHandler, dog.LoadHandler)
	r := routers.InitRouters()
	r.Run(":8888")
}
