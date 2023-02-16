package mysql

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once = sync.Once{}
)

func NewDB() *gorm.DB {
	return db
}

func InitDB() {
	dsn := "MiniTikTok:root@tcp(49.232.155.203:3306)/minitiktok?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	InitDBWithDSN(dsn)
}

func InitDBWithDSN(dsn string) {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("连接数据库失败, err: " + err.Error())
		}
	})
}

func UpdateModel() {
	InitDB()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Follow{})
	db.AutoMigrate(&Like{})
	db.AutoMigrate(&Video{})
	db.AutoMigrate(&Comment{})
}
