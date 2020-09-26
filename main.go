package main

import (
	_ "beegoResp/routers"
	"github.com/astaxie/beego"
)

func main() {
	/*
	config :=beego.AppConfig
	appName :=config.String("appname")
	fmt.Println("应用名称",appName)
	port ,err :=config.Int("httpport")
	if err !=nil {
		panic("配置文件解析失败，请检查配置文件")
	}
	fmt.Println(port)
	dbDriveername :=config.String("db_driveerName")
	dbUser :=config.String("db_user")
	dbPassword :=config.String("db_password")
	dbIp :=config.String("db_ip")
	dbName :=config.String("db_name")

	connUrl :=dbUser+":"+dbPassword +"@tcp" + dbIp + dbName +"?chareset=utf8"
	db , err :=sql.Open(dbDriveername,connUrl)
	if err !=nil {
		panic("数据库连接错误，请检查配置")
	}

	 */

	beego.Run()
}

