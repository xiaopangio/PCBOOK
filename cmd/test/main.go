// Package test  @Author xiaobaiio 2023/3/4 13:00:00
package main

import (
	"fmt"
	"github.com/xiaopangio/pcbook/orm"
	"github.com/xiaopangio/pcbook/orm/dal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(orm.Dsn()))
	if err != nil {
		fmt.Errorf("cannot open db: %w", err)
	}
	dal.SetDefault(db)
	user, err := dal.User.Where(dal.User.UserID.Eq(1)).Preload(dal.User.Role).First()
	if err != nil {
		fmt.Errorf("no record")
	}
	fmt.Println(user)
	role, err := dal.Role.Where(dal.Role.RoleID.Eq(1)).First()
	if err != nil {
		fmt.Errorf("no record")
	}
	fmt.Println(role)
}
