// Package orm  @Author xiaobaiio 2023/3/4 15:11:00
package orm

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var v *viper.Viper
var DB *gorm.DB

func init() {
	v = viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("app")
	v.AddConfigPath("./config/")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %w", err))
	}
	db, err := gorm.Open(mysql.Open(Dsn()))
	if err != nil {
		log.Fatal("cannot open db: ", err)
	}
	DB = db
}
func Dsn() string {
	root := v.GetString("database.root")
	password := v.GetString("database.password")
	ip := v.GetString("database.ipaddress")
	port := v.GetString("database.port")
	dbName := v.GetString("database.db_name")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", root, password, ip, port, dbName)
}
