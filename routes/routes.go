package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1 := r.Group("/api/v1")
	//注册
	v1.POST("/signup", controllers.SignUpHandler)
	//登入
	v1.POST("/login", controllers.LoginHandler)
	v1.GET("/community", controllers.CommunityHandler)
	v1.GET("/community/:id", controllers.CommunityDetailHandler)
	v1.GET("/post/:id", controllers.GetPostDetailHandler)
	v1.GET("/posts", controllers.GetPostListHandler)
	v1.GET("/posts2", controllers.GetPostListHandler2)
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.POST("/post", controllers.CreatePostHandler)

		v1.POST("/vote", controllers.PostVoteController)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
