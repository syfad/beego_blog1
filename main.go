package main

import (
	_ "beego_blog1/routers"
	_ "beego_blog1/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

