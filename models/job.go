package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Job struct {
	FCSKID string `gorm:"primaryKey;column:FCSKID;type:char;size:8;not null;index;" json:"fcskid"`
	FCCODE string `gorm:"index:idx_fccode,unique" json:"fccode"`
	FCNAME string `json:"fcname,omitempty"`
}

func (Job) TableName() string {
	return "JOB"
}

func (obj *Job) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
