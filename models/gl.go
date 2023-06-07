package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Gl struct {
	FCACCHART string `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
	FCACLINK  string `gorm:"column:FCACLINK;" json:"fcaclink"  form:"fcaclink" ` //$.
	// FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCATSTEP string `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" ` // I
	// FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	// FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	// FCCLOSED   string    `gorm:"column:FCCLOSED;" json:"fcclosed"  form:"fcclosed" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" ` //$.
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	// FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	// FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	// FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" ` // $$$+
	FCDEPT    string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" ` // E
	FCGLHEAD  string `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	// FCGRP      string    `gorm:"column:FCGRP;" json:"fcgrp"  form:"fcgrp" `
	// FCISBUDGET string    `gorm:"column:FCISBUDGET;" json:"fcisbudget"  form:"fcisbudget" `
	// FCISINBROW string    `gorm:"column:FCISINBROW;" json:"fcisinbrow"  form:"fcisinbrow" `
	// FCISSUMOFB string    `gorm:"column:FCISSUMOFB;" json:"fcissumofb"  form:"fcissumofb" `
	FCJOB     string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" ` // $.
	// FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	// FCPAYMENT  string    `gorm:"column:FCPAYMENT;" json:"fcpayment"  form:"fcpayment" `
	FCPROJ    string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCRECALBY string `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
	FCSECT    string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	// FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ  string `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	// FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	// FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	// FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	// FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	// FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	// FMDETAIL   string    `gorm:"column:FMDETAIL;" json:"fmdetail"  form:"fmdetail" `
	// FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FNAMT      float64   `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNCALLINE  float64   `gorm:"column:FNCALLINE;" json:"fncalline"  form:"fncalline" ` // 0
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	// FTLASRECAL string    `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	// FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	Corp      *Corp       `gorm:"foreignKey:FCCORP;references:FCSKID;" json:"corp"`
	Branch    *Branch     `gorm:"foreignKey:FCBRANCH;references:FCSKID;" json:"branch"`
	Dept      *Department `gorm:"foreignKey:FCDEPT;references:FCSKID;" json:"department"`
	Sect      *Section    `gorm:"foreignKey:FCSECT;references:FCSKID;" json:"section"`
	Glhead    *Glhead     `gorm:"foreignKey:FCGLHEAD;references:FCSKID;" json:"glhead"`
	Acchart   *Acchart    `gorm:"foreignKey:FCACCHART;references:FCSKID;" json:"acchart"`
	Job       *Job        `gorm:"foreignKey:FCJOB;references:FCSKID;" json:"job"`
	CreatedBy *Employee   `gorm:"foreignKey:FCCREATEBY;references:FCSKID;" json:"created_by"`
	UpdatedBy *Employee   `gorm:"foreignKey:FCCORRECTB;references:FCSKID;" json:"updated_by"`
}

func (Gl) TableName() string {
	return "GL"
}

func (obj *Gl) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	obj.FIMILLISEC = time.Now().Unix()
	obj.FTDATETIME = time.Now()
	obj.FTLASTEDIT = time.Now()
	obj.FTLASTUPD = time.Now()
	obj.FCACLINK = "$."
	obj.FCATSTEP = "I"
	obj.FCCREATEAP = "$."
	obj.FCDATASER = "$$$+"
	obj.FCEAFTERR = "E"
	obj.FCLUPDAPP = "$."
	obj.FNCALLINE = 0
	return
}

func (obj *Gl) BeforeUpdate(tx *gorm.DB) (err error) {
	obj.FTLASTEDIT = time.Now()
	obj.FTLASTUPD = time.Now()
	return
}
