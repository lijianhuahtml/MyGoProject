package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 自定义中间件 连接器
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("userSession", "userid-1")

		context.Next()  // 放行
		context.Abort() // 阻止
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	// 创建一个服务
	r := gin.Default()

	// 设置icon图标
	//r.Use(favicon.New("gin/Outline_fuben3.svg"))

	// 加载静态页面
	r.LoadHTMLGlob("gin/templates/*")
	// 加载资源文件
	r.Static("gin/static", "./gin/static")

	// 访问地址，处理我们的请求， request response
	// 获取 /user/:id 这类型路由绑定的参数
	//r.GET("/user/:id", func(c *gin.Context) {
	//
	//	// 获取url参数id
	//	id := c.Param("id")
	//	c.String(http.StatusOK, "%s的信息", id)
	//
	//})

	r.GET("/index", func(c *gin.Context) {
		//c.JSON()  json 数据
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "tzzsb",
		})
	})

	//Gin RustFul
	//r.PUT()
	//r.DELETE()
	//r.POST()

	//传参
	//user?userid=xxx&username=lijianhua
	r.GET("/user", myHandler(), func(context *gin.Context) {

		// 取出中间件中的值
		usersession := context.MustGet("userSession")

		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":      userid,
			"username":    username,
			"usersession": usersession,
		})
	})

	//user/info/1/lijianhua
	r.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 前端给后端传递 json
	r.POST("/json", func(context *gin.Context) {
		// request.body
		b, _ := context.GetRawData()

		var m map[string]interface{}
		// 包装为json数据
		_ = json.Unmarshal(b, &m)
		context.JSON(http.StatusOK, m)
	})

	// 处理表单
	// 支持函数式编程 =>
	r.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	// 路由
	r.GET("/test", func(context *gin.Context) {
		// 重定向 301
		context.Redirect(301, "https://www.daidu.com")
	})

	// 404
	r.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	// 路由组 /user/add
	userGroup := r.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.GET("/login")
		userGroup.GET("/logout")
	}
	orderGroup := r.Group("/order")
	{
		orderGroup.GET("/add")
		orderGroup.POST("/delete")
	}

	// 服务器端口
	err := r.Run(":8088")
	if err != nil {
		return
	}

}
