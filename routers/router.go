package routers

import (
	"github.com/gin-gonic/gin"
	"gorestfulapiforcms/middleware/log"
	"gorestfulapiforcms/pkg/setting"
	"gorestfulapiforcms/routers/j"
)

func
InitRouter() *gin.Engine {
	r := gin.New()//初始化gin
/*	gin.ForceConsoleColor()//控制台的日志颜色控制
	r.Use(gin.Logger())//将日志输出到控制台*/
	r.Use(gin.Recovery())
	gin.SetMode(setting.Config().Run_mode)
	//r.GET("/auth", j.GetAuth)

	//j文件夹的api是给内部提供的json
	group_j := r.Group("/j")
	//group_j.Use(log.LoggerToFile(), jwt.JWT())
	group_j.Use(log.LoggerToFile())
	{
		//获取标签列表
		group_j.GET("/tags", j.GetTags)
		//新建标签
		group_j.POST("/tags", j.AddTag)
		//更新指定标签
		group_j.PUT("/tags/:id", j.EditTag)
		//删除指定标签
		group_j.DELETE("/tags/:id", j.DeleteTag)

		//获取文章列表
		group_j.GET("/articles", j.GetArticles)
		//获取指定文章
		group_j.GET("/articles/:id", j.GetArticle)
		//新建文章
		group_j.POST("/articles", j.AddArticle)
		//更新指定文章
		group_j.PUT("/articles/:id", j.EditArticle)
		//删除指定文章
		group_j.DELETE("/articles/:id", j.DeleteArticle)
	}

	//j文件夹下是给外部提供的
	group_x := r.Group("/x")
	//group_j.Use(log.LoggerToFile(), jwt.JWT())
	group_x.Use()
	{
		//获取标签列表
	}

	return r
}
