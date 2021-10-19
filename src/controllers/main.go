package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"icpscan/src/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var RedisDb *redis.Client

func init() {
	var cfg config.DbConfig
	if err := load(&cfg, "icpscan.config"); err != nil {
		log.Fatalf("err is %v", err)
	}
	fmt.Printf("cf is %+v", cfg)
	connection, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/icp_scan?charset=utf8mb4&parseTime=True&loc=Local", cfg.UserName, cfg.Password, cfg.Addr, cfg.Port)),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db = connection

	RedisDb = redis.NewClient(&redis.Options{
		Addr: "gva-redis:6379",
		// Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func load(cfg interface{}, f string) error {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, cfg)
}
