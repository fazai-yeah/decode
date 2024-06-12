package service

import (
	"decode/common"
	"decode/models/DatabaseStruct"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strings"
)

// ProcessFileAndSave 解析文件路径，提取元数据并保存到数据库
func ProcessFileAndSave(filePath string, db *gorm.DB) (string, error) {
	parts := strings.Split(filePath, "/")
	fileName := parts[len(parts)-1]
	fileParts := strings.Split(fileName, "_")

	if len(fileParts) < 6 {
		return "", common.ErrNotEnoughParts
	}

	record := DatabaseStruct.SateFY4BL1FileTab{
		DFileID:   fileParts[0],
		DDataID:   fileParts[1],
		VFileName: fileName,
	}
	log.Info().Msg("这是名称" + fileName)
	log.Info().Msg(common.StructToString(record))
	if result := db.Table("sate_fy4b_l1_file_tab").Create(&record); result.Error != nil {
		return "", result.Error
	}

	return record.DFileID, nil
}
