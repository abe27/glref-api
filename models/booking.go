package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Booking struct {
	FCSKID    string   `gorm:"primaryKey;column:FCSKID;type:char;size:8;not null;index;" json:"fcskid"`
	FCCODE    string   `gorm:"column:FCCODE" json:"fccode,omitempty"`
	FCNAME    string   `json:"fcname,omitempty"`
	FCREFTYPE string   `gorm:"column:FCREFTYPE" json:"fcreftype,omitempty"`
	FCPREFIX  string   `json:"fcprefix,omitempty"`
	REFTYPE   *Reftype `gorm:"foreignKey:FCREFTYPE;references:FCSKID;" json:"ref_type"`
}

func (Booking) TableName() string {
	return "BOOK"
}

func (obj *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
