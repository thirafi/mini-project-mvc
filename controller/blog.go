package controller

import (
	"net/http"
	"project_structure/model"
	"project_structure/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBlogsController(c echo.Context) error {
	blogs, e := usecase.GetListBlogs()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"blogs":  blogs,
	})
}

func GetBlogController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	blog, err := usecase.GetBlog(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get blog",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"blogs":  blog,
	})
}

// create new blog
func CreateBlogController(c echo.Context) error {
	blog := model.Blog{}
	c.Bind(&blog)

	if err := usecase.CreateBlog(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create blog",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteBlog(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete blog",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf blog tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	blog := model.Blog{}
	c.Bind(&blog)
	blog.ID = uint(id)
	if err := usecase.UpdateBlog(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update blog",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf blog tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
	})
}
