package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	 */
	c.GetString("user")//返回字符串
	c.GetInt("age")//返回整数

	userName :=c.Ctx.Input.Query("user")//user : Postman中的key值
	userPassword :=c.Ctx.Input.Query("pwd")//pwd : Postman中的key值
	//使用固定值进行数据校验
	//qiangzi  123456
	if userName != "qiangzi" || userPassword != "123456"{
		c.Ctx.WriteString("对不起，数据校验错误。")
		return
	}
	c.Ctx.WriteString("恭喜，数据校验成功。")
}
