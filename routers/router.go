package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorestfulapiforcms/middleware/jwt"
	"gorestfulapiforcms/middleware/log"
	"gorestfulapiforcms/routers/api"

	"gorestfulapiforcms/pkg/setting"

	"gorestfulapiforcms/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.Config().Run_mode)

	r.GET("/auth", api.GetAuth)

	api_v1 := r.Group("/api/v1")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api_v1.Use(log.LoggerToFile(), jwt.JWT())
	{
		//获取标签列表
		api_v1.GET("/tags", v1.GetTags)
		//新建标签
		api_v1.POST("/tags", v1.AddTag)
		//更新指定标签
		api_v1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		api_v1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		api_v1.GET("/articles", v1.GetArticles)
		//获取指定文章
		api_v1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		api_v1.POST("/articles", v1.AddArticle)
		//更新指定文章
		api_v1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		api_v1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
