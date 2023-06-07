package models

import (
	"fmt"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Corp struct {
	// FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCSKID string `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	// FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	// FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	// FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	// FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCCODE   string `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCNAME   string `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2  string `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCTAXID  string `gorm:"column:FCTAXID;" json:"fctaxid"  form:"fctaxid" `
	FCADDR1  string `gorm:"column:FCADDR1;" json:"fcaddr1"  form:"fcaddr1" `
	FCADDR12 string `gorm:"column:FCADDR12;" json:"fcaddr12"  form:"fcaddr12" `
	FCADDR2  string `gorm:"column:FCADDR2;" json:"fcaddr2"  form:"fcaddr2" `
	FCADDR22 string `gorm:"column:FCADDR22;" json:"fcaddr22"  form:"fcaddr22" `
	FCTEL    string `gorm:"column:FCTEL;" json:"fctel"  form:"fctel" `
	FCFAX    string `gorm:"column:FCFAX;" json:"fcfax"  form:"fcfax" `
	// FDUSERDATE string    `gorm:"column:FDUSERDATE;" json:"fduserdate"  form:"fduserdate" `
	// FDCALDATE  time.Time `gorm:"column:FDCALDATE;" json:"fdcaldate"  form:"fdcaldate" default:"now"`
	// FCBACCDATE string    `gorm:"column:FCBACCDATE;" json:"fcbaccdate"  form:"fcbaccdate" `
	// FCBACCMTH  string    `gorm:"column:FCBACCMTH;" json:"fcbaccmth"  form:"fcbaccmth" `
	// FNROUNDDEP float64   `gorm:"column:FNROUNDDEP;" json:"fnrounddep"  form:"fnrounddep" `
	// FCACCBOOK  string    `gorm:"column:FCACCBOOK;" json:"fcaccbook"  form:"fcaccbook" `
	// FDCPBEGDAT string    `gorm:"column:FDCPBEGDAT;" json:"fdcpbegdat"  form:"fdcpbegdat" `
	// FDCPENDDAT string    `gorm:"column:FDCPENDDAT;" json:"fdcpenddat"  form:"fdcpenddat" `
	// FCBRANCHIS string    `gorm:"column:FCBRANCHIS;" json:"fcbranchis"  form:"fcbranchis" `
	// FCBRANCHI2 string    `gorm:"column:FCBRANCHI2;" json:"fcbranchi2"  form:"fcbranchi2" `
	// FCDEPTIS   string    `gorm:"column:FCDEPTIS;" json:"fcdeptis"  form:"fcdeptis" `
	// FCDEPTIS2  string    `gorm:"column:FCDEPTIS2;" json:"fcdeptis2"  form:"fcdeptis2" `
	// FCSECTIS   string    `gorm:"column:FCSECTIS;" json:"fcsectis"  form:"fcsectis" `
	// FCSECTIS2  string    `gorm:"column:FCSECTIS2;" json:"fcsectis2"  form:"fcsectis2" `
	// FCCTRLSTOC string    `gorm:"column:FCCTRLSTOC;" json:"fcctrlstoc"  form:"fcctrlstoc" `
	// FCGOODSTYP string    `gorm:"column:FCGOODSTYP;" json:"fcgoodstyp"  form:"fcgoodstyp" `
	// FCRAWTYPE  string    `gorm:"column:FCRAWTYPE;" json:"fcrawtype"  form:"fcrawtype" `
	// FCSEMITYPE string    `gorm:"column:FCSEMITYPE;" json:"fcsemitype"  form:"fcsemitype" `
	// FCGOODSCOS string    `gorm:"column:FCGOODSCOS;" json:"fcgoodscos"  form:"fcgoodscos" `
	// FCRAWCOST  string    `gorm:"column:FCRAWCOST;" json:"fcrawcost"  form:"fcrawcost" `
	// FCSEMICOST string    `gorm:"column:FCSEMICOST;" json:"fcsemicost"  form:"fcsemicost" `
	// FCVATISOUT string    `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	// FCVATTYPE  string    `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	// FNVATRATE  float64   `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	// FCSHOWCOMP string    `gorm:"column:FCSHOWCOMP;" json:"fcshowcomp"  form:"fcshowcomp" `
	// FCEDPRICE  string    `gorm:"column:FCEDPRICE;" json:"fcedprice"  form:"fcedprice" `
	// FCEDNAME   string    `gorm:"column:FCEDNAME;" json:"fcedname"  form:"fcedname" `
	// FCEDCOMPO  string    `gorm:"column:FCEDCOMPO;" json:"fcedcompo"  form:"fcedcompo" `
	// FCARNDUE   string    `gorm:"column:FCARNDUE;" json:"fcarndue"  form:"fcarndue" `
	// FCAROVER   string    `gorm:"column:FCAROVER;" json:"fcarover"  form:"fcarover" `
	// FCAPNDUE   string    `gorm:"column:FCAPNDUE;" json:"fcapndue"  form:"fcapndue" `
	// FCAPOVER   string    `gorm:"column:FCAPOVER;" json:"fcapover"  form:"fcapover" `
	// FCAMTPICT  string    `gorm:"column:FCAMTPICT;" json:"fcamtpict"  form:"fcamtpict" `
	// FCVATPICT  string    `gorm:"column:FCVATPICT;" json:"fcvatpict"  form:"fcvatpict" `
	// FCPRICEPIC string    `gorm:"column:FCPRICEPIC;" json:"fcpricepic"  form:"fcpricepic" `
	// FCQTYPICT  string    `gorm:"column:FCQTYPICT;" json:"fcqtypict"  form:"fcqtypict" `
	// FCCORPCHAR string    `gorm:"column:FCCORPCHAR;" json:"fccorpchar"  form:"fccorpchar" `
	// FMMEMDATA  string    `gorm:"column:FMMEMDATA;" json:"fmmemdata"  form:"fmmemdata" `
	// FNSELL1PCT float64   `gorm:"column:FNSELL1PCT;" json:"fnsell1pct"  form:"fnsell1pct" `
	// FNSELL2PCT float64   `gorm:"column:FNSELL2PCT;" json:"fnsell2pct"  form:"fnsell2pct" `
	// FNSELL3PCT float64   `gorm:"column:FNSELL3PCT;" json:"fnsell3pct"  form:"fnsell3pct" `
	// FNSELL4PCT float64   `gorm:"column:FNSELL4PCT;" json:"fnsell4pct"  form:"fnsell4pct" `
	// FCSECURMOD string    `gorm:"column:FCSECURMOD;" json:"fcsecurmod"  form:"fcsecurmod" `
	// FCTYNAMECO string    `gorm:"column:FCTYNAMECO;" json:"fctynameco"  form:"fctynameco" `
	// FCTYNAMEPR string    `gorm:"column:FCTYNAMEPR;" json:"fctynamepr"  form:"fctynamepr" `
	// FNDECLIMIT float64   `gorm:"column:FNDECLIMIT;" json:"fndeclimit"  form:"fndeclimit" `
	// FNDECLIM2  float64   `gorm:"column:FNDECLIM2;" json:"fndeclim2"  form:"fndeclim2" `
	// FNDECLIM3  float64   `gorm:"column:FNDECLIM3;" json:"fndeclim3"  form:"fndeclim3" `
	// FNDECLIM4  float64   `gorm:"column:FNDECLIM4;" json:"fndeclim4"  form:"fndeclim4" `
	// FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	// FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	// FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	// FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	// FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	// FCPLTYPE   string    `gorm:"column:FCPLTYPE;" json:"fcpltype"  form:"fcpltype" `
	// FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	// FCCOMMERCI string    `gorm:"column:FCCOMMERCI;" json:"fccommerci"  form:"fccommerci" `
	// FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	// FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	// FCDEPT     string    `gorm:"column:FCDEPT;" json:"fcdept"  form:"fcdept" `
	// FCJOB      string    `gorm:"column:FCJOB;" json:"fcjob"  form:"fcjob" `
	// FCSWCAT    string    `gorm:"column:FCSWCAT;" json:"fcswcat"  form:"fcswcat" `
	// FCHCODE    string    `gorm:"column:FCHCODE;" json:"fchcode"  form:"fchcode" `
	// FCPROJ     string    `gorm:"column:FCPROJ;" json:"fcproj"  form:"fcproj" `
	// FCPROJIS   string    `gorm:"column:FCPROJIS;" json:"fcprojis"  form:"fcprojis" `
	// FCPROJIS2  string    `gorm:"column:FCPROJIS2;" json:"fcprojis2"  form:"fcprojis2" `
	// FCJOBIS    string    `gorm:"column:FCJOBIS;" json:"fcjobis"  form:"fcjobis" `
	// FCJOBIS2   string    `gorm:"column:FCJOBIS2;" json:"fcjobis2"  form:"fcjobis2" `
	// FCCORPIS   string    `gorm:"column:FCCORPIS;" json:"fccorpis"  form:"fccorpis" `
	// FCCORPIS2  string    `gorm:"column:FCCORPIS2;" json:"fccorpis2"  form:"fccorpis2" `
	// FDBUDGDATE time.Time `gorm:"column:FDBUDGDATE;" json:"fdbudgdate"  form:"fdbudgdate" default:"now"`
	// FCSUBBRIS  string    `gorm:"column:FCSUBBRIS;" json:"fcsubbris"  form:"fcsubbris" `
	// FCSUBBRIS2 string    `gorm:"column:FCSUBBRIS2;" json:"fcsubbris2"  form:"fcsubbris2" `
	// FMMEMDATA2 string    `gorm:"column:FMMEMDATA2;" json:"fmmemdata2"  form:"fmmemdata2" `
	// FCDXCODE   string    `gorm:"column:FCDXCODE;" json:"fcdxcode"  form:"fcdxcode" `
	// FCHOLDER   string    `gorm:"column:FCHOLDER;" json:"fcholder"  form:"fcholder" `
	// FCSTOCKAGE string    `gorm:"column:FCSTOCKAGE;" json:"fcstockage"  form:"fcstockage" `
	// FNBUDYR    float64   `gorm:"column:FNBUDYR;" json:"fnbudyr"  form:"fnbudyr" `
	// FCCURRENCY string    `gorm:"column:FCCURRENCY;" json:"fccurrency"  form:"fccurrency" `
	// FCTRADENO  string    `gorm:"column:FCTRADENO;" json:"fctradeno"  form:"fctradeno" `
	// FCSTDCURR  string    `gorm:"column:FCSTDCURR;" json:"fcstdcurr"  form:"fcstdcurr" `
	// FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	// FCMCTRLSTO string    `gorm:"column:FCMCTRLSTO;" json:"fcmctrlsto"  form:"fcmctrlsto" `
	// FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	// FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" default:"now"`
	// FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	// FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	// FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	// FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	// FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	// FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	// FMERRMSG   string    `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	// FCSTATUS   string    `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	// FDINACTIVE string    `gorm:"column:FDINACTIVE;" json:"fdinactive"  form:"fdinactive" `
	// FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	// FCSDNAME   string    `gorm:"column:FCSDNAME;" json:"fcsdname"  form:"fcsdname" `
	// FCPVKEY    string    `gorm:"column:FCPVKEY;" json:"fcpvkey"  form:"fcpvkey" `
	// FCDSSDNAME string    `gorm:"column:FCDSSDNAME;" json:"fcdssdname"  form:"fcdssdname" `
	// FNCHGPASSW float64   `gorm:"column:FNCHGPASSW;" json:"fnchgpassw"  form:"fnchgpassw" `
	// FNWITHCHG  float64   `gorm:"column:FNWITHCHG;" json:"fnwithchg"  form:"fnwithchg" `
	// FNLOCKDATE float64   `gorm:"column:FNLOCKDATE;" json:"fnlockdate"  form:"fnlockdate" `
	// FCSHOWROP  string    `gorm:"column:FCSHOWROP;" json:"fcshowrop"  form:"fcshowrop" `
	// FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	// FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	// FCBKCORPID string    `gorm:"column:FCBKCORPID;" json:"fcbkcorpid"  form:"fcbkcorpid" `
	// FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	// FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	// FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	// FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	// FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	// FDBGPERIOD string    `gorm:"column:FDBGPERIOD;" json:"fdbgperiod"  form:"fdbgperiod" `
	// FCBGPERIOD string    `gorm:"column:FCBGPERIOD;" json:"fcbgperiod"  form:"fcbgperiod" `
	// FCHOLDCORP string    `gorm:"column:FCHOLDCORP;" json:"fcholdcorp"  form:"fcholdcorp" `
	// FNHSHARES  float64   `gorm:"column:FNHSHARES;" json:"fnhshares"  form:"fnhshares" `
	// FNTSHARES  float64   `gorm:"column:FNTSHARES;" json:"fntshares"  form:"fntshares" `
	// FCFREECORP string    `gorm:"column:FCFREECORP;" json:"fcfreecorp"  form:"fcfreecorp" `
	// FCCUSTSHAR string    `gorm:"column:FCCUSTSHAR;" json:"fccustshar"  form:"fccustshar" `
	// FCSUPPSHAR string    `gorm:"column:FCSUPPSHAR;" json:"fcsuppshar"  form:"fcsuppshar" `
	// FCPDSHARE  string    `gorm:"column:FCPDSHARE;" json:"fcpdshare"  form:"fcpdshare" `
	// FCACCSHARE string    `gorm:"column:FCACCSHARE;" json:"fcaccshare"  form:"fcaccshare" `
	// FNTAXBEG1  float64   `gorm:"column:FNTAXBEG1;" json:"fntaxbeg1"  form:"fntaxbeg1" `
	// FNTAXEND1  float64   `gorm:"column:FNTAXEND1;" json:"fntaxend1"  form:"fntaxend1" `
	// FNTAXRATE1 float64   `gorm:"column:FNTAXRATE1;" json:"fntaxrate1"  form:"fntaxrate1" `
	// FNTAXBEG2  float64   `gorm:"column:FNTAXBEG2;" json:"fntaxbeg2"  form:"fntaxbeg2" `
	// FNTAXEND2  float64   `gorm:"column:FNTAXEND2;" json:"fntaxend2"  form:"fntaxend2" `
	// FNTAXRATE2 float64   `gorm:"column:FNTAXRATE2;" json:"fntaxrate2"  form:"fntaxrate2" `
	// FNTAXBEG3  float64   `gorm:"column:FNTAXBEG3;" json:"fntaxbeg3"  form:"fntaxbeg3" `
	// FNTAXEND3  float64   `gorm:"column:FNTAXEND3;" json:"fntaxend3"  form:"fntaxend3" `
	// FNTAXRATE3 float64   `gorm:"column:FNTAXRATE3;" json:"fntaxrate3"  form:"fntaxrate3" `
	// FNTAXBEG4  float64   `gorm:"column:FNTAXBEG4;" json:"fntaxbeg4"  form:"fntaxbeg4" `
	// FNTAXEND4  float64   `gorm:"column:FNTAXEND4;" json:"fntaxend4"  form:"fntaxend4" `
	// FNTAXRATE4 float64   `gorm:"column:FNTAXRATE4;" json:"fntaxrate4"  form:"fntaxrate4" `
	// FNTAXBEG5  float64   `gorm:"column:FNTAXBEG5;" json:"fntaxbeg5"  form:"fntaxbeg5" `
	// FNTAXEND5  float64   `gorm:"column:FNTAXEND5;" json:"fntaxend5"  form:"fntaxend5" `
	// FNTAXRATE5 float64   `gorm:"column:FNTAXRATE5;" json:"fntaxrate5"  form:"fntaxrate5" `
	// FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	// FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	// FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	// FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	// FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	// FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	// FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	// FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	// FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	// FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	// FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	// FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	// FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	// FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	// FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	// FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	// FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	// FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	// FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	// FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	// FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	// FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	// FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	// FCCNCOSTS  string    `gorm:"column:FCCNCOSTS;" json:"fccncosts"  form:"fccncosts" `
	// FCCNCOSTB  string    `gorm:"column:FCCNCOSTB;" json:"fccncostb"  form:"fccncostb" `
	// FCCOSTZERO string    `gorm:"column:FCCOSTZERO;" json:"fccostzero"  form:"fccostzero" `
	// FNCOSTZERO float64   `gorm:"column:FNCOSTZERO;" json:"fncostzero"  form:"fncostzero" `
	// FCCOSTACNT string    `gorm:"column:FCCOSTACNT;" json:"fccostacnt"  form:"fccostacnt" `
	// FNCOSTACNT float64   `gorm:"column:FNCOSTACNT;" json:"fncostacnt"  form:"fncostacnt" `
	// FCCOSTRETW string    `gorm:"column:FCCOSTRETW;" json:"fccostretw"  form:"fccostretw" `
	// FNCOSTRETW float64   `gorm:"column:FNCOSTRETW;" json:"fncostretw"  form:"fncostretw" `
	// FCPERIODTY string    `gorm:"column:FCPERIODTY;" json:"fcperiodty"  form:"fcperiodty" `
	// FCBOICARD  string    `gorm:"column:FCBOICARD;" json:"fcboicard"  form:"fcboicard" `
	// FCBOIGROUP string    `gorm:"column:FCBOIGROUP;" json:"fcboigroup"  form:"fcboigroup" `
	// FCCOSTBTYP string    `gorm:"column:FCCOSTBTYP;" json:"fccostbtyp"  form:"fccostbtyp" `
	// FCVATREG   string    `gorm:"column:FCVATREG;" json:"fcvatreg"  form:"fcvatreg" `
	// FCCODEONL  string    `gorm:"column:FCCODEONL;" json:"fccodeonl"  form:"fccodeonl" `
	// FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	// FCLOCKTRAN string    `gorm:"column:FCLOCKTRAN;" json:"fclocktran"  form:"fclocktran" `
	// FCNO       string    `gorm:"column:FCNO;" json:"fcno"  form:"fcno" `
	// FCFLOOR    string    `gorm:"column:FCFLOOR;" json:"fcfloor"  form:"fcfloor" `
	// FCROOMNO   string    `gorm:"column:FCROOMNO;" json:"fcroomno"  form:"fcroomno" `
	// FCBUILDING string    `gorm:"column:FCBUILDING;" json:"fcbuilding"  form:"fcbuilding" `
	// FCBUILDIN2 string    `gorm:"column:FCBUILDIN2;" json:"fcbuildin2"  form:"fcbuildin2" `
	// FCMOO      string    `gorm:"column:FCMOO;" json:"fcmoo"  form:"fcmoo" `
	// FCSOI      string    `gorm:"column:FCSOI;" json:"fcsoi"  form:"fcsoi" `
	// FCSOI2     string    `gorm:"column:FCSOI2;" json:"fcsoi2"  form:"fcsoi2" `
	// FCROAD     string    `gorm:"column:FCROAD;" json:"fcroad"  form:"fcroad" `
	// FCROAD2    string    `gorm:"column:FCROAD2;" json:"fcroad2"  form:"fcroad2" `
	// FCQCJANGWA string    `gorm:"column:FCQCJANGWA;" json:"fcqcjangwa"  form:"fcqcjangwa" `
	// FCQNJANGWA string    `gorm:"column:FCQNJANGWA;" json:"fcqnjangwa"  form:"fcqnjangwa" `
	// FCQNJANGW2 string    `gorm:"column:FCQNJANGW2;" json:"fcqnjangw2"  form:"fcqnjangw2" `
	// FCQCAMPHUR string    `gorm:"column:FCQCAMPHUR;" json:"fcqcamphur"  form:"fcqcamphur" `
	// FCQNAMPHUR string    `gorm:"column:FCQNAMPHUR;" json:"fcqnamphur"  form:"fcqnamphur" `
	// FCQNAMPHU2 string    `gorm:"column:FCQNAMPHU2;" json:"fcqnamphu2"  form:"fcqnamphu2" `
	// FCQCTAMBON string    `gorm:"column:FCQCTAMBON;" json:"fcqctambon"  form:"fcqctambon" `
	// FCQNTAMBON string    `gorm:"column:FCQNTAMBON;" json:"fcqntambon"  form:"fcqntambon" `
	// FCQNTAMBO2 string    `gorm:"column:FCQNTAMBO2;" json:"fcqntambo2"  form:"fcqntambo2" `
	// FCZIP      string    `gorm:"column:FCZIP;" json:"fczip"  form:"fczip" `
	// FCEMAIL    string    `gorm:"column:FCEMAIL;" json:"fcemail"  form:"fcemail" `
	// FCQCCOUNTY string    `gorm:"column:FCQCCOUNTY;" json:"fcqccounty"  form:"fcqccounty" `
	// FCSHOW2UM  string    `gorm:"column:FCSHOW2UM;" json:"fcshow2um"  form:"fcshow2um" `
	// FCCHGDUE   string    `gorm:"column:FCCHGDUE;" json:"fcchgdue"  form:"fcchgdue" `
	// FTSRCUPD   time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
	// FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	// FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	// FCADUCFACC string    `gorm:"column:FCADUCFACC;" json:"fcaducfacc"  form:"fcaducfacc" `
	// FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	// FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	// FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
}

func (Corp) TableName() string {
	return "CORP"
}

func (obj *Corp) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
