package main

import (
	"github.com/go-sql-driver/mysql"
	"gopkg.in/redis.v4"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"database/sql"
	"content"
)

type Option struct {
	Port int `yaml:"port"`
	Mysql mysql.Config `yaml:"mysql"`
	Redis redis.Options `yaml:"redis"`
	UGC struct{
		Dir string `yaml:"dir"`
		Prefix string `yaml:"prefix"`
	} `yaml:"ugc_files"`
}

var inst *Option
var db *sql.DB

/**
初始化运行环境
 */
func EnvironmentInitialize(configFile string) (*Option, error) {
	inst = &Option{
		Port: 80,
	}
	c, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(c, inst)
	if err != nil {
		return nil, err
	}
	sqlDSN := inst.Mysql.FormatDSN()
	db, err = sql.Open("mysql", sqlDSN)
	if err != nil {
		return nil, err
	}
	content.InitManager(db)
	return inst, nil
}

/**
销毁环境
 */
func EnvironmentFinish() {
	db.Close()
}

func EnvironmentGet() *Option {
	return inst
}
