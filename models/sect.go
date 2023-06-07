package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Section struct {
	FCCODE     string      `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP     string      `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCDEPT     string      `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string      `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCFCHR     string      `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCNAME     string      `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2    string      `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCSKID     string      `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"`
	FTDATETIME time.Time   `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time   `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time   `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTSRCUPD   time.Time   `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
	CORP       *Corp       `gorm:"foreignKey:FCCORP;references:FCSKID;" json:"corp"`
	DEPT       *Department `gorm:"foreignKey:FCDEPT;references:FCSKID;" json:"dept"`
}

func (Section) TableName() string {
	return "SECT"
}

func (obj *Section) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
