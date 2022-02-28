package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "github.com/grokify/beegoutil/demo1/routers"
)

func main() {
	web.Run()
}
