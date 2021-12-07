package main

import (
	"fmt"
	"github.com/ewenliu/gin-example/7.Gorm/db"
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (b BaseModel) BeforeCreate(scope *gorm.Scope) {
	_ = scope.SetColumn("create_time", time.Now())
	_ = scope.SetColumn("update_time", time.Now())
}

func (b BaseModel ) BeforeUpdate(scope *gorm.Scope) {
	_ = scope.SetColumn("update_time", time.Now())
}

type Tag struct {

	BaseModel
	Id int `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Description string `json:"description"`
	Group int `json:"group"`

	CategoryId int `json:"category_id"`
}

func (Tag) TableName() string {
	return "tag"
}

func main() {
	db.SetupDB()

	var t = new(Tag)
	session := db.GetDB()
	session.First(&t)
	t.Description = fmt.Sprintf("gorm test")
	session.Save(&t)
	//fmt.Println(t.Id)
}
