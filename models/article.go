package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 获取单个文章
func GetArticle(id int) (article Article) {
	db.Where("id = ? ", id).First(&article)
	db.Model(&article).Related(&article.Tag, "tag_id")
	return
}

// 查看文章是否存在
func ExitArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ? ", id).First(&article)
	return article.ID > 0
}

// 获取文章列表
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

// 获取文章总条数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

// 添加文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

// 编辑文章
func EditArticle(id int, data map[string]interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)
	return true
}

// 删除文章
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}

// 更新文章前触发的事件
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("modifiedOn", time.Now().Unix())
	return nil
}

// 插入文章前触发的事件
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}
