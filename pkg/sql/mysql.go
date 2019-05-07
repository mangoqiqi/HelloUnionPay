package sql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func NewEngine(database string, user string, pass string, ip string, port string) (engine *xorm.Engine, err error) {
	return xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		user, pass, ip, port, database))
}

