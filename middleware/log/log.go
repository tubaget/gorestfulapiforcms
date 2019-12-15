package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorestfulapiforcms/pkg/setting"
	"log"
	"os"
	"strconv"
	"time"
)

var absolute_file_name = ""

func init() {
	//用数组串联整个目录结构，没有哪层就创建哪层  获取日志写入路径没有则创建
	now := time.Now()

	file_name := strconv.Itoa(now.Day()) + ".log"

	var paths = [3] string {setting.Config().Log.Log_path, strconv.Itoa(now.Year()), strconv.Itoa(int(now.Month()))}

	now_path := ""
	for _, sub_path := range paths {
		now_path += sub_path + "/"
		fi, _ := os.Stat(now_path)

		if fi == nil {
			err := os.Mkdir(now_path, os.ModePerm)
			fmt.Println(now_path)

			if err != nil{
				fmt.Println(err)
			}
		}
	}

	absolute_file_name = now_path + file_name
	fi, _ := os.Stat(absolute_file_name)
	if fi == nil {
		os.Create(absolute_file_name) // 文件不存在就创建
	}
}
// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	//写入文件
	src, err := os.OpenFile(absolute_file_name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	//实例化
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{}
	//设置输出
	logger.Out = src
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	//logger.SetFormatter(&amp;logrus.TextFormatter{
	//	TimestampFormat:"2006-01-02 15:04:05",
	//})
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()

		header := c.Request.Header

		logIn(header)
		// 日志格式 -- 基础内容
		logger.Infof("%3d | %13v | %15s | %s | %s",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)



/*		jsonRes, err := pkg.MapToJson(header)
		if err != nil {
			fmt.Printf("Convert json to map failed with error: %+v\n", err)
		}

		logIn(header)*/
	}
}

//插入文本内容
func logIn(param string) {
	//写入文件
	src, err := os.OpenFile(absolute_file_name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	lll := log.New(src, "", 0)

	lll.Println(param)
}