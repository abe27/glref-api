package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Branch struct {
	// FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	// FCADDR1    string    `gorm:"column:FCADDR1;" json:"fcaddr1"  form:"fcaddr1" `
	// FCADDR12   string    `gorm:"column:FCADDR12;" json:"fcaddr12"  form:"fcaddr12" `
	// FCADDR2    string    `gorm:"column:FCADDR2;" json:"fcaddr2"  form:"fcaddr2" `
	// FCADDR22   string    `gorm:"column:FCADDR22;" json:"fcaddr22"  form:"fcaddr22" `
	// FCADDR3    string    `gorm:"column:FCADDR3;" json:"fcaddr3"  form:"fcaddr3" `
	// FCADDR32   string    `gorm:"column:FCADDR32;" json:"fcaddr32"  form:"fcaddr32" `
	// FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	// FCAPVPOTYP string    `gorm:"column:FCAPVPOTYP;" json:"fcapvpotyp"  form:"fcapvpotyp" `
	// FCAPVPRTYP string    `gorm:"column:FCAPVPRTYP;" json:"fcapvprtyp"  form:"fcapvprtyp" `
	// FCAPVQSTYP string    `gorm:"column:FCAPVQSTYP;" json:"fcapvqstyp"  form:"fcapvqstyp" `
	// FCAPVSITYP string    `gorm:"column:FCAPVSITYP;" json:"fcapvsityp"  form:"fcapvsityp" `
	// FCAPVSOTYP string    `gorm:"column:FCAPVSOTYP;" json:"fcapvsotyp"  form:"fcapvsotyp" `
	// FCBAKDIRNA string    `gorm:"column:FCBAKDIRNA;" json:"fcbakdirna"  form:"fcbakdirna" `
	// FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	// FCBKACCPAY string    `gorm:"column:FCBKACCPAY;" json:"fcbkaccpay"  form:"fcbkaccpay" `
	// FCBPRCTYPE string    `gorm:"column:FCBPRCTYPE;" json:"fcbprctype"  form:"fcbprctype" `
	// FCCHKMPRC  string    `gorm:"column:FCCHKMPRC;" json:"fcchkmprc"  form:"fcchkmprc" `
	FCCODE string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCORP string `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	// FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	// FCCOSTBASE string    `gorm:"column:FCCOSTBASE;" json:"fccostbase"  form:"fccostbase" `
	// FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	// FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	// FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	// FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	// FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	// FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	// FCDEFAPRIN string    `gorm:"column:FCDEFAPRIN;" json:"fcdefaprin"  form:"fcdefaprin" `
	// FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	// FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	// FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	// FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	// FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	// FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	// FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	// FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	// FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	// FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	// FCDXCODE   string    `gorm:"column:FCDXCODE;" json:"fcdxcode"  form:"fcdxcode" `
	// FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	// FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	// FCFAX      string    `gorm:"column:FCFAX;" json:"fcfax"  form:"fcfax" `
	// FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	// FCFRONTAUD string    `gorm:"column:FCFRONTAUD;" json:"fcfrontaud"  form:"fcfrontaud" `
	// FCFRONTLOC string    `gorm:"column:FCFRONTLOC;" json:"fcfrontloc"  form:"fcfrontloc" `
	// FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	// FCGLBILLAP string    `gorm:"column:FCGLBILLAP;" json:"fcglbillap"  form:"fcglbillap" `
	// FCGLBILLAR string    `gorm:"column:FCGLBILLAR;" json:"fcglbillar"  form:"fcglbillar" `
	// FCGLCURR   string    `gorm:"column:FCGLCURR;" json:"fcglcurr"  form:"fcglcurr" `
	// FCHCODE    string    `gorm:"column:FCHCODE;" json:"fchcode"  form:"fchcode" `
	// FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	// FCLBDATASE string    `gorm:"column:FCLBDATASE;" json:"fclbdatase"  form:"fclbdatase" `
	// FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	// FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	// FCMACHINEN string    `gorm:"column:FCMACHINEN;" json:"fcmachinen"  form:"fcmachinen" `
	// FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCNAME  string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2 string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	// FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	// FCPATSTEP  string    `gorm:"column:FCPATSTEP;" json:"fcpatstep"  form:"fcpatstep" `
	// FCPAYINFEE string    `gorm:"column:FCPAYINFEE;" json:"fcpayinfee"  form:"fcpayinfee" `
	// FCPAYINVCT string    `gorm:"column:FCPAYINVCT;" json:"fcpayinvct"  form:"fcpayinvct" `
	// FCPOSTCNTY string    `gorm:"column:FCPOSTCNTY;" json:"fcpostcnty"  form:"fcpostcnty" `
	// FCPPOLIBY  string    `gorm:"column:FCPPOLIBY;" json:"fcppoliby"  form:"fcppoliby" `
	// FCPYMRUNTY string    `gorm:"column:FCPYMRUNTY;" json:"fcpymrunty"  form:"fcpymrunty" `
	// FCREALCUST string    `gorm:"column:FCREALCUST;" json:"fcrealcust"  form:"fcrealcust" `
	// FCREVACCBK string    `gorm:"column:FCREVACCBK;" json:"fcrevaccbk"  form:"fcrevaccbk" `
	// FCSDNAME   string    `gorm:"column:FCSDNAME;" json:"fcsdname"  form:"fcsdname" `
	// FCSECTBILL string    `gorm:"column:FCSECTBILL;" json:"fcsectbill"  form:"fcsectbill" `
	// FCSECTITEM string    `gorm:"column:FCSECTITEM;" json:"fcsectitem"  form:"fcsectitem" `
	// FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	// FCSENDETAX string    `gorm:"column:FCSENDETAX;" json:"fcsendetax"  form:"fcsendetax" `
	// FCSEQ      string    `gorm:"column:FCSEQ;" json:"fcseq"  form:"fcseq" `
	// FCSERVEVAT string    `gorm:"column:FCSERVEVAT;" json:"fcservevat"  form:"fcservevat" `
	// FCSHOWDRCR string    `gorm:"column:FCSHOWDRCR;" json:"fcshowdrcr"  form:"fcshowdrcr" `
	// FCSHOWVCHT string    `gorm:"column:FCSHOWVCHT;" json:"fcshowvcht"  form:"fcshowvcht" `
	FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	// FCSPRCTYPE string    `gorm:"column:FCSPRCTYPE;" json:"fcsprctype"  form:"fcsprctype" `
	// FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	// FCSTATUS   string    `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	// FCSUBBR    string    `gorm:"column:FCSUBBR;" json:"fcsubbr"  form:"fcsubbr" `
	// FCTEL      string    `gorm:"column:FCTEL;" json:"fctel"  form:"fctel" `
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
	// FCVCHBILL  string    `gorm:"column:FCVCHBILL;" json:"fcvchbill"  form:"fcvchbill" `
	// FCVCHCOOR  string    `gorm:"column:FCVCHCOOR;" json:"fcvchcoor"  form:"fcvchcoor" `
	// FCWHBUY    string    `gorm:"column:FCWHBUY;" json:"fcwhbuy"  form:"fcwhbuy" `
	// FCWHSALE   string    `gorm:"column:FCWHSALE;" json:"fcwhsale"  form:"fcwhsale" `
	// FDINACTIVE string    `gorm:"column:FDINACTIVE;" json:"fdinactive"  form:"fdinactive" `
	// FDLASTCLOS string    `gorm:"column:FDLASTCLOS;" json:"fdlastclos"  form:"fdlastclos" `
	// FDLASTDELE string    `gorm:"column:FDLASTDELE;" json:"fdlastdele"  form:"fdlastdele" `
	// FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	// FMERRMSG   string    `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	// FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	// FNPYMRUNL  float64   `gorm:"column:FNPYMRUNL;" json:"fnpymrunl"  form:"fnpymrunl" `
	// FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	// FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	// FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	// FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	// FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	// FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	// FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	// FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	// FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTSRCUPD   time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
}

func (Branch) TableName() string {
	return "BRANCH"
}

func (obj *Branch) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
