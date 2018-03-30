package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ginblog/router/api"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode("debug")

	router.GET("api/message", api.Getpersons) //获取文章列表(所有文章)

	//router.GET("api/article/:id", api.GetSingleArticle) //获取指定的文章

	router.PUT("api/editarticle/:id", api.UpdateArticle) //编辑文章

	router.GET("api/delarticle/:id", api.DelArticle) //删除指定文章

	router.POST("api/addarticle", api.Addpersons) //添加

	//router.GET("test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
		//	"hello": "world",
	//	})
	//})

	return router
}
