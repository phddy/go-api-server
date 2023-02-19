package common

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB

func init() {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.DB.Id, config.DB.Pw, config.DB.Host, config.DB.Database)
	print(conn)
	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", config.Database.Id, config.Database.Pw, config.Database.Host))
	//if err != nil {
	//	panic(err)
	//}
	//
	//db.SetConnMaxLifetime(time.Minute * 3)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", conn)
	orm.RunSyncdb("default", false, true)
	//defer db.Close()
}

//func GetConnection() *sql.DB {
//	return db
//}
