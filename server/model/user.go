package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"server/storage"
)

// mysql
var sqlDb *sql.DB

func Test(msg string) (string, interface{}) {

	db, _ := storage.Connect("cha")
	data, _ := storage.GetAll(db, "SELECT * FROM `user` WHERE `id` =1")
	//data := storage.Insert(db, "INSERT INTO user (`username`,`password`,`nickname`,`email`,`remark`,`status`,`createtime`) VALUES( 'vv', 'vv', 'vv' ,'vv' ,'vv' ,'5' ,'ff' )")
	//data, _ := storage.GetRow(db, "SELECT * FROM `user` LIMIT 1")
	//data, _ := storage.Exec(db, "UPDATE user SET `username` = 'qinyinbo' WHERE id = 1 LIMIT 1")
	//fmt.Println(data)
	return msg, data
}
