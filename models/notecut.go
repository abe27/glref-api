package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Notecut struct {
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCHILDH   string    `gorm:"column:FCCHILDH;" json:"fcchildh"  form:"fcchildh" `
	FCCHILDI   string    `gorm:"column:FCCHILDI;" json:"fcchildi"  form:"fcchildi" `
	FCCHILDTYP string    `gorm:"column:FCCHILDTYP;" json:"fcchildtyp"  form:"fcchildtyp" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMASTERH  string    `gorm:"column:FCMASTERH;" json:"fcmasterh"  form:"fcmasterh" `
	FCMASTERI  string    `gorm:"column:FCMASTERI;" json:"fcmasteri"  form:"fcmasteri" `
	FCMASTERTY string    `gorm:"column:FCMASTERTY;" json:"fcmasterty"  form:"fcmasterty" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNQTY      float64   `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNUMQTY    float64   `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	Corp       *Corp     `gorm:"foreignKey:FCCORP;references:FCSKID;" json:"corp"`
	Branch     *Branch   `gorm:"foreignKey:FCBRANCH;references:FCSKID;" json:"branch"`
	Orderh     *Orderh   `gorm:"foreignKey:FCCHILDH;references:FCSKID;" json:"orderh"`
	Orderi     *Orderi   `gorm:"foreignKey:FCCHILDI;references:FCSKID;" json:"orderi"`
	Glref      *Glref    `gorm:"foreignKey:FCMASTERH;references:FCSKID;" json:"glref"`
	Refprod    *Refprod  `gorm:"foreignKey:FCMASTERI;references:FCSKID;" json:"refsprod"`
}

func (Notecut) TableName() string {
	return "NOTECUT"
}

func (obj *Notecut) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("G%sF", id)
	obj.FCDATASER = "$$$+"
	obj.FCEAFTERR = "E"
	obj.FCLUPDAPP = "$/"
	obj.FCMASTERTY = "PO"
	obj.FCCHILDTYP = "PR"
	obj.FNUMQTY = 1
	obj.FCCREATEAP = "$/"
	obj.FIMILLISEC = time.Now().Unix()
	obj.FTDATETIME = time.Now() // time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	obj.FTLASTEDIT = time.Now() // time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	obj.FTLASTUPD = time.Now()  //  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	return
}

func (obj *Notecut) BeforeUpdate(tx *gorm.DB) (err error) {
	obj.FTLASTEDIT = time.Now() // time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	obj.FTLASTUPD = time.Now()  //  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	return
}
