package routers

import (
	"fmt"
	"gin-blog/middleware/jwt"
	"gin-blog/routers/api"
	"gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//gin.SetMode(setting.AppSetting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	apiTest := r.Group("/api/test")
	apiTest.GET("/value", func(c *gin.Context) {
		name := c.Query("name")
		fmt.Println(c.Request.URL.Query())
		c.JSON(http.StatusOK, []gin.H{
			{
				"name":  name,
				"value": 200,
			},
		})
	})
	apiTest.GET("ip", func(c *gin.Context) {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			fmt.Println(err)
		}
		ip := []string{}
		for _, address := range addrs {
			// 检查ip地址判断是否回环地址
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = append(ip, ipnet.IP.String())
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"ip": ip,
		})
	})

	return r
}
