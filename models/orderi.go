package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Orderi struct {
	FCACCHART  string    `gorm:"column:FCACCHART;" json:"fcacchart"  form:"fcacchart" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCAPPROVEB string    `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCBUDCHART string    `gorm:"column:FCBUDCHART;" json:"fcbudchart"  form:"fcbudchart" `
	FCCOOR     string    `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCOUNTER  string    `gorm:"column:FCCOUNTER;" json:"fccounter"  form:"fccounter" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCTRLSTOC string    `gorm:"column:FCCTRLSTOC;" json:"fcctrlstoc"  form:"fcctrlstoc" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDELITIME string    `gorm:"column:FCDELITIME;" json:"fcdelitime"  form:"fcdelitime" `
	FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	FCDISCSTR  string    `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCFORMULAS string    `gorm:"column:FCFORMULAS;" json:"fcformulas"  form:"fcformulas" `
	FCGAG      string    `gorm:"column:FCGAG;" json:"fcgag"  form:"fcgag" `
	FCGVPOLICY string    `gorm:"column:FCGVPOLICY;" json:"fcgvpolicy"  form:"fcgvpolicy" `
	FCIOTYPE   string    `gorm:"column:FCIOTYPE;" json:"fciotype"  form:"fciotype" `
	FCISCHG    string    `gorm:"column:FCISCHG;" json:"fcischg"  form:"fcischg" `
	FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLOT      string    `gorm:"column:FCLOT;" json:"fclot"  form:"fclot" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMAINJOBH string    `gorm:"column:FCMAINJOBH;" json:"fcmainjobh"  form:"fcmainjobh" `
	FCMSTEP    string    `gorm:"column:FCMSTEP;" json:"fcmstep"  form:"fcmstep" `
	FCORDERH   string    `gorm:"column:FCORDERH;" json:"fcorderh"  form:"fcorderh" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPDSTHPAC string    `gorm:"column:FCPDSTHPAC;" json:"fcpdsthpac"  form:"fcpdsthpac" `
	FCPFORMULA string    `gorm:"column:FCPFORMULA;" json:"fcpformula"  form:"fcpformula" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPRICE    string    `gorm:"column:FCPRICE;" json:"fcprice"  form:"fcprice" `
	FCPRICEBY  string    `gorm:"column:FCPRICEBY;" json:"fcpriceby"  form:"fcpriceby" `
	FCPRIORITY string    `gorm:"column:FCPRIORITY;" json:"fcpriority"  form:"fcpriority" `
	FCPROCHAIN string    `gorm:"column:FCPROCHAIN;" json:"fcprochain"  form:"fcprochain" `
	FCPROD     string    `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCPRODTYPE string    `gorm:"column:FCPRODTYPE;" json:"fcprodtype"  form:"fcprodtype" `
	FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	FCRECALBY  string    `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
	FCREFPDTYP string    `gorm:"column:FCREFPDTYP;" json:"fcrefpdtyp"  form:"fcrefpdtyp" `
	FCREFTYPE  string    `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	FCREMLOCAT string    `gorm:"column:FCREMLOCAT;" json:"fcremlocat"  form:"fcremlocat" `
	FCROOTSEQ  string    `gorm:"column:FCROOTSEQ;" json:"fcrootseq"  form:"fcrootseq" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSHOWCOMP string    `gorm:"column:FCSHOWCOMP;" json:"fcshowcomp"  form:"fcshowcomp" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSTEP2    string    `gorm:"column:FCSTEP2;" json:"fcstep2"  form:"fcstep2" `
	FCSTORE    string    `gorm:"column:FCSTORE;" json:"fcstore"  form:"fcstore" `
	FCSTUM     string    `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
	FCSTUMSTD  string    `gorm:"column:FCSTUMSTD;" json:"fcstumstd"  form:"fcstumstd" `
	FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	FCSUBJOBH  string    `gorm:"column:FCSUBJOBH;" json:"fcsubjobh"  form:"fcsubjobh" `
	FCSUBTXID  string    `gorm:"column:FCSUBTXID;" json:"fcsubtxid"  form:"fcsubtxid" `
	FCTXID     string    `gorm:"column:FCTXID;" json:"fctxid"  form:"fctxid" `
	FCTXIDLINK string    `gorm:"column:FCTXIDLINK;" json:"fctxidlink"  form:"fctxidlink" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUM       string    `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUMSTD    string    `gorm:"column:FCUMSTD;" json:"fcumstd"  form:"fcumstd" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVATISOUT string    `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATTYPE  string    `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FCWHOUSE   string    `gorm:"column:FCWHOUSE;" json:"fcwhouse"  form:"fcwhouse" `
	FDAPPROVE  time.Time `gorm:"column:FDAPPROVE;" json:"fdapprove"  form:"fdapprove" default:"now"`
	FDDATE     time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	FDDELIVERY time.Time `gorm:"column:FDDELIVERY;" json:"fddelivery"  form:"fddelivery" default:"now"`
	FDEXPIRE   string    `gorm:"column:FDEXPIRE;" json:"fdexpire"  form:"fdexpire" `
	FDMFGDATE  string    `gorm:"column:FDMFGDATE;" json:"fdmfgdate"  form:"fdmfgdate" `
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMERRMSG   string    `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMREMARK   string    `gorm:"column:FMREMARK;" json:"fmremark"  form:"fmremark" `
	FMREMARK10 string    `gorm:"column:FMREMARK10;" json:"fmremark10"  form:"fmremark10" `
	FMREMARK2  string    `gorm:"column:FMREMARK2;" json:"fmremark2"  form:"fmremark2" `
	FMREMARK3  string    `gorm:"column:FMREMARK3;" json:"fmremark3"  form:"fmremark3" `
	FMREMARK4  string    `gorm:"column:FMREMARK4;" json:"fmremark4"  form:"fmremark4" `
	FMREMARK5  string    `gorm:"column:FMREMARK5;" json:"fmremark5"  form:"fmremark5" `
	FMREMARK6  string    `gorm:"column:FMREMARK6;" json:"fmremark6"  form:"fmremark6" `
	FMREMARK7  string    `gorm:"column:FMREMARK7;" json:"fmremark7"  form:"fmremark7" `
	FMREMARK8  string    `gorm:"column:FMREMARK8;" json:"fmremark8"  form:"fmremark8" `
	FMREMARK9  string    `gorm:"column:FMREMARK9;" json:"fmremark9"  form:"fmremark9" `
	FMTEXTKE   string    `gorm:"column:FMTEXTKE;" json:"fmtextke"  form:"fmtextke" `
	FNBACKQTY  float64   `gorm:"column:FNBACKQTY;" json:"fnbackqty"  form:"fnbackqty" `
	FNDEFAPRIC float64   `gorm:"column:FNDEFAPRIC;" json:"fndefapric"  form:"fndefapric" `
	FNDISCAMT  float64   `gorm:"column:FNDISCAMT;" json:"fndiscamt"  form:"fndiscamt" `
	FNDISCAMTK float64   `gorm:"column:FNDISCAMTK;" json:"fndiscamtk"  form:"fndiscamtk" `
	FNDISCPCN  float64   `gorm:"column:FNDISCPCN;" json:"fndiscpcn"  form:"fndiscpcn" `
	FNDOQTY    float64   `gorm:"column:FNDOQTY;" json:"fndoqty"  form:"fndoqty" `
	FNPACKQTY  float64   `gorm:"column:FNPACKQTY;" json:"fnpackqty"  form:"fnpackqty" `
	FNPLANQTY  float64   `gorm:"column:FNPLANQTY;" json:"fnplanqty"  form:"fnplanqty" `
	FNPRICE    float64   `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
	FNPRICEKE  float64   `gorm:"column:FNPRICEKE;" json:"fnpriceke"  form:"fnpriceke" `
	FNPRQTY    float64   `gorm:"column:FNPRQTY;" json:"fnprqty"  form:"fnprqty" `
	FNQTY      float64   `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNSTQTY    float64   `gorm:"column:FNSTQTY;" json:"fnstqty"  form:"fnstqty" `
	FNSTUMQTY  float64   `gorm:"column:FNSTUMQTY;" json:"fnstumqty"  form:"fnstumqty" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNUMQTY    float64   `gorm:"column:FNUMQTY;" json:"fnumqty"  form:"fnumqty" `
	FNVATAMT   float64   `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FNVATRATE  float64   `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FNWTAXAMT  float64   `gorm:"column:FNWTAXAMT;" json:"fnwtaxamt"  form:"fnwtaxamt" `
	FNWTAXAMTK float64   `gorm:"column:FNWTAXAMTK;" json:"fnwtaxamtk"  form:"fnwtaxamtk" `
	FNWTAXRATE float64   `gorm:"column:FNWTAXRATE;" json:"fnwtaxrate"  form:"fnwtaxrate" `
	FNXRATE    float64   `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASRECAL string    `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	// Corp       *Corp      `gorm:"foreignKey:FCCORP;references:FCSKID;" json:"corp"`
	// Branch     *Branch    `gorm:"foreignKey:FCBRANCH;references:FCSKID;" json:"branch"`
	// Dept       *Dept      `gorm:"foreignKey:FCDEPT;references:FCSKID;" json:"department"`
	// Sect       *Sect      `gorm:"foreignKey:FCSECT;references:FCSKID;" json:"section"`
	// Job        *Job       `gorm:"foreignKey:FCJOB;references:FCSKID;" json:"job"`
	// Orderh     *Orderh    `gorm:"foreignKey:FCORDERH;references:FCSKID;" json:"order_head"`
	// Coor       *Coor      `gorm:"foreignKey:FCCOOR;references:FCSKID;" json:"coor"`
	// Product    *Product `gorm:"foreignKey:FCPROD;references:FCSKID;" json:"product"`
	// Unit       *Unit      `gorm:"foreignKey:FCUM;references:FCSKID;" json:"unit"`
	// UnitSTD    *Unit      `gorm:"foreignKey:FCUMSTD;references:FCSKID;" json:"unit_standard"`
	// WHouse     *WHouse    `gorm:"foreignKey:FCWHOUSE;references:FCSKID;" json:"whouse"`
	// Proj       *Proj      `gorm:"foreignKey:FCPROJ;references:FCSKID;" json:"proj"`
	// Stum       *Unit      `gorm:"foreignKey:FCSTUM;references:FCSK;" json:"stum"`
	// StumStd    *Unit      `gorm:"foreignKey:FCSTUMSTD;references:FCSK;" json:"stum_std"`
	// CreatedBy  *Empl      `gorm:"foreignKey:FCCREATEBY;references:FCSKID;" json:"created_by"`
}

func (Orderi) TableName() string {
	return "ORDERI"
}

func (obj *Orderi) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("G%sF", id)
	return
}

type OrderiView struct {
	FCORDERH     string  `gorm:"column:FCORDERH;" json:"fcorderh"  form:"fcorderh" `
	FCPROD       string  `gorm:"column:FCPROD;" json:"fcprod"  form:"fcprod" `
	FCSEQ        string  `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	FCSKID       string  `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FNBACKQTY    float64 `gorm:"column:FNBACKQTY;" json:"fnbackqty"  form:"fnbackqty" `
	FNQTY        float64 `gorm:"column:FNQTY;" json:"fnqty"  form:"fnqty" `
	FNRECEIVEQTY float64 `gorm:"column:FNRECEIVEQTY;" json:"fnreceiveqty"`
}

func (OrderiView) TableName() string {
	return "ORDERI"
}
