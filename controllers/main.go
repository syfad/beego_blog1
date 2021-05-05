package controllers

import (
	"beego_blog1/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

//首页展示
func (c *MainController) Index() {
	var list []*models.Post
	post := models.Post{}
	//文章表句柄
	query := orm.NewOrm().QueryTable(&post).Filter("status", 0)
	_, err := query.All(&list)
	if err != nil{
		fmt.Println(err)
	}
	//fmt.Println(list)
	// list传到页面
	c.Data["list"] = list
	// Layout是指定不动的部分
	theme := "double"
	c.Layout = theme + "/layout.html"
	c.TplName = theme + "/index.html"
	// LayoutSections
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["head"] = theme + "/head.html"
	c.LayoutSections["banner"] = theme + "/banner.html"
	c.LayoutSections["middle"] = theme + "/middle.html"
	c.LayoutSections["right"] = theme + "/right.html"
	c.LayoutSections["foot"] = theme + "/foot.html"
}
