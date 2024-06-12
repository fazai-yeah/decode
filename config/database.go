package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=192.168.2.80 port= 30432 user=postgres password=Zdf@7618 dbname=moriss_smdb "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 设置搜索路径
	//db.Exec("SET search_path TO sod")
	// 这里可以添加自动迁移
	return db
}
