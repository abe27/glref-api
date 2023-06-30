package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Glref struct {
	FCATSTEP   string      `gorm:"column:FCATSTEP;" json:"fcatstep"  form:"fcatstep" `
	FCBOOK     string      `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH   string      `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCCODE     string      `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOOR     string      `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string      `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string      `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCREATEBY string      `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string      `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCDATASER  string      `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" default:"$$$+"`
	FCDELICOOR string      `gorm:"column:FCDELICOOR;" json:"fcdelicoor"  form:"fcdelicoor" `
	FCDEPT     string      `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCEAFTERR  string      `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCFRWHOUSE string      `gorm:"column:FCFRWHOUSE;" json:"fcfrwhouse"  form:"fcfrwhouse" `
	FCGID      string      `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCGLHEAD   string      `gorm:"column:FCGLHEAD;" json:"fcglhead"  form:"fcglhead" `
	FCJOB      string      `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCPROJ     string      `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCREFNO    string      `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE  string      `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCRFTYPE   string      `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	FCSECT     string      `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSKID     string      `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"`
	FCSTEP     string      `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" default:"I"`
	FCTOWHOUSE string      `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	FCVATCOOR  string      `gorm:"column:FCVATCOOR;" json:"fcvatcoor"  form:"fcvatcoor" `
	FDDATE     time.Time   `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FDRECEDATE time.Time   `gorm:"column:FDRECEDATE;" json:"fdrecedate"`
	FDDUEDATE  time.Time   `gorm:"column:FDDUEDATE;" json:"fdduedate"`
	FIMILLISEC int64       `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FNAFTDEP   float64     `gorm:"column:FNAFTDEP;" json:"fnaftdep"  form:"fnaftdep" defualt:"0"`
	FNAFTDEPKE float64     `gorm:"column:FNAFTDEPKE;" json:"fnaftdepke"  form:"fnaftdepke" defualt:"0"`
	FNAMT      float64     `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	FNXRATE    float64     `gorm:"column:FNXRATE;" json:"fnxate"`
	FNAMTKE    float64     `gorm:"column:FNAMTKE;" json:"fnamtke"`
	FNVATAMTKE float64     `gorm:"column:FNVATAMTKE;" json:"fnvatamtke"`
	FNCREDTERM int         `gorm:"column:FNCREDTERM;" json:"fncredterm"`
	FCVATTYPE  string      `gorm:"column:FCVATTYPE;" json:"fcvattype"`
	FNVATRATE  float64     `gorm:"column:FNVATRATE;" json:"fnvatrate"`
	FNVATAMT   float64     `gorm:"column:FNVATAMT;" json:"fnvatamt"`
	FNPAYAMT   float64     `gorm:"column:FNPAYAMT;" json:"fnpayamt"`
	FNSPAYAMT  float64     `gorm:"column:FNSPAYAMT;" json:"fnspayamt"`
	FNBEFOAMT  float64     `gorm:"column:FNBEFOAMT;" json:"fnbefoamt"`
	FCVATISOUT string      `gorm:"column:FCVATISOUT;" json:"fcvatisout"`
	FCISCASH   string      `gorm:"column:FCISCASH;" json:"fciscash"`
	FCHASRET   string      `gorm:"column:FCHASRET;" json:"fcchasret"`
	FCVATDUE   string      `gorm:"column:FCVATDUE;" json:"fcvatdue"`
	FNSTOCKUPD float64     `gorm:"column:FNSTOCKUPD;" json:"fnstockupd"`
	FTDATETIME time.Time   `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time   `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time   `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FCLUPDAPP  string      `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FMMEMDATA  string      `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	FCCREATEAP string      `gorm:"column:FCCREATEAP;" json:"fccreateap"`
	BOOK       *Booking    `gorm:"foreignKey:FCBOOK;references:FCSKID;" json:"book"`
	BRANCH     *Branch     `gorm:"foreignKey:FCBRANCH;references:FCSKID;" json:"branch"`
	COOR       *Coor       `gorm:"foreignKey:FCCOOR;references:FCSKID;" json:"coor"`
	CORP       *Corp       `gorm:"foreignKey:FCCORP;references:FCSKID;" json:"corp"`
	CORRECTB   *Employee   `gorm:"foreignKey:FCCORRECTB;references:FCSKID;" json:"correctb"`
	DEPT       *Department `gorm:"foreignKey:FCDEPT;references:FCSKID;" json:"department"`
	JOB        *Job        `gorm:"foreignKey:FCJOB;references:FCSKID;" json:"job"`
	PROJ       *Proj       `gorm:"foreignKey:FCPROJ;references:FCSKID;" json:"proj"`
	TOWHOUSE   *Whs        `gorm:"foreignKey:FCTOWHOUSE;references:FCSKID;" json:"to_whouse"`
	FROMWHOUSE *Whs        `gorm:"foreignKey:FCFRWHOUSE;references:FCSKID;" json:"from_whouse"`
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
	obj.FCSTEP = "1"
	obj.FCLUPDAPP = "$0"
	obj.FCVATTYPE = "1"
	obj.FNVATRATE = 7
	obj.FNVATAMT = 0
	obj.FNCREDTERM = 0
	obj.FCISCASH = "N"
	obj.FCHASRET = "Y"
	obj.FCVATDUE = "Y"
	obj.FCEAFTERR = "E"
	obj.FNPAYAMT = 0
	obj.FNSPAYAMT = 0
	obj.FNBEFOAMT = 0
	obj.FCVATISOUT = "Y"
	obj.FNXRATE = 1
	obj.FNAMTKE = 0
	obj.FNVATAMTKE = 0
	obj.FNSTOCKUPD = 1
	obj.FCCREATEAP = "$/"
	obj.FDRECEDATE = time.Now()
	obj.FDDUEDATE = time.Now()
	obj.FIMILLISEC = time.Now().Unix()
	obj.FTDATETIME = time.Now()
	obj.FTLASTEDIT = time.Now()
	obj.FTLASTUPD = time.Now()
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
	FromWhs    string        `json:"from_whs"`
	ToWhs      string        `json:"to_whs"`
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

type GlrefHistory struct {
	FCSKID    string    `gorm:"primaryKey;size:8;unique;index;" json:"fcskid"`
	FCCODE    string    `gorm:"size:8;unique;index;" json:"fccode"`
	FCREFNO   string    `gorm:"size:25;" json:"fcrefno"`
	FCDATE    time.Time `json:"fcdate"`
	FCINVOICE string    `json:"fcinvoice"`
	FCPONO    string    `json:"fcpono"`
	FCREMARK  string    `json:"fcremark"`
	FCSTATUS  int       `json:"fcstatus"`
	FCTYPE    string    `gorm:"size:8;not null;" json:"fctype"`
	CreatedAt time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *GlrefHistory) BeforeCreate(tx *gorm.DB) (err error) {
	obj.FCSTATUS = 0
	return
}
