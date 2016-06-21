package utils

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //导入驱动
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbcon := "saltfish:saltfish@tcp(114.215.145.118:3306)/hello?charset=utf8&loc=Asia%2fShanghai"
	orm.RegisterDataBase("default", "mysql", dbcon, 30)
}
