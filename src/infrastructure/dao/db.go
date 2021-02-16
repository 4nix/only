// Package dao 数据持久层 mysql
package dao

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/BurntSushi/toml"
)

// MySQLconfig 数据库配置
type MySQLconfig struct {
	User     string
	Pass     string
	Database string
	Host     string
}

// Config Toml 配置
type Config struct {
	Mysql MySQLconfig
}

// DB 数据连接对象
var DB *sqlx.DB
var err error

func init() {
	var env = flag.String("env", "", "环境变量")
	flag.Parse()

	conf := new(Config)
	if *env == "production" {
		toml.DecodeFile("src/application/config/db.toml", &conf)
	} else {
		toml.DecodeFile("src/application/config/db_dev.toml", &conf)
	}

	DB, err = sqlx.Open("mysql", conf.Mysql.User+":"+conf.Mysql.Pass+"@tcp("+conf.Mysql.Host+":3306)/"+conf.Mysql.Database)

	fmt.Println(conf)
	fmt.Println(err)
	fmt.Println(DB)
	fmt.Println(conf.Mysql.Host)

	fmt.Println("-env:", *env)

	// p := Groups{}
	// DB.Get(&p, "SELECT * FROM nose_group")
	// fmt.Println(err)
	// fmt.Println(conf.Dev)
	// fmt.Println(DB)
	// fmt.Println(p)
}
