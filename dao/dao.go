package dao

import (
	models "blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	AddUser(user *models.User)
	GetAllUsers() []models.User
	Login(username string) models.User

	AddPost(post *models.Post)
	GetAllPost() []models.Post
	GetPost(pid int) models.Post
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to init db:", err)
	}

	Mgr = &manager{db: db}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}

func (mgr *manager) AddUser(user *models.User) {
	mgr.db.Create(user)
}
func (mgr *manager) GetAllUsers() []models.User {
	var users = make([]models.User, 10)
	mgr.db.Find(&users)
	return users
}
func (mgr *manager) Login(username string) models.User {
	var user models.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

func (mgr *manager) AddPost(post *models.Post) {
	mgr.db.Create(&post)
}

func (mgr *manager) GetAllPost() []models.Post {
	var posts = make([]models.Post, 10)
	mgr.db.Find(&posts)
	return posts
}

func (mgr *manager) GetPost(pid int) models.Post {
	var post models.Post
	mgr.db.First(&post, pid)
	return post
}
