// Package dao 数据持久层 mysql
package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/BurntSushi/toml"
)

// Dbconfig 数据库配置
type Dbconfig struct {
	User     string
	Pass     string
	Database string
	Host     string
}

// Config Toml 配置
type Config struct {
	Dev Dbconfig
}

// DB 数据连接对象
var DB *sqlx.DB
var err error

func init() {
	conf := new(Config)
	toml.DecodeFile("src/application/config/db.toml", &conf)

	DB, err = sqlx.Open("mysql", conf.Dev.User+":"+conf.Dev.Pass+"@tcp("+conf.Dev.Host+":3306)/"+conf.Dev.Database)

	fmt.Println(err)
	fmt.Println(DB)
	fmt.Println(conf)

	// p := Groups{}
	// DB.Get(&p, "SELECT * FROM nose_group")
	// fmt.Println(err)
	// fmt.Println(conf.Dev)
	// fmt.Println(DB)
	// fmt.Println(p)
}
