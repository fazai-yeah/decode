package DatabaseStruct

import (
	"time"
)

// SateFY4BL1FileTab corresponds to the structure of the PostgreSQL table
type SateFY4BL1FileTab struct {
	DFileID            string     `gorm:"primary_key;column:d_file_id"`
	DDataID            string     `gorm:"column:d_data_id"`
	DIymdhm            time.Time  `gorm:"column:d_iymdhm"`
	DRymdhm            time.Time  `gorm:"column:d_rymdhm"`
	DDateTime          time.Time  `gorm:"column:d_datetime"`
	DStorageSite       string     `gorm:"column:d_storage_site"`
	DFileSize          int64      `gorm:"column:d_file_size"`
	VFiletime          string     `gorm:"column:v_filetime"`
	VLevel             string     `gorm:"column:v_level"`
	VCCCC              string     `gorm:"column:v_cccc"`
	V0100701           string     `gorm:"column:v01007_01"`
	V0201901           string     `gorm:"column:v02019_01"`
	V05310             string     `gorm:"column:v05310"`
	V01310             string     `gorm:"column:v01310"`
	V02402             string     `gorm:"column:v02402"`
	V2900101           string     `gorm:"column:v29001_01"`
	V04001             int        `gorm:"column:v04001"`
	V04002             int        `gorm:"column:v04002"`
	V04003             int        `gorm:"column:v04003"`
	V04004             int        `gorm:"column:v04004"`
	V04005             int        `gorm:"column:v04005"`
	V04006             *int       `gorm:"column:v04006"`
	V05301             string     `gorm:"column:v05301"`
	V01015Rev          string     `gorm:"column:v01015_rev"`
	VFileFormat        string     `gorm:"column:v_file_format"`
	VFileName          string     `gorm:"column:v_file_name"`
	VFileNameSource    string     `gorm:"column:v_file_name_source"`
	V404101            string     `gorm:"column:v40410_1"`
	V404102            string     `gorm:"column:v40410_2"`
	V404103            string     `gorm:"column:v40410_3"`
	VRetain1C          string     `gorm:"column:v_retain1_c"`
	VRetain2C          string     `gorm:"column:v_retain2_c"`
	VRetain3C          string     `gorm:"column:v_retain3_c"`
	DSourceID          string     `gorm:"column:d_source_id"`
	DUpdateTime        *time.Time `gorm:"column:d_update_time"`
	DFileSaveHierarchy int        `gorm:"column:d_file_save_hierarchy;default:0"`
}

// TableName overrides the table name used by GORM for the Alert model.
func (SateFY4BL1FileTab) TableName() string {
	return "sod.sate_fy4b_l1_file_tab"
}
