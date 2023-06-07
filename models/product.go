package models

import (
	"fmt"
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Product struct {
	FCACCBCASH string    `gorm:"column:FCACCBCASH;" json:"fcaccbcash"  form:"fcaccbcash" `
	FCACCBCRED string    `gorm:"column:FCACCBCRED;" json:"fcaccbcred"  form:"fcaccbcred" `
	FCACCCONSP string    `gorm:"column:FCACCCONSP;" json:"fcaccconsp"  form:"fcaccconsp" `
	FCACCEXP1  string    `gorm:"column:FCACCEXP1;" json:"fcaccexp1"  form:"fcaccexp1" `
	FCACCITEM  string    `gorm:"column:FCACCITEM;" json:"fcaccitem"  form:"fcaccitem" `
	FCACCSCASH string    `gorm:"column:FCACCSCASH;" json:"fcaccscash"  form:"fcaccscash" `
	FCACCSCOST string    `gorm:"column:FCACCSCOST;" json:"fcaccscost"  form:"fcaccscost" `
	FCACCSCRED string    `gorm:"column:FCACCSCRED;" json:"fcaccscred"  form:"fcaccscred" `
	FCACCVER   string    `gorm:"column:FCACCVER;" json:"fcaccver"  form:"fcaccver" `
	FCACCWIP   string    `gorm:"column:FCACCWIP;" json:"fcaccwip"  form:"fcaccwip" `
	FCACCWR1   string    `gorm:"column:FCACCWR1;" json:"fcaccwr1"  form:"fcaccwr1" `
	FCACCWR2   string    `gorm:"column:FCACCWR2;" json:"fcaccwr2"  form:"fcaccwr2" `
	FCADDAPVBY string    `gorm:"column:FCADDAPVBY;" json:"fcaddapvby"  form:"fcaddapvby" `
	FCAPPNAME  string    `gorm:"column:FCAPPNAME;" json:"fcappname"  form:"fcappname" `
	FCAPPROVEB string    `gorm:"column:FCAPPROVEB;" json:"fcapproveb"  form:"fcapproveb" `
	FCAUTOGW1  string    `gorm:"column:FCAUTOGW1;" json:"fcautogw1"  form:"fcautogw1" `
	FCBAKYRHIS string    `gorm:"column:FCBAKYRHIS;" json:"fcbakyrhis"  form:"fcbakyrhis" `
	FCBARCODE  string    `gorm:"column:FCBARCODE;" json:"fcbarcode"  form:"fcbarcode" `
	FCBRANCH   string    `gorm:"column:FCBRANCH;" json:"fcbranch"  form:"fcbranch" `
	FCBUDCHART string    `gorm:"column:FCBUDCHART;" json:"fcbudchart"  form:"fcbudchart" `
	FCBVATISOU string    `gorm:"column:FCBVATISOU;" json:"fcbvatisou"  form:"fcbvatisou" `
	FCBWHOUSE  string    `gorm:"column:FCBWHOUSE;" json:"fcbwhouse"  form:"fcbwhouse" `
	FCCODE     string    `gorm:"column:FCCODE;" json:"fccode"  form:"fccode" `
	FCCOLSEND  string    `gorm:"column:FCCOLSEND;" json:"fccolsend"  form:"fccolsend" `
	FCCOPYREMA string    `gorm:"column:FCCOPYREMA;" json:"fccopyrema"  form:"fccopyrema" `
	FCCORP     string    `gorm:"column:FCCORP;" json:"fccorp"  form:"fccorp" `
	FCCORRECTB string    `gorm:"column:FCCORRECTB;" json:"fccorrectb"  form:"fccorrectb" `
	FCCOSTBASE string    `gorm:"column:FCCOSTBASE;" json:"fccostbase"  form:"fccostbase" `
	FCCREATEAP string    `gorm:"column:FCCREATEAP;" json:"fccreateap"  form:"fccreateap" `
	FCCREATEBY string    `gorm:"column:FCCREATEBY;" json:"fccreateby"  form:"fccreateby" `
	FCCREATETY string    `gorm:"column:FCCREATETY;" json:"fccreatety"  form:"fccreatety" `
	FCCTRLSTOC string    `gorm:"column:FCCTRLSTOC;" json:"fcctrlstoc"  form:"fcctrlstoc" `
	FCCUACC    string    `gorm:"column:FCCUACC;" json:"fccuacc"  form:"fccuacc" `
	FCCURR1    string    `gorm:"column:FCCURR1;" json:"fccurr1"  form:"fccurr1" `
	FCCURR2    string    `gorm:"column:FCCURR2;" json:"fccurr2"  form:"fccurr2" `
	FCCURRENCY string    `gorm:"column:FCCURRENCY;" json:"fccurrency"  form:"fccurrency" `
	FCCUST     string    `gorm:"column:FCCUST;" json:"fccust"  form:"fccust" `
	FCDATAIMP  string    `gorm:"column:FCDATAIMP;" json:"fcdataimp"  form:"fcdataimp" `
	FCDATASER  string    `gorm:"column:FCDATASER;" json:"fcdataser"  form:"fcdataser" `
	FCDELAPVBY string    `gorm:"column:FCDELAPVBY;" json:"fcdelapvby"  form:"fcdelapvby" `
	FCDFAUMRAT string    `gorm:"column:FCDFAUMRAT;" json:"fcdfaumrat"  form:"fcdfaumrat" `
	FCDISCSTR  string    `gorm:"column:FCDISCSTR;" json:"fcdiscstr"  form:"fcdiscstr" `
	FCDTYPE1   string    `gorm:"column:FCDTYPE1;" json:"fcdtype1"  form:"fcdtype1" `
	FCDTYPE2   string    `gorm:"column:FCDTYPE2;" json:"fcdtype2"  form:"fcdtype2" `
	FCDTYPE3   string    `gorm:"column:FCDTYPE3;" json:"fcdtype3"  form:"fcdtype3" `
	FCDTYPE4   string    `gorm:"column:FCDTYPE4;" json:"fcdtype4"  form:"fcdtype4" `
	FCDTYPE5   string    `gorm:"column:FCDTYPE5;" json:"fcdtype5"  form:"fcdtype5" `
	FCDTYPE6   string    `gorm:"column:FCDTYPE6;" json:"fcdtype6"  form:"fcdtype6" `
	FCDTYPE7   string    `gorm:"column:FCDTYPE7;" json:"fcdtype7"  form:"fcdtype7" `
	FCDTYPE8   string    `gorm:"column:FCDTYPE8;" json:"fcdtype8"  form:"fcdtype8" `
	FCDTYPE9   string    `gorm:"column:FCDTYPE9;" json:"fcdtype9"  form:"fcdtype9" `
	FCEAFTERR  string    `gorm:"column:FCEAFTERR;" json:"fceafterr"  form:"fceafterr" `
	FCEDPRICE  string    `gorm:"column:FCEDPRICE;" json:"fcedprice"  form:"fcedprice" `
	FCEDTAPVBY string    `gorm:"column:FCEDTAPVBY;" json:"fcedtapvby"  form:"fcedtapvby" `
	FCEXPCUR1  string    `gorm:"column:FCEXPCUR1;" json:"fcexpcur1"  form:"fcexpcur1" `
	FCEXPCUR2  string    `gorm:"column:FCEXPCUR2;" json:"fcexpcur2"  form:"fcexpcur2" `
	FCFCHR     string    `gorm:"column:FCFCHR;" json:"fcfchr"  form:"fcfchr" `
	FCFCHRS    string    `gorm:"column:FCFCHRS;" json:"fcfchrs"  form:"fcfchrs" `
	FCFORMULAS string    `gorm:"column:FCFORMULAS;" json:"fcformulas"  form:"fcformulas" `
	FCGID      string    `gorm:"column:FCGID;" json:"fcgid"  form:"fcgid" `
	FCHASSALE  string    `gorm:"column:FCHASSALE;" json:"fchassale"  form:"fchassale" `
	FCISUSED   string    `gorm:"column:FCISUSED;" json:"fcisused"  form:"fcisused" `
	FCIUM      string    `gorm:"column:FCIUM;" json:"fcium"  form:"fcium" `
	FCLASTUM   string    `gorm:"column:FCLASTUM;" json:"fclastum"  form:"fclastum" `
	FCLID      string    `gorm:"column:FCLID;" json:"fclid"  form:"fclid" `
	FCLOCATION string    `gorm:"column:FCLOCATION;" json:"fclocation"  form:"fclocation" `
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP;" json:"fclupdapp"  form:"fclupdapp" `
	FCMANFLAG  string    `gorm:"column:FCMANFLAG;" json:"fcmanflag"  form:"fcmanflag" `
	FCMANUFAC  string    `gorm:"column:FCMANUFAC;" json:"fcmanufac"  form:"fcmanufac" `
	FCMCTRLSTO string    `gorm:"column:FCMCTRLSTO;" json:"fcmctrlsto"  form:"fcmctrlsto" `
	FCMEASURE  string    `gorm:"column:FCMEASURE;" json:"fcmeasure"  form:"fcmeasure" `
	FCMFUM     string    `gorm:"column:FCMFUM;" json:"fcmfum"  form:"fcmfum" `
	FCMRPXFER  string    `gorm:"column:FCMRPXFER;" json:"fcmrpxfer"  form:"fcmrpxfer" `
	FCNAME     string    `gorm:"column:FCNAME;" json:"fcname"  form:"fcname" `
	FCNAME2    string    `gorm:"column:FCNAME2;" json:"fcname2"  form:"fcname2" `
	FCORGCODE  string    `gorm:"column:FCORGCODE;" json:"fcorgcode"  form:"fcorgcode" `
	FCPDCAT    string    `gorm:"column:FCPDCAT;" json:"fcpdcat"  form:"fcpdcat" `
	FCPDGRP    string    `gorm:"column:FCPDGRP;" json:"fcpdgrp"  form:"fcpdgrp" `
	FCPLANT    string    `gorm:"column:FCPLANT;" json:"fcplant"  form:"fcplant" `
	FCPRCGRP   string    `gorm:"column:FCPRCGRP;" json:"fcprcgrp"  form:"fcprcgrp" `
	FCPROCURE  string    `gorm:"column:FCPROCURE;" json:"fcprocure"  form:"fcprocure" `
	FCPRTYPE   string    `gorm:"column:FCPRTYPE;" json:"fcprtype"  form:"fcprtype" `
	FCSECT     string    `gorm:"column:FCSECT;" json:"fcsect"  form:"fcsect" `
	FCSELTAG   string    `gorm:"column:FCSELTAG;" json:"fcseltag"  form:"fcseltag" `
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;size:8;unique;index;" json:"fcskid"  form:"fcskid" `
	FCSNAME    string    `gorm:"column:FCSNAME;" json:"fcsname"  form:"fcsname" `
	FCSNAME2   string    `gorm:"column:FCSNAME2;" json:"fcsname2"  form:"fcsname2" `
	FCSOURCE   string    `gorm:"column:FCSOURCE;" json:"fcsource"  form:"fcsource" `
	FCSRCUPD   string    `gorm:"column:FCSRCUPD;" json:"fcsrcupd"  form:"fcsrcupd" `
	FCSTATUS   string    `gorm:"column:FCSTATUS;" json:"fcstatus"  form:"fcstatus" `
	FCSTCTRLST string    `gorm:"column:FCSTCTRLST;" json:"fcstctrlst"  form:"fcstctrlst" `
	FCSTEP     string    `gorm:"column:FCSTEP;" json:"fcstep"  form:"fcstep" `
	FCSTIUM    string    `gorm:"column:FCSTIUM;" json:"fcstium"  form:"fcstium" `
	FCSTUM     string    `gorm:"column:FCSTUM;" json:"fcstum"  form:"fcstum" `
	FCSTUM1    string    `gorm:"column:FCSTUM1;" json:"fcstum1"  form:"fcstum1" `
	FCSTUM2    string    `gorm:"column:FCSTUM2;" json:"fcstum2"  form:"fcstum2" `
	FCSTUM3    string    `gorm:"column:FCSTUM3;" json:"fcstum3"  form:"fcstum3" `
	FCSTUM4    string    `gorm:"column:FCSTUM4;" json:"fcstum4"  form:"fcstum4" `
	FCSUBSTITU string    `gorm:"column:FCSUBSTITU;" json:"fcsubstitu"  form:"fcsubstitu" `
	FCSUPP     string    `gorm:"column:FCSUPP;" json:"fcsupp"  form:"fcsupp" `
	FCSUPPBOI  string    `gorm:"column:FCSUPPBOI;" json:"fcsuppboi"  form:"fcsuppboi" `
	FCSWHOUSE  string    `gorm:"column:FCSWHOUSE;" json:"fcswhouse"  form:"fcswhouse" `
	FCTAXTYPE  string    `gorm:"column:FCTAXTYPE;" json:"fctaxtype"  form:"fctaxtype" `
	FCTYPE     string    `gorm:"column:FCTYPE;" json:"fctype"  form:"fctype" `
	FCU1ACC    string    `gorm:"column:FCU1ACC;" json:"fcu1acc"  form:"fcu1acc" `
	FCU1STATUS string    `gorm:"column:FCU1STATUS;" json:"fcu1status"  form:"fcu1status" `
	FCU2STATUS string    `gorm:"column:FCU2STATUS;" json:"fcu2status"  form:"fcu2status" `
	FCU3STATUS string    `gorm:"column:FCU3STATUS;" json:"fcu3status"  form:"fcu3status" `
	FCU4STATUS string    `gorm:"column:FCU4STATUS;" json:"fcu4status"  form:"fcu4status" `
	FCU5STATUS string    `gorm:"column:FCU5STATUS;" json:"fcu5status"  form:"fcu5status" `
	FCU6STATUS string    `gorm:"column:FCU6STATUS;" json:"fcu6status"  form:"fcu6status" `
	FCU7STATUS string    `gorm:"column:FCU7STATUS;" json:"fcu7status"  form:"fcu7status" `
	FCU8STATUS string    `gorm:"column:FCU8STATUS;" json:"fcu8status"  form:"fcu8status" `
	FCU9STATUS string    `gorm:"column:FCU9STATUS;" json:"fcu9status"  form:"fcu9status" `
	FCUDATE    string    `gorm:"column:FCUDATE;" json:"fcudate"  form:"fcudate" `
	FCUM       string    `gorm:"column:FCUM;" json:"fcum"  form:"fcum" `
	FCUM1      string    `gorm:"column:FCUM1;" json:"fcum1"  form:"fcum1" `
	FCUM2      string    `gorm:"column:FCUM2;" json:"fcum2"  form:"fcum2" `
	FCUM3      string    `gorm:"column:FCUM3;" json:"fcum3"  form:"fcum3" `
	FCUM4      string    `gorm:"column:FCUM4;" json:"fcum4"  form:"fcum4" `
	FCUSEPDSER string    `gorm:"column:FCUSEPDSER;" json:"fcusepdser"  form:"fcusepdser" `
	FCUTIME    string    `gorm:"column:FCUTIME;" json:"fcutime"  form:"fcutime" `
	FCVATISOUT string    `gorm:"column:FCVATISOUT;" json:"fcvatisout"  form:"fcvatisout" `
	FCVATTYPE  string    `gorm:"column:FCVATTYPE;" json:"fcvattype"  form:"fcvattype" `
	FCWHATPRIC string    `gorm:"column:FCWHATPRIC;" json:"fcwhatpric"  form:"fcwhatpric" `
	FDADD      string    `gorm:"column:FDADD;" json:"fdadd"  form:"fdadd" `
	FDINACTIVE string    `gorm:"column:FDINACTIVE;" json:"fdinactive"  form:"fdinactive" `
	FDLASTBUYD string    `gorm:"column:FDLASTBUYD;" json:"fdlastbuyd"  form:"fdlastbuyd" `
	FDLASTSELL string    `gorm:"column:FDLASTSELL;" json:"fdlastsell"  form:"fdlastsell" `
	FIMILLISEC int64     `gorm:"column:FIMILLISEC;" json:"fimillisec"  form:"fimillisec" `
	FMERRMSG   string    `gorm:"column:FMERRMSG;" json:"fmerrmsg"  form:"fmerrmsg" `
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG;" json:"fmextratag"  form:"fmextratag" `
	FMINACTIVE string    `gorm:"column:FMINACTIVE;" json:"fminactive"  form:"fminactive" `
	FMPICNAME  string    `gorm:"column:FMPICNAME;" json:"fmpicname"  form:"fmpicname" `
	FNAGELONG  float64   `gorm:"column:FNAGELONG;" json:"fnagelong"  form:"fnagelong" `
	FNAVGCOST  float64   `gorm:"column:FNAVGCOST;" json:"fnavgcost"  form:"fnavgcost" `
	FNBATCHSIZ float64   `gorm:"column:FNBATCHSIZ;" json:"fnbatchsiz"  form:"fnbatchsiz" `
	FNCOMMISSI float64   `gorm:"column:FNCOMMISSI;" json:"fncommissi"  form:"fncommissi" `
	FNDLCOST   float64   `gorm:"column:FNDLCOST;" json:"fndlcost"  form:"fndlcost" `
	FNDMCOST   float64   `gorm:"column:FNDMCOST;" json:"fndmcost"  form:"fndmcost" `
	FNDUTYPCN  float64   `gorm:"column:FNDUTYPCN;" json:"fndutypcn"  form:"fndutypcn" `
	FNEOQTY    float64   `gorm:"column:FNEOQTY;" json:"fneoqty"  form:"fneoqty" `
	FNEXPFAC   float64   `gorm:"column:FNEXPFAC;" json:"fnexpfac"  form:"fnexpfac" `
	FNFIXCOST  float64   `gorm:"column:FNFIXCOST;" json:"fnfixcost"  form:"fnfixcost" `
	FNGROSSWG  float64   `gorm:"column:FNGROSSWG;" json:"fngrosswg"  form:"fngrosswg" `
	FNHEIGHT   float64   `gorm:"column:FNHEIGHT;" json:"fnheight"  form:"fnheight" `
	FNISCONSUM float64   `gorm:"column:FNISCONSUM;" json:"fnisconsum"  form:"fnisconsum" `
	FNIUMQTY   float64   `gorm:"column:FNIUMQTY;" json:"fniumqty"  form:"fniumqty" `
	FNLASTBUYP float64   `gorm:"column:FNLASTBUYP;" json:"fnlastbuyp"  form:"fnlastbuyp" `
	FNLENGTH   float64   `gorm:"column:FNLENGTH;" json:"fnlength"  form:"fnlength" `
	FNLTDMFG   float64   `gorm:"column:FNLTDMFG;" json:"fnltdmfg"  form:"fnltdmfg" `
	FNLTDPO    float64   `gorm:"column:FNLTDPO;" json:"fnltdpo"  form:"fnltdpo" `
	FNLTDPPARE float64   `gorm:"column:FNLTDPPARE;" json:"fnltdppare"  form:"fnltdppare" `
	FNLTDQC    float64   `gorm:"column:FNLTDQC;" json:"fnltdqc"  form:"fnltdqc" `
	FNLTDSUBC  float64   `gorm:"column:FNLTDSUBC;" json:"fnltdsubc"  form:"fnltdsubc" `
	FNMAXPRICE float64   `gorm:"column:FNMAXPRICE;" json:"fnmaxprice"  form:"fnmaxprice" `
	FNMFUMQTY  float64   `gorm:"column:FNMFUMQTY;" json:"fnmfumqty"  form:"fnmfumqty" `
	FNMINPRICE float64   `gorm:"column:FNMINPRICE;" json:"fnminprice"  form:"fnminprice" `
	FNOHCOST   float64   `gorm:"column:FNOHCOST;" json:"fnohcost"  form:"fnohcost" `
	FNPRICE    float64   `gorm:"column:FNPRICE;" json:"fnprice"  form:"fnprice" `
	FNREORDERP float64   `gorm:"column:FNREORDERP;" json:"fnreorderp"  form:"fnreorderp" `
	FNSAFETYQT float64   `gorm:"column:FNSAFETYQT;" json:"fnsafetyqt"  form:"fnsafetyqt" `
	FNSALEPRIC float64   `gorm:"column:FNSALEPRIC;" json:"fnsalepric"  form:"fnsalepric" `
	FNSCOSTKE1 float64   `gorm:"column:FNSCOSTKE1;" json:"fnscostke1"  form:"fnscostke1" `
	FNSCOSTKE2 float64   `gorm:"column:FNSCOSTKE2;" json:"fnscostke2"  form:"fnscostke2" `
	FNSTDCOST  float64   `gorm:"column:FNSTDCOST;" json:"fnstdcost"  form:"fnstdcost" `
	FNSTIUMQTY float64   `gorm:"column:FNSTIUMQTY;" json:"fnstiumqty"  form:"fnstiumqty" `
	FNSTUMQTY1 float64   `gorm:"column:FNSTUMQTY1;" json:"fnstumqty1"  form:"fnstumqty1" `
	FNSTUMQTY2 float64   `gorm:"column:FNSTUMQTY2;" json:"fnstumqty2"  form:"fnstumqty2" `
	FNSTUMQTY3 float64   `gorm:"column:FNSTUMQTY3;" json:"fnstumqty3"  form:"fnstumqty3" `
	FNSTUMQTY4 float64   `gorm:"column:FNSTUMQTY4;" json:"fnstumqty4"  form:"fnstumqty4" `
	FNTRANSEXP float64   `gorm:"column:FNTRANSEXP;" json:"fntransexp"  form:"fntransexp" `
	FNU1CNT    float64   `gorm:"column:FNU1CNT;" json:"fnu1cnt"  form:"fnu1cnt" `
	FNU2CNT    float64   `gorm:"column:FNU2CNT;" json:"fnu2cnt"  form:"fnu2cnt" `
	FNU3CNT    float64   `gorm:"column:FNU3CNT;" json:"fnu3cnt"  form:"fnu3cnt" `
	FNU4CNT    float64   `gorm:"column:FNU4CNT;" json:"fnu4cnt"  form:"fnu4cnt" `
	FNU5CNT    float64   `gorm:"column:FNU5CNT;" json:"fnu5cnt"  form:"fnu5cnt" `
	FNU6CNT    float64   `gorm:"column:FNU6CNT;" json:"fnu6cnt"  form:"fnu6cnt" `
	FNU7CNT    float64   `gorm:"column:FNU7CNT;" json:"fnu7cnt"  form:"fnu7cnt" `
	FNU8CNT    float64   `gorm:"column:FNU8CNT;" json:"fnu8cnt"  form:"fnu8cnt" `
	FNU9CNT    float64   `gorm:"column:FNU9CNT;" json:"fnu9cnt"  form:"fnu9cnt" `
	FNUMQTY1   float64   `gorm:"column:FNUMQTY1;" json:"fnumqty1"  form:"fnumqty1" `
	FNUMQTY2   float64   `gorm:"column:FNUMQTY2;" json:"fnumqty2"  form:"fnumqty2" `
	FNUMQTY3   float64   `gorm:"column:FNUMQTY3;" json:"fnumqty3"  form:"fnumqty3" `
	FNUMQTY4   float64   `gorm:"column:FNUMQTY4;" json:"fnumqty4"  form:"fnumqty4" `
	FNUMRATIO  float64   `gorm:"column:FNUMRATIO;" json:"fnumratio"  form:"fnumratio" `
	FNUMRATIOA float64   `gorm:"column:FNUMRATIOA;" json:"fnumratioa"  form:"fnumratioa" `
	FNUMRATIOB float64   `gorm:"column:FNUMRATIOB;" json:"fnumratiob"  form:"fnumratiob" `
	FNVARCOST  float64   `gorm:"column:FNVARCOST;" json:"fnvarcost"  form:"fnvarcost" `
	FNVATRATE  float64   `gorm:"column:FNVATRATE;" json:"fnvatrate"  form:"fnvatrate" `
	FNVOLUME   float64   `gorm:"column:FNVOLUME;" json:"fnvolume"  form:"fnvolume" `
	FNWIDTH    float64   `gorm:"column:FNWIDTH;" json:"fnwidth"  form:"fnwidth" `
	FNWTAXRATE float64   `gorm:"column:FNWTAXRATE;" json:"fnwtaxrate"  form:"fnwtaxrate" `
	FTDATETIME time.Time `gorm:"column:FTDATETIME;" json:"ftdatetime"  form:"ftdatetime" default:"now"`
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT;" json:"ftlastedit"  form:"ftlastedit" `
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD;" json:"ftlastupd"  form:"ftlastupd" default:"now"`
	FTSRCUPD   time.Time `gorm:"column:FTSRCUPD;" json:"ftsrcupd"  form:"ftsrcupd" default:"now"`
}

func (Product) TableName() string {
	return "PROD"
}

func (obj *Product) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(6)
	obj.FCSKID = fmt.Sprintf("AB%s", id)
	return
}
