package controllers

import (
	"beego_blog1/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type LinkController struct {
	beego.Controller

}

func (c *LinkController) Link() {
	var link models.Link
	if c.Ctx.Request.Method == "POST"{
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &link); err != nil{
			fmt.Println(err)
			c.Ctx.WriteString("parse error \r\n")

		}else {
			link.Insert()
			fmt.Println("link =", link)
			c.Ctx.WriteString("parse success \r\n")
		}
	}
}