package util

import (
	"gorestfulapiforcms/pkg/setting"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.Config().Basic.Page_size
	}

	return result
}
