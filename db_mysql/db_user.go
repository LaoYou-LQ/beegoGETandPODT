package db_mysql

import (
	"beegoResp/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	//连接数据库
	config := beego.AppConfig
	//配置
	//conf\app.conf
	dbDriveername := config.String("db_driveerName")//mysql
	dbUser := config.String("db_user")//root
	dbPassword := config.String("db_password")//123456
	dbIp := config.String("db_ip")//（127.0.0.1:3306）
	dbName := config.String("db_name")//数据库名称

	connUrl := dbUser + ":" + dbPassword + "@tcp" + dbIp+"/" + dbName + "?charset=utf8"

	//打开数据库。sql.Open
	db, err1 := sql.Open(dbDriveername, connUrl)
	//数据库如果打开错误的话，就没必要继续往下执行了，（执行了也是错的）所以Panic
	if err1 != nil {
		//fmt.Println(err1.Error())
		panic("数据库连接错误，请检查配置")
	}
	//为全局变量赋值，连接没有错误就把值给Db这个全局变量
	Db = db
	fmt.Println("连接成功咯！！")
}

//将信息保存到数据库中
func InserUser(user models.User) (int64, error) {
	//用户密码需要保密，使用md5加密，
	hashMd5 := md5.New()
	hashMd5.Write([]byte(user.Password)) //获得结构体user中的用户密码并粉碎
	bytes := hashMd5.Sum(nil)  //粉碎后返回
	//将用户密码变成十六进制并返回
	user.Password = hex.EncodeToString(bytes) //hex.EncodeToString:返回十六进制编码
	fmt.Println("用户名：", user.Usere)
	//Exec:执行
	//（insert into 表名 （数据库表中的值） values（有多少值就写多少问号） 实参.值）
	res, err := Db.Exec("insert into users (users, birthday , address ,nick ,password) values(?,?,?,?,?)", user.Usere, user.Birthday, user.Address, user.Nick, user.Password)
	//保存数据时遇到错误
	if err != nil {
		return -1, err
	}
	//返回受影响的行数
	num, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	return num, nil
}
