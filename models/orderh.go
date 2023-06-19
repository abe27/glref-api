package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Orderh struct {
	// FCACSTEP   string    `gorm:"column:FCACSTEP;" json:"fcacstep"  form:"fcacstep" `
	// FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	// FCAPPROVEB string    `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	// FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	// FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	FCBOOK   string `gorm:"column:FCBOOK;" json:"fcbook"  form:"fcbook" `
	FCBRANCH string `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	// FCCANCELBY string    `gorm:"column:FCCANCELBY;" json:"fccancelby"  form:"fccancelby" `
	// FCCARRIER  string    `gorm:"column:FCCARRIER;" json:"fccarrier"  form:"fccarrier" `
	// FCCARRYTYP string    `gorm:"column:FCCARRYTYP;" json:"fccarrytyp"  form:"fccarrytyp" `
	// FCCARTYPE  string    `gorm:"column:FCCARTYPE;" json:"fccartype"  form:"fccartype" `
	FCCODE string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	// FCCONTACTN string    `gorm:"column:FCCONTACTN;" json:"fccontactn"  form:"fccontactn" `
	FCCOOR     string `gorm:"column:FCCOOR;" json:"fccoor"  form:"fccoor" `
	FCCORP     string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	// FCCOUNTER  string    `gorm:"column:FCCOUNTER;" json:"fccounter"  form:"fccounter" `
	// FCCRAPRVB  string    `gorm:"column:FCCRAPRVB;" json:"fccraprvb"  form:"fccraprvb" `
	FCCREATEAP string `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	// FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	// FCCRSTEP   string    `gorm:"column:FCCRSTEP;" json:"fccrstep"  form:"fccrstep" `
	// FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	// FCCURRENCY string    `gorm:"column:FCCURRENCY;" json:"fccurrency"  form:"fccurrency" `
	// FCCVEHICLE string    `gorm:"column:FCCVEHICLE;" json:"fccvehicle"  form:"fccvehicle" `
	// FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER string `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	// FCDELIBY   string    `gorm:"column:FCDELIBY;" json:"fcdeliby"  form:"fcdeliby" `
	FCDELICOOR string `gorm:"column:FCDELICOOR;" json:"fcdelicoor"  form:"fcdelicoor" `
	// FCDELIEMPL string    `gorm:"column:FCDELIEMPL;" json:"fcdeliempl"  form:"fcdeliempl" `
	FCDEPT string `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	// FCDISCSTR  string    `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	// FCDOCAPRVB string    `gorm:"column:FCDOCAPRVB;" json:"fcdocaprvb"  form:"fcdocaprvb" `
	// FCDOCFLOWI string    `gorm:"column:FCDOCFLOWI;" json:"fcdocflowi"  form:"fcdocflowi" `
	// FCDOCSTEP  string    `gorm:"column:FCDOCSTEP;" json:"fcdocstep"  form:"fcdocstep" `
	// FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	// FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	// FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	// FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	// FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	// FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	// FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	// FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	// FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEAFTERR string `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	// FCEMPL     string    `gorm:"column:FCEMPL;" json:"fcempl"  form:"fcempl" `
	// FCEMZONE   string    `gorm:"column:FCEMZONE;" json:"fcemzone"  form:"fcemzone" `
	// FCFRWHOUSE string    `gorm:"column:FCFRWHOUSE;" json:"fcfrwhouse"  form:"fcfrwhouse" `
	// FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	// FCHASCHAIN string    `gorm:"column:FCHASCHAIN;" json:"fchaschain"  form:"fchaschain" `
	// FCHASRET   string    `gorm:"column:FCHASRET;" json:"fchasret"  form:"fchasret" `
	// FCISCASH   string    `gorm:"column:FCISCASH;" json:"fciscash"  form:"fciscash" `
	FCISPDPART string `gorm:"column:FCISPDPART;" json:"fcispdpart"  form:"fcispdpart" `
	FCJOB      string `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	// FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	// FCLINKAPP1 string    `gorm:"column:FCLINKAPP1;" json:"fclinkapp1"  form:"fclinkapp1" `
	// FCLINKAPP2 string    `gorm:"column:FCLINKAPP2;" json:"fclinkapp2"  form:"fclinkapp2" `
	// FCLINKSTP1 string    `gorm:"column:FCLINKSTP1;" json:"fclinkstp1"  form:"fclinkstp1" `
	// FCLINKSTP2 string    `gorm:"column:FCLINKSTP2;" json:"fclinkstp2"  form:"fclinkstp2" `
	// FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	// FCMSTEP    string    `gorm:"column:FCMSTEP;" json:"fcmstep"  form:"fcmstep" `
	// FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPAYTERM string `gorm:"column:FCPAYTERM;" json:"fcpayterm"  form:"fcpayterm" `
	// FCPDBRAND  string    `gorm:"column:FCPDBRAND;" json:"fcpdbrand"  form:"fcpdbrand" `
	// FCPERSON   string    `gorm:"column:FCPERSON;" json:"fcperson"  form:"fcperson" `
	// FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	// FCPRIORITY string    `gorm:"column:FCPRIORITY;" json:"fcpriority"  form:"fcpriority" `
	FCPROJ string `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	// FCREASONRQ string    `gorm:"column:FCREASONRQ;" json:"fcreasonrq"  form:"fcreasonrq" `
	// FCRECALBY  string    `gorm:"column:FCRECALBY;" json:"fcrecalby"  form:"fcrecalby" `
	// FCRECOMENB string    `gorm:"column:FCRECOMENB;" json:"fcrecomenb"  form:"fcrecomenb" `
	// FCRECVMAN  string    `gorm:"column:FCRECVMAN;" json:"fcrecvman"  form:"fcrecvman" `
	FCREFNO   string `gorm:"column:FCREFNO;" json:"fcrefno"  form:"fcrefno" `
	FCREFTYPE string `gorm:"column:FCREFTYPE;" json:"fcreftype"  form:"fcreftype" `
	// FCREQMEN   string    `gorm:"column:FCREQMEN;" json:"fcreqmen"  form:"fcreqmen" `
	// FCREQPERSN string    `gorm:"column:FCREQPERSN;" json:"fcreqpersn"  form:"fcreqpersn" `
	FCRFTYPE string `gorm:"column:FCRFTYPE;" json:"fcrftype"  form:"fcrftype" `
	// FCRNGTIME  string    `gorm:"column:FCRNGTIME;" json:"fcrngtime"  form:"fcrngtime" `
	// FCROUTEDEL string    `gorm:"column:FCROUTEDEL;" json:"fcroutedel"  form:"fcroutedel" `
	FCSECT string `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	// FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	// FCSHOWPLED string    `gorm:"column:FCSHOWPLED;" json:"fcshowpled"  form:"fcshowpled" `
	FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	// FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	// FCSTAT     string    `gorm:"column:FCSTAT;" json:"fcstat"  form:"fcstat" `
	// FCSTATAPPV string    `gorm:"column:FCSTATAPPV;" json:"fcstatappv"  form:"fcstatappv" `
	FCSTEP string `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	// FCSTEP2    string    `gorm:"column:FCSTEP2;" json:"fcstep2"  form:"fcstep2" `
	// FCSTEPX1   string    `gorm:"column:FCSTEPX1;" json:"fcstepx1"  form:"fcstepx1" `
	// FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	// FCTHRDPTY  string    `gorm:"column:FCTHRDPTY;" json:"fcthrdpty"  form:"fcthrdpty" `
	// FCTOWHOUSE string    `gorm:"column:FCTOWHOUSE;" json:"fctowhouse"  form:"fctowhouse" `
	// FCTRADETRM string    `gorm:"column:FCTRADETRM;" json:"fctradetrm"  form:"fctradetrm" `
	// FCTRANTYPE string    `gorm:"column:FCTRANTYPE;" json:"fctrantype"  form:"fctrantype" `
	// FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	// FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	// FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	// FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	// FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	// FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	// FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	// FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	// FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	// FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	// FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	// FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	// FCVATDUE   string    `gorm:"column:FCVATDUE;" json:"fcvatdue"  form:"fcvatdue" `
	FCVATISOUT string `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATTYPE  string `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	// FCXFERSTEP string    `gorm:"column:FCXFERSTEP;" json:"fcxferstep"  form:"fcxferstep" `
	FDAPPROVE time.Time `gorm:"column:FDAPPROVE;" json:"fdapprove"  form:"fdapprove" default:"now"`
	// FDCRAPRV   string    `gorm:"column:FDCRAPRV;" json:"fdcraprv"  form:"fdcraprv" `
	FDDATE time.Time `gorm:"column:FDDATE;" json:"fddate"  form:"fddate" default:"now"`
	// FDDOCAPRV  string    `gorm:"column:FDDOCAPRV;" json:"fddocaprv"  form:"fddocaprv" `
	FDDUEDATE  time.Time `gorm:"column:FDDUEDATE;" json:"fdduedate"  form:"fdduedate" default:"now"`
	FDRECEDATE time.Time `gorm:"column:FDRECEDATE;" json:"fdrecedate"  form:"fdrecedate" default:"now"`
	// FDRECOMENB string    `gorm:"column:FDRECOMENB;" json:"fdrecomenb"  form:"fdrecomenb" `
	FDREQDATE  time.Time `gorm:"column:FDREQDATE;" json:"fdreqdate"  form:"fdreqdate" default:"now"`
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMDOCFLOW  string    `gorm:"column:FMDOCFLOW;" json:"fmdocflow"  form:"fmdocflow" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	// FMMEMDATA2 string    `gorm:"column:FMMEMDATA2;" json:"fmmemdata2"  form:"fmmemdata2" `
	// FMMEMDATA3 string    `gorm:"column:FMMEMDATA3;" json:"fmmemdata3"  form:"fmmemdata3" `
	// FMMEMDATA4 string    `gorm:"column:FMMEMDATA4;" json:"fmmemdata4"  form:"fmmemdata4" `
	// FMMEMDATA5 string    `gorm:"column:FMMEMDATA5;" json:"fmmemdata5"  form:"fmmemdata5" `
	// FMMEMDATA6 string    `gorm:"column:FMMEMDATA6;" json:"fmmemdata6"  form:"fmmemdata6" `
	// FMMEMDATA7 string    `gorm:"column:FMMEMDATA7;" json:"fmmemdata7"  form:"fmmemdata7" `
	// FMMEMDATA8 string    `gorm:"column:FMMEMDATA8;" json:"fmmemdata8"  form:"fmmemdata8" `
	// FMMEMDATA9 string    `gorm:"column:FMMEMDATA9;" json:"fmmemdata9"  form:"fmmemdata9" `
	// FMMEMDATAA string    `gorm:"column:FMMEMDATAA;" json:"fmmemdataa"  form:"fmmemdataa" `
	FNAMT float64 `gorm:"column:FNAMT;" json:"fnamt"  form:"fnamt" `
	// FNAMT2     float64   `gorm:"column:FNAMT2;" json:"fnamt2"  form:"fnamt2" `
	FNAMTKE    float64 `gorm:"column:FNAMTKE;" json:"fnamtke"  form:"fnamtke" `
	FNCREDTERM int     `gorm:"column:FNCREDTERM;" json:"fncredterm"  form:"fncredterm" `
	// FNDAYASSUR float64   `gorm:"column:FNDAYASSUR;" json:"fndayassur"  form:"fndayassur" `
	// FNDISCAMT1 float64   `gorm:"column:FNDISCAMT1;" json:"fndiscamt1"  form:"fndiscamt1" `
	// FNDISCAMT2 float64   `gorm:"column:FNDISCAMT2;" json:"fndiscamt2"  form:"fndiscamt2" `
	// FNDISCAMTI float64   `gorm:"column:FNDISCAMTI;" json:"fndiscamti"  form:"fndiscamti" `
	// FNDISCAMTK float64   `gorm:"column:FNDISCAMTK;" json:"fndiscamtk"  form:"fndiscamtk" `
	// FNDISCPCN1 float64   `gorm:"column:FNDISCPCN1;" json:"fndiscpcn1"  form:"fndiscpcn1" `
	// FNDISCPCN2 float64   `gorm:"column:FNDISCPCN2;" json:"fndiscpcn2"  form:"fndiscpcn2" `
	// FNDISCPCN3 float64   `gorm:"column:FNDISCPCN3;" json:"fndiscpcn3"  form:"fndiscpcn3" `
	// FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	// FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	// FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	// FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	// FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	// FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	// FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	// FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	// FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNVATAMT   float64   `gorm:"column:FNVATAMT;" json:"fnvatamt"  form:"fnvatamt" `
	FNVATAMTKE float64   `gorm:"column:FNVATAMTKE;" json:"fnvatamtke"  form:"fnvatamtke" `
	FNVATRATE  float64   `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FNXRATE    float64   `gorm:"column:FNXRATE;" json:"fnxrate"  form:"fnxrate" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	// FTLASRECAL string    `gorm:"column:FTLASRECAL;" json:"ftlasrecal"  form:"ftlasrecal" `
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	// FTSRCUPD   string    `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" `
	// FTSTATAPPV string    `gorm:"column:FTSTATAPPV;" json:"ftstatappv"  form:"ftstatappv" `
	// FTXFER     string    `gorm:"column:FTXFER;" json:"ftxfer"  form:"ftxfer" `
	Corp      *Corp       `gorm:"foreignKey:FCCORP;references:FCSKID;" json:"corp"`
	Book      *Booking    `gorm:"foreignKey:FCBOOK;references:FCSKID;" json:"book"`
	Branch    *Branch     `gorm:"foreignKey:FCBRANCH;references:FCSKID;" json:"branch"`
	Dept      *Department `gorm:"foreignKey:FCDEPT;references:FCSKID;" json:"department"`
	Sect      *Section    `gorm:"foreignKey:FCSECT;references:FCSKID;" json:"section"`
	Job       *Job        `gorm:"foreignKey:FCJOB;references:FCSKID;" json:"job"`
	Coor      *Coor       `gorm:"foreignKey:FCCOOR;references:FCSKID;" json:"coor"`
	CreatedBy *Employee   `gorm:"foreignKey:FCCREATEBY;references:FCSKID;" json:"created_by"`
	UpdatedBy *Employee   `gorm:"foreignKey:FCCORRECTB;references:FCSKID;" json:"updated_by"`
	Proj      *Proj       `gorm:"foreignKey:FCPROJ;references:FCSKID;" json:"proj"`
	DeliverTo *Coor       `gorm:"foreignKey:FCDELICOOR;references:FCSKID;" json:"delivery_to"`
	Payterm   *Payterm    `gorm:"foreignKey:FCPAYTERM;references:FCSKID;" json:"paymenterm"`
}

func (Orderh) TableName() string {
	return "ORDERH"
}
func (obj *Orderh) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("G%sF", id)
	return
}
