package main

import (
	"fmt"
	"github.com/ewenliu/gin-example/7.Gorm/db"
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (b BaseModel) BeforeCreate(scope *gorm.Scope) {
	_ = scope.SetColumn("create_time", time.Now())
	_ = scope.SetColumn("update_time", time.Now())
}

func (b BaseModel) BeforeUpdate(scope *gorm.Scope) {
	fmt.Println(time.Now())
	_ = scope.SetColumn("update_time", time.Now())
}

type Tag struct {
	BaseModel
	Id          int    `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Group       int    `json:"group" gorm:"SMALLINT"`
	CategoryId  int    `json:"category_id"`
}

func (Tag) TableName() string {
	return "tag"
}

type Result struct{
	CategoryId int
	Total int
}

func main() {
	db.SetupDB()

	var t = new(Tag)
	session := db.GetDB()
	defer func() {
		_ = session.Close()
	}()

	// 查询
	// 单条查询
	//session.First(&t)
	// 多条查询
	var tags []Tag
	session.Find(&tags)
	// 条件查询
	//var count int
	//Debug() 打印原始sql
	//session.Debug().Where("description = ? AND id = ?", "描述", "0").Find(&t)
	//session.Where("name like ?", "%lyw%").Order("id desc").Find(&tags)
	// count + order by
	//session.Where("name like ?", "%lyw%").Order("id desc").Find(&tags).Count(&count)
	// group by
	//var results []Result
	//session.Model(Tag{}).Select("category_id, count(*) as total").Group("category_id").Scan(&results)
	//fmt.Println(results)

	// 裸写sql
	//sql := "select category_id, count(*) as total from tag group by category_id"
	//session.Raw(sql).Scan(&results)

	// 链式过滤
	session = session.Where("name like ?", "%lyw%")
	session.Debug().Where("description = ?", "描述").Find(&t)


	// 更新
	//t.Description = fmt.Sprintf("gorm test2")
	//session.Save(&t)

	// 插入
	//t1 := Tag{Name: "lywtest1", Description: "描述", Group: 1, CategoryId: 6}
	//session.Create(&t1)

	// 删除
	//fmt.Println(t.Id)
}
