package usecase

import (
	"errors"
	"fmt"
	"project_structure/model"
	"project_structure/repository/database"
)

type BlogUsecase interface {
	CreateBlog(blog *model.Blog) error
	GetBlog(id uint) (blog model.Blog, err error)
	GetListBlogs() (blogs []model.Blog, err error)
	UpdateBlog(blog *model.Blog) (err error)
	DeleteBlog(id uint) (err error)
}

func CreateBlog(blog *model.Blog) error {

	// check name cannot be empty
	if blog.Title == "" {
		return errors.New("blog title cannot be empty")
	}

	// check email
	if blog.Content == "" {
		return errors.New("blog content cannot be empty")
	}

	err := database.CreateBlog(blog)
	if err != nil {
		return err
	}

	return nil
}

func GetBlog(id uint) (blog model.Blog, err error) {
	blog, err = database.GetBlog(id)
	if err != nil {
		fmt.Println("GetBlog: Error getting blog from database")
		return
	}
	return
}

func GetListBlogs() (blogs []model.Blog, err error) {
	blogs, err = database.GetBlogs()
	if err != nil {
		fmt.Println("GetListBlogs: Error getting blogs from database")
		return
	}
	return
}

func UpdateBlog(blog *model.Blog) (err error) {
	err = database.UpdateBlog(blog)
	if err != nil {
		fmt.Println("UpdateBlog : Error updating blog, err: ", err)
		return
	}

	return
}

func DeleteBlog(id uint) (err error) {
	blog := model.Blog{}
	blog.ID = id
	err = database.DeleteBlog(&blog)
	if err != nil {
		fmt.Println("DeleteBlog : error deleting blog, err: ", err)
		return
	}

	return
}
