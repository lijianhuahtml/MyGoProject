package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

	// 服务器端口
	err := r.Run(":8088")
	if err != nil {
		return
	}

}
