package routers

import (
	"golang-dts/middleware/jwt"
	"golang-dts/pkg/setting"

	"golang-dts/routers/api"
	v1 "golang-dts/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	// token获取
	r.GET("/auth", api.GetAuth)

	// 设置路由组
	apiV1 := r.Group("/api/v1").Use(jwt.JWT())
	{
		// 获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		// 新建标签
		apiV1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiV1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiV1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiV1.POST("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "童伟珍到此一游20210706",
		})
	})

	return r
}
