package routers

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/users", controller.ListUser)
	e.POST("/users", controller.AddUser)

	e.GET("/", controller.Index)
	e.POST("/register", controller.Register)
	e.GET("/register", controller.GoRegister)

	e.POST("/login", controller.Login)
	e.GET("/login", controller.GoLogin)

	e.GET("/post_index", controller.GetPostIndex)
	e.POST("/post", controller.AddPost)
	e.GET("/post", controller.GoAddPost)
	e.GET("/post_detail", controller.PostDetail)

	e.Run()
}
