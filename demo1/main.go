package main

import (
	"github.com/astaxie/beego"
	_ "github.com/grokify/beego-oauth2-demo/demo1/routers"
)

func main() {
	beego.Run()
}
