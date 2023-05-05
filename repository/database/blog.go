package database

import (
	"project_structure/config"
	"project_structure/model"
)

func CreateBlog(blog *model.Blog) error {
	if err := config.DB.Create(blog).Error; err != nil {
		return err
	}
	return nil
}

func GetBlogs() (blogs []model.Blog, err error) {
	if err = config.DB.Find(&blogs).Error; err != nil {
		return
	}
	return
}

func GetBlog(id uint) (blog model.Blog, err error) {
	blog.ID = id
	if err = config.DB.First(&blog).Error; err != nil {
		return
	}
	return
}

func GetBlogsByUserId(userID uint) (blog model.Blog, err error) {
	blog.UserID = userID
	if err = config.DB.Find(&blog).Error; err != nil {
		return
	}
	return
}

func UpdateBlog(blog *model.Blog) error {
	if err := config.DB.Updates(blog).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBlog(blog *model.Blog) error {
	if err := config.DB.Delete(blog).Error; err != nil {
		return err
	}
	return nil
}
