package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	// 创建带默认中间件（日志与恢复）的 Gin 路由器
	engine := gin.Default()
	// 图标
	engine.Use(favicon.New("snow.ico"))
	// 加载静态页面
	engine.LoadHTMLGlob("template/*")
	// 加载资源文件
	engine.Static("/static", "./static")

	// 定义路由，处理请求
	engine.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"message": "你好啊"})
	})
	engine.GET("/hello", func(context *gin.Context) {
		// 返回 JSON 响应
		context.JSON(http.StatusOK, gin.H{"message": "hello, get"})
	})
	engine.POST("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "helo, post"})
	})

	// 重定向
	engine.GET("/re", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	// 404页面
	engine.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	// 接收前端传来的参数
	// use query: /user/info?userid=1&username=jerry
	engine.GET("/user/info", func(context *gin.Context) {
		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{"userid": userid, "username": username})
	})
	// use param: /user/info/2/tom
	engine.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{"userid": userid, "username": username})
	})

	// form表单提交
	// 使用中间件
	engine.POST("/user/add", myHandler(), func(context *gin.Context) {
		str := context.MustGet("usersession").(string)
		log.Println("==========>", str)

		username := context.PostForm("username")
		password := context.PostForm("password")

		context.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"password": password,
		})
	})

	// 监听端口 8080，启动服务器
	// http://localhost:8080
	_ = engine.Run(":8080")
}

// 自定义中间件
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 通过自定义的中间件，设置的值，在后续的处理只要调用了这个中间件的都可以拿到这里的参数
		context.Set("usersession", "valueid-01")
		context.Next() // 放行
		//context.Abort()  // 阻止
	}
}
