package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Glref struct {
	FCATSTEP   string    `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
	FCBOOK     string    `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" default:"$$$+"`
	FCDELICOOR string    `gorm:"column:FCDELICOOR;" json:"fcdelicoor"  form:"fcdelicoor" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCFRWHOUSE string    `gorm:"column:FCFRWHOUSE;" json:"fcfrwhouse"  form:"fcfrwhouse" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGLHEAD   string    `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCREFNO    string    `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string    `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"`
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" default:"I"`
	FCTOWHOUSE string    `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	FCVATCOOR  string    `gorm:"column:FCVATCOOR;" json:"fcvatcoor"  form:"fcvatcoor" `
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FNAFTDEP   float64   `gorm:"column:FNAFTDEP;" json:"fnaftdep"  form:"fnaftdep" defualt:"0"`
	FNAFTDEPKE float64   `gorm:"column:FNAFTDEPKE;" json:"fnaftdepke"  form:"fnaftdepke" defualt:"0"`
	FNAMT      float64   `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
}

func (Glref) TableName() string {
	return "GLREF"
}

func (obj *Glref) BeforeCreate(tx *gorm.DB) (err error) {
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

func (obj *Glref) BeforeUpdate(tx *gorm.DB) (err error) {
	obj.FTLASTEDIT = time.Now()
	obj.FTLASTUPD = time.Now()
	return
}

type GlrefForm struct {
	Prefix     string        `json:"prefix"`
	Type       string        `json:"type"`
	Step       string        `json:"step"`
	Branch     string        `json:"branch"`
	Job        string        `json:"job"`
	Corp       string        `json:"corp"`
	Proj       string        `json:"proj"`
	RecDate    time.Time     `json:"recdate"`
	Booking    string        `json:"booking"`
	Whs        string        `json:"whs"`
	Coor       string        `json:"coor"`
	Department string        `json:"department"`
	InvoiceNo  string        `json:"invoice_no"`
	Items      []GlrefDetail `json:"items"`
}

type GlrefDetail struct {
	Product string  `json:"product"`
	Qty     float64 `json:"qty"`
	Unit    string  `json:"unit"`
	Price   float64 `json:"price"`
}
