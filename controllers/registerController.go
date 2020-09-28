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
//最开始当然要配置好MySQL数据的啦（app.conf  db_user.go  config）
//1.用ioutil.ReadAll接收数据
//2.接收完后用json.Unmarshal解析数据
//3.解析完后用Db.Exec保存到数据库中，并返回收到影响的行数。（num, err := res.RowsAffected()）
//Db.Exec（执行）

//该方法用于处理post类型请求
func (r *RegisterConeroller) Post()  {

	//接收数据
	body , err :=ioutil.ReadAll(r.Ctx.Request.Body)//一直读取，直到遇到错误。
	if err!=nil {
		//读取数据遇到错误，在postman中 “打印”《数据接收错误》
		r.Ctx.WriteString("数据接收错误")
		return
	}

	//解析数据
	//user: models文件中的user.go中的结构体
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
