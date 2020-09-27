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
/*     //post第一种方法
func (c *MainController) Post()  {
	//post请求form
	//var re models.Respon
	name := c.Ctx.Request.FormValue("name")
	age := c.Ctx.Request.FormValue("age")
	if name !="qiangzi" && age !="12346" {
		c.Ctx.WriteString("对不起，账号或密码错误")
		return
	}
	c.Ctx.WriteString(name)
	c.Ctx.WriteString(age)
	c.Ctx.WriteString("登入成功")
}

 */
/*
        /第二种方法
func (c *MainController) Post()  {
/respon：结构体
	var respon models.Respon   //引用结构体
	data ,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil {
		c.Ctx.WriteString("数据获取失败")
		return
	}
	err =json.Unmarshal(data, &respon )//解析数据
	if err !=nil {
		c.Ctx.WriteString("数据解析错误")//在postman中 “打印”
		return
	}
	c.Ctx.WriteString("数据解析成功啦！")

}
 */
