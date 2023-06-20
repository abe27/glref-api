package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Glhead struct {
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `    //E
	FCSTYLE    string    `gorm:"column:FCSTYLE;" json:"fcstyle"  form:"fcstyle" ` //A
	FCLINK     string    `gorm:"column:FCLINK;" json:"fclink"  form:"fclink" `    // $C
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FCACCBOOK  string    `gorm:"column:FCACCBOOK;" json:"fcaccbook"  form:"fcaccbook" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FMREMARK2  string    `gorm:"column:FMREMARK2;" json:"fmremark2"  form:"fmremark2" `
	FMREMARK3  string    `gorm:"column:FMREMARK3;" json:"fmremark3"  form:"fmremark3" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" ` //E
	FCATSTEP   string    `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `    //I
	FNLOCKED   float64   `gorm:"column:FNLOCKED;" json:"fnlocked"  form:"fnlocked" `    //0
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	Corp       *Corp     `gorm:"foreignKey:FCCORP;references:FCSKID;" json:"corp"`
	Branch     *Branch   `gorm:"foreignKey:FCBRANCH;references:FCSKID;" json:"branch"`
	Accbook    *Accbook  `gorm:"foreignKey:FCACCBOOK;references:FCSKID;" json:"accbook"`
	CreatedBy  *Employee `gorm:"foreignKey:FCCREATEBY;references:FCSKID;" json:"created_by"`
	UpdatedBy  *Employee `gorm:"foreignKey:FCCORRECTB;references:FCSKID;" json:"updated_by"`

	// FTLASRECAL string `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
	// FCCANLINK  string `gorm:"column:FCCANLINK;" json:"fccanlink"  form:"fccanlink" `
	// FCCHKDT    string `gorm:"column:FCCHKDT;" json:"fcchkdt"  form:"fcchkdt" `
	// FCCLOSED   string `gorm:"column:FCCLOSED;" json:"fcclosed"  form:"fcclosed" `
	// FCCMD1     string `gorm:"column:FCCMD1;" json:"fccmd1"  form:"fccmd1" `

	// FCAPPNAME  string `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	// FCAPPROVER string `gorm:"column:FCAPPROVER;" json:"fcapprover"  form:"fcapprover" `

	// FCCPATADJ  string `gorm:"column:FCCPATADJ;" json:"fccpatadj"  form:"fccpatadj" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	// FCCREATETY string `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	// FCCUACC    string `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	// FCDATAIMP  string `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `

	// FCDOCFLOWI string `gorm:"column:FCDOCFLOWI;" json:"fcdocflowi"  form:"fcdocflowi" `
	// FCDTYPE1   string `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	// FCDTYPE2   string `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	// FCDTYPE3   string `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	// FCDTYPE4   string `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	// FCDTYPE5   string `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	// FCDTYPE6   string `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	// FCDTYPE7   string `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	// FCDTYPE8   string `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	// FCDTYPE9   string `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `

	// FCENTRYTIM string `gorm:"column:FCENTRYTIM;" json:"fcentrytim"  form:"fcentrytim" `
	// FCFIX      string `gorm:"column:FCFIX;" json:"fcfix"  form:"fcfix" `
	// FCFROMPGL  string `gorm:"column:FCFROMPGL;" json:"fcfrompgl"  form:"fcfrompgl" `
	// FCGLHEADRE string `gorm:"column:FCGLHEADRE;" json:"fcglheadre"  form:"fcglheadre" `
	// FCISBUDGET string `gorm:"column:FCISBUDGET;" json:"fcisbudget"  form:"fcisbudget" `
	// FCLID      string `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `

	// FCLOCKBY   string `gorm:"column:FCLOCKBY;" json:"fclockby"  form:"fclockby" `
	FCLUPDAPP string `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	// FCMANDATE  string `gorm:"column:FCMANDATE;" json:"fcmandate"  form:"fcmandate" `
	// FCORGCODE  string `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	// FCPOSTBY   string `gorm:"column:FCPOSTBY;" json:"fcpostby"  form:"fcpostby" `
	// FCPOSTCORR string `gorm:"column:FCPOSTCORR;" json:"fcpostcorr"  form:"fcpostcorr" `
	// FCRECALBY  string `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
	// FCSELTAG   string `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `

	// FCSRCUPD string `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `

	// FCU1ACC    string `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	// FCU1STATUS string `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	// FCU2STATUS string `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	// FCU3STATUS string `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	// FCU4STATUS string `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	// FCU5STATUS string `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	// FCU6STATUS string `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	// FCU7STATUS string `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	// FCU8STATUS string `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	// FCU9STATUS string `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	// FCUDATE    string `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	// FCUTIME    string `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	// FDAUDIT    string `gorm:"column:FDAUDIT;" json:"fdaudit"  form:"fdaudit" `

	// FDENTRYDAT string `gorm:"column:FDENTRYDAT;" json:"fdentrydat"  form:"fdentrydat" `
	// FDMANDATE  string `gorm:"column:FDMANDATE;" json:"fdmandate"  form:"fdmandate" `
	// FMDOCFLOW  string `gorm:"column:FMDOCFLOW;" json:"fmdocflow"  form:"fmdocflow" `
	// FMERRMSG   string `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	// FMEXTRATAG string `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `

	// FNISAUDIT float64  `gorm:"column:FNISAUDIT;" json:"fnisaudit"  form:"fnisaudit" `

	// FNU1CNT   float64  `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	// FNU2CNT   float64  `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	// FNU3CNT   float64  `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	// FNU4CNT   float64  `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	// FNU5CNT   float64  `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	// FNU6CNT   float64  `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	// FNU7CNT   float64  `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	// FNU8CNT   float64  `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	// FNU9CNT   float64  `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	// FTLOCK    string   `gorm:"column:FTLOCK;" json:"ftlock"  form:"ftlock" `
	// FTSRCUPD  string   `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
}

func (Glhead) TableName() string {
	return "GLHEAD"
}
func (obj *Glhead) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)

	gid, _ := g.New(26)
	obj.FCGID = gid

	obj.FIMILLISEC = time.Now().Unix()
	obj.FCLUPDAPP = "$/"
	obj.FCCREATEAP = "$/"
	obj.FCDATASER = "$$$+"
	obj.FCSTEP = "E"
	obj.FCSTYLE = "A"
	obj.FCLINK = "$C"
	obj.FCEAFTERR = "E"
	obj.FCATSTEP = "I"
	obj.FNLOCKED = 0
	obj.FTDATETIME = time.Now()
	obj.FTLASTEDIT = time.Now()
	obj.FTLASTUPD = time.Now()
	return
}
