package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"icpscan/src/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var cfg config.DbConfig
	if err := load(&cfg, "icpscan.config"); err != nil {
		log.Fatalf("err is %v", err)
	}
	connection, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/icpscan?charset=utf8mb4&parseTime=True&loc=Local", cfg.UserName, cfg.Password, cfg.Addr, cfg.Port)),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db = connection
}


func load(cfg interface{}, f string) error {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, cfg)
}
