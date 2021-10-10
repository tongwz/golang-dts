package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 获取标签表数据
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// 获取数据总条数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 查询是否有这个标签
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	return tag.ID > 0

}

// 查询是否有这个标签通过id
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ? ", id).First(&tag)
	return tag.ID > 0
}

// 添加标签
func AddTag(name string, state int, CreatedBy string) bool {
	db.Create(&Tag{
		Name:      name,
		CreatedBy: CreatedBy,
		State:     state,
	})
	return true
}

// 修改标签
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ? ", id).Update(data)
	return true
}

// 标签更新数据前需要统一更新的字段
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

// 插入标签数据的时候统一加入的字段
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func DeleteTag(id int) bool {
	db.Where("id = ? ", id).Delete(&Tag{})
	return true
}
