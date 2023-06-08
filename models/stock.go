package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Stock struct {
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCWHOUSE   string    `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCLOT      string    `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FNAVGCOST  float64   `gorm:"column:FNAVGCOST;" json:"fnavgcost"  form:"fnavgcost" `
	FNQTY      float64   `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNSTQTYSTD float64   `gorm:"column:FNSTQTYSTD;" json:"fnstqtystd"  form:"fnstqtystd" `
	FNPRICE    float64   `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
	FNEDPRICE  float64   `gorm:"column:FNEDPRICE;" json:"fnedprice"  form:"fnedprice" `
	FNCOST     float64   `gorm:"column:FNCOST;" json:"fncost"  form:"fncost" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FNUNQCQTY  float64   `gorm:"column:FNUNQCQTY;" json:"fnunqcqty"  form:"fnunqcqty" `
	FNUNDOQTY  float64   `gorm:"column:FNUNDOQTY;" json:"fnundoqty"  form:"fnundoqty" `
	FNSOALLOCQ float64   `gorm:"column:FNSOALLOCQ;" json:"fnsoallocq"  form:"fnsoallocq" `
	FCGAG      string    `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
	FDMFGDATE  string    `gorm:"column:FDMFGDATE;" json:"fdmfgdate"  form:"fdmfgdate" `
	FDEXPIRE   string    `gorm:"column:FDEXPIRE;" json:"fdexpire"  form:"fdexpire" `
	FNDOQTY    float64   `gorm:"column:FNDOQTY;" json:"fndoqty"  form:"fndoqty" `
	FNGRNQTY   float64   `gorm:"column:FNGRNQTY;" json:"fngrnqty"  form:"fngrnqty" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FNALLOCQTY float64   `gorm:"column:FNALLOCQTY;" json:"fnallocqty"  form:"fnallocqty" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FDRECDATE  string    `gorm:"column:FDRECDATE;" json:"fdrecdate"  form:"fdrecdate" `
	FCLOCATION string    `gorm:"column:FCLOCATION;" json:"fclocation"  form:"fclocation" `
	FCSTORE    string    `gorm:"column:FCSTORE;" json:"fcstore"  form:"fcstore" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	Product    Product   `gorm:"foreignKey:FCPROD;references:FCSKID;" json:"product"`
	WHouse     Whs       `gorm:"foreignKey:FCWHOUSE;references:FCSKID;" json:"whouse"`
}

func (Stock) TableName() string {
	return "STOCK"
}

func (obj *Stock) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("G%sF", id)
	uid, _ := g.New(26)
	obj.FCGID = uid
	obj.FCDATASER = "$$$+"
	obj.FIMILLISEC = time.Now().Unix()
	obj.FTDATETIME = time.Now()
	obj.FTLASTEDIT = time.Now()
	obj.FTLASTUPD = time.Now()
	obj.FCLUPDAPP = "$0"
	return
}

func (obj *Stock) BeforeUpdate(tx *gorm.DB) (err error) {
	obj.FTLASTEDIT = time.Now()
	obj.FTLASTUPD = time.Now()
	return
}
