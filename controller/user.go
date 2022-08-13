package controller

import (
	"blog/dao"
	models "blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"net/http"
	"strconv"
)

func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)
}

func ListUser(c *gin.Context) {
	users := dao.Mgr.GetAllUsers()
	c.HTML(http.StatusOK, "user.html", users)
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)

	c.Redirect(http.StatusMovedPermanently, "/")
}

func GoRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.HTML(http.StatusOK, "login.html", "用户名不存在！")
		fmt.Println("用户名不存在！")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(http.StatusOK, "login.html", "密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	}
}

func GoLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(http.StatusOK, "postIndex.html", posts)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := models.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(http.StatusMovedPermanently, "/post_index")

}

func GoAddPost(c *gin.Context) {
	c.HTML(http.StatusOK, "post.html", nil)
}

func PostDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)

	content := blackfriday.Run([]byte(p.Content))

	c.HTML(http.StatusOK, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}
