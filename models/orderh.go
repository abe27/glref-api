package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Orderh struct {
	FCBOOK     string      `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH   string      `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE     string      `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOOR     string      `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string      `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string      `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEAP string      `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string      `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCDATASER  string      `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDELICOOR string      `gorm:"column:FCDELICOOR;" json:"fcdelicoor"  form:"fcdelicoor" `
	FCDEPT     string      `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string      `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCISPDPART string      `gorm:"column:FCISPDPART;" json:"fcispdpart"  form:"fcispdpart" `
	FCJOB      string      `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCPAYTERM  string      `gorm:"column:FCPAYTERM;" json:"fcpayterm"  form:"fcpayterm" `
	FCPROJ     string      `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCREFNO    string      `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string      `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string      `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSECT     string      `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSKID     string      `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSTEP     string      `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCVATISOUT string      `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATTYPE  string      `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FDAPPROVE  time.Time   `gorm:"column:FDAPPROVE;" json:"fdapprove"  form:"fdapprove" default:"now"`
	FDDATE     time.Time   `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FDDUEDATE  time.Time   `gorm:"column:FDDUEDATE;" json:"fdduedate"  form:"fdduedate" default:"now"`
	FDRECEDATE time.Time   `gorm:"column:FDRECEDATE;" json:"fdrecedate"  form:"fdrecedate" default:"now"`
	FDREQDATE  time.Time   `gorm:"column:FDREQDATE;" json:"fdreqdate"  form:"fdreqdate" default:"now"`
	FIMILLISEC int64       `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMDOCFLOW  string      `gorm:"column:FMDOCFLOW;" json:"fmdocflow"  form:"fmdocflow" `
	FMEXTRATAG string      `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMMEMDATA  string      `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FNAMT      float64     `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNAMTKE    float64     `gorm:"column:FNAMTKE;" json:"fnamtke"  form:"fnamtke" `
	FNCREDTERM int         `gorm:"column:FNCREDTERM;" json:"fncredterm"  form:"fncredterm" `
	FNVATAMT   float64     `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FNVATAMTKE float64     `gorm:"column:FNVATAMTKE;" json:"fnvatamtke"  form:"fnvatamtke" `
	FNVATRATE  float64     `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FNXRATE    float64     `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FTDATETIME time.Time   `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time   `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  time.Time   `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	Corp       *Corp       `gorm:"foreignKey:FCCORP;references:FCSKID;" json:"corp"`
	Book       *Booking    `gorm:"foreignKey:FCBOOK;references:FCSKID;" json:"book"`
	Branch     *Branch     `gorm:"foreignKey:FCBRANCH;references:FCSKID;" json:"branch"`
	Dept       *Department `gorm:"foreignKey:FCDEPT;references:FCSKID;" json:"department"`
	Sect       *Section    `gorm:"foreignKey:FCSECT;references:FCSKID;" json:"section"`
	Job        *Job        `gorm:"foreignKey:FCJOB;references:FCSKID;" json:"job"`
	Coor       *Coor       `gorm:"foreignKey:FCCOOR;references:FCSKID;" json:"coor"`
	CreatedBy  *Employee   `gorm:"foreignKey:FCCREATEBY;references:FCSKID;" json:"created_by"`
	UpdatedBy  *Employee   `gorm:"foreignKey:FCCORRECTB;references:FCSKID;" json:"updated_by"`
	Proj       *Proj       `gorm:"foreignKey:FCPROJ;references:FCSKID;" json:"proj"`
	DeliverTo  *Coor       `gorm:"foreignKey:FCDELICOOR;references:FCSKID;" json:"delivery_to"`
	Payterm    *Payterm    `gorm:"foreignKey:FCPAYTERM;references:FCSKID;" json:"paymenterm"`
}

func (Orderh) TableName() string {
	return "ORDERH"
}
func (obj *Orderh) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("G%sF", id)
	return
}
