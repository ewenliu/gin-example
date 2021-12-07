package main

import (
	"fmt"
	settings "github.com/ewenliu/gin-example/6.settings/env"
)

func main() {
	fmt.Println(settings.Env["GinMode"])
	fmt.Println(settings.Env["Port"])
}
