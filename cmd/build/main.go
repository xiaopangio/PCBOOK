// Package build  @Author xiaobaiio 2023/3/4 8:52:00
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var v *viper.Viper

func init() {
	v = viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("app")
	v.AddConfigPath("./config/")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %w", err))
	}
}
func dsn() string {
	root := v.GetString("database.root")
	password := v.GetString("database.password")
	ip := v.GetString("database.ipaddress")
	port := v.GetString("database.port")
	dbName := v.GetString("database.db_name")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parsetime=true&loc=local", root, password, ip, port, dbName)
}
func main() {
	dsn := dsn()
	_, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Errorf("cannot open db: %w", err)
	}

}
