package models

//构建用户结构体，用于表示数据结构的定义。
//因返回的值是json格式，所以`json: " "`

type User struct {
	Usere string `json:"name"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
	Nick string `json:"nick"`
	Password string `json:"password"`
}
