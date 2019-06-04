package main

import (
	"cs/model"
	//"database/sql"
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//打开数据库
	////DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	//db, err := sql.Open("mysql", "root:wang0203@tcp(192.168.86.128:3306)/test?charset=utf8")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	////插入数据
	//stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	//
	//res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	//
	//id, err := res.LastInsertId()
	//fmt.Println(id)
	db := conn()
	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// 自动迁移模式
	db.AutoMigrate(&model.UserInfo{})
}
func conn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:wang0203@tcp(192.168.86.128:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}
	return db
}
func init() {
	db := conn()
	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// 自动迁移模式
	db.AutoMigrate(&model.UserInfo{})
}
