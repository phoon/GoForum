package repository

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/phoon/go-forum/config"
	"github.com/phoon/go-forum/repository/model"
)

var (
	db   *gorm.DB
	err  error
	_once sync.Once
)

func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Fields.Database.Username,
		config.Fields.Database.Password,
		config.Fields.Database.Host,
		config.Fields.Database.DBName,
	)

	db, err = gorm.Open("mysql", dsn)

	if config.Fields.Mode == "debug" {
		db = db.Debug()
	}

	if err != nil {
		log.Fatalln("Fail to open database -> ", err.Error())
	}
	db.AutoMigrate(model.Models...)
	//创建外键
	db.Model(&model.Topic{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&model.Topic{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")

	db.Model(&model.Comment{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&model.Comment{}).AddForeignKey("topic_id", "topics(id)", "CASCADE", "CASCADE")
}

func Start() {
	_once.Do(initDB)
}
