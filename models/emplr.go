package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Employee struct {
	FCSKID  string `gorm:"primaryKey;column:FCSKID;type:char;size:8;not null;index;" json:"fcskid"`
	FCLOGIN string `json:"fcslogin,omitempty"`
	FCPW    string `json:"fcspassword,omitempty"`
}

func (Employee) TableName() string {
	return "EMPLR"
}

func (obj *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}

type FrmLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthSession struct {
	Header   string      `json:"header"`
	JwtType  string      `json:"jwt_type,omitempty"`
	JwtToken string      `json:"jwt_token,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
