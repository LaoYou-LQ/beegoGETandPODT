package controllers

import (
	"beegoResp/db_mysql"
	"beegoResp/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

//该结构体用于处理/register请求

type RegisterConeroller struct {
	beego.Controller
}

//该方法用于处理post类型请求
func (r *RegisterConeroller) Post()  {

	//接收数据
	body , err :=ioutil.ReadAll(r.Ctx.Request.Body)
	if err!=nil {
		r.Ctx.WriteString("数据接收错误")
		return
	}

	//解析数据

	var user models.User
	err = json.Unmarshal(body,&user)
	if err !=nil {
		r.Ctx.WriteString("数据解析错误")
		return
	}

	//保存已解析的数据
	num , err :=db_mysql.InserUser(user)
	fmt.Println(err.Error())
	if err !=nil {
		r.Ctx.WriteString("数据存储失败")
		return
	}
	fmt.Println(num)

	res :=models.Respon{
		Code:0,
		Message:"保存好了",
		Data:nil ,
	}
	r.Data["json"] = &res
	r.ServeJSON()
}
