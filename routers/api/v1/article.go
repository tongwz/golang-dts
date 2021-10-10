package v1

import (
	"golang-dts/models"
	"golang-dts/pkg/e"
	"golang-dts/pkg/logging"
	"golang-dts/pkg/setting"
	"golang-dts/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type ArticleInfo struct {
	TID       int    `json:"tag_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Desc      string `json:"desc"`
	CreatedBy string `json:"created_by" binding:"required"`
}

// 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("Id必须大于0")

	code := e.INVALID_PARAMS

	var data interface{}
	if !valid.HasErrors() {
		if models.ExitArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只能是0或者1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章
func AddArticle(c *gin.Context) {
	var info ArticleInfo

	state := 1

	maps := make(map[string]interface{})

	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  err.Error(),
			"data": make(map[string]string),
		})
		return
	}

	valid := validation.Validation{}

	valid.Required(info.Title, "title").Message("文章标题必填")
	valid.Required(info.Content, "content").Message("文章内容必填")
	valid.Required(info.CreatedBy, "created_by").Message("创建人必填")
	valid.Min(info.TID, 1, "tag_id").Message("文章标签不能小于0")

	// 默认参数错误
	code := e.INVALID_PARAMS

	maps["tag_id"] = info.TID
	maps["title"] = info.Title
	maps["content"] = info.Content
	maps["created_by"] = info.CreatedBy
	maps["desc"] = info.Desc
	maps["state"] = state

	if !valid.HasErrors() {
		code = e.SUCCESS
		models.AddArticle(maps)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": info,
	})

}

// 修改文章
func EditArticle(c *gin.Context) {
	maps := make(map[string]interface{})
	type EditArticleInfo struct {
		TagID      int    `json:"tag_id" binding:"required"`
		Title      string `json:"title"`
		Content    string `json:"content"`
		Desc       string `json:"desc"`
		ModifiedBy string `json:"modified_by" binding:"required"`
	}
	valid := validation.Validation{}
	code := e.INVALID_PARAMS

	var info EditArticleInfo

	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
	}
	id := com.StrTo(c.Param("id")).MustInt()

	valid.Min(id, 1, "id").Message("ID大于0")
	valid.MaxSize(info.Title, 100, "title").Message("标题最长100字符")
	valid.MaxSize(info.Desc, 255, "desc").Message("简述最长255字符")
	valid.MaxSize(info.Content, 65536, "content").Message("内容最长为65535字符")
	valid.Required(info.ModifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(info.ModifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	if !valid.HasErrors() {
		if models.ExitArticleByID(id) {
			if models.ExistTagByID(info.TagID) {
				maps["tag_id"] = info.TagID
				maps["title"] = info.Title
				maps["content"] = info.Content
				maps["desc"] = info.Desc
				maps["modified_by"] = info.ModifiedBy
				models.EditArticle(id, maps)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			// 文章id不存在的情况
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		if models.ExitArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
