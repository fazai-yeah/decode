package service

import (
	"decode/models/DatabaseStruct"
	"encoding/xml"
	"fmt"
	"gorm.io/gorm"
	"os"
	"time"
)

func HandleAlertFile(filePath string, db *gorm.DB) (string, error) {
	// 打开XML文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	// 解析XML文件
	var alert DatabaseStruct.AlertTab
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&alert)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return "", err
	}
	times := time.Now()
	alert.CreateTime = &times
	parsetTime := times
	parsetTime, err = time.Parse("2006-01-02 15:04:05Z07:00", alert.SendTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return "", err
	}
	fmt.Println("Parsed time:", parsetTime) // 调试输出解析后的完整时间
	alert.SendTime = parsetTime.Format("2006-01-02 15:04:05")
	// 保存到数据库
	err = db.Create(&alert).Error
	if err != nil {
		fmt.Println("Error saving to database:", err)
		return "", err
	}
	fmt.Println("Data saved to database")
	return "success", nil

}
