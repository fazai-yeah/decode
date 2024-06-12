package DatabaseStruct

import (
	"time"
)

// Alert represents the "sod"."alert" table in PostgreSQL and also defines XML mapping for parsing.
type AlertTab struct {
	Identifier string     `gorm:"column:identifier;type:varchar(255)" xml:"identifier"`
	Sender     string     `gorm:"column:sender;type:varchar(255)" xml:"sender"`
	SendTime   string     `gorm:"column:sendtime;type:timestamp" xml:"sendTime"`
	Status     string     `gorm:"column:status;type:varchar(255)" xml:"status"`
	MsgType    string     `gorm:"column:msgtype;type:varchar(255)" xml:"msgType"`
	SCopy      string     `gorm:"column:scopy;type:varchar(255)" xml:"scope"`
	Code       string     `gorm:"column:code;type:varchar(255)" xml:"code"`
	Note       string     `gorm:"column:note;type:varchar(255)" xml:"note"`
	Info       Info       `gorm:"-" xml:"info"`
	CreateTime *time.Time `gorm:"column:creattime;type:timestamp" xml:"creattime"`
	UpdateTime *time.Time `gorm:"column:updatetime;type:timestamp" xml:"updatetime"`
}
type Info struct {
	Methods []Methods `xml:"method"`
}
type Methods struct {
	MethodName  string `xml:"methodName"`
	Message     string `xml:"message"`
	AudienceGrp string `xml:"audienceGrp"`
}

// TableName overrides the table name used by GORM for the Alert model.
func (AlertTab) TableName() string {
	return "sod.alert"
}
