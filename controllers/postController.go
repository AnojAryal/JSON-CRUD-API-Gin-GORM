package controllers

import (
	"github.com/anojaryal/json-crud-api-gin-gorm/initializers"
	"github.com/anojaryal/json-crud-api-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

//create a post
func PostCreate(c*gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body:body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(201, gin.H{
		"post": post,
	})

}

//fetch all posts
func PostsIndex(c*gin.Context) {
	//get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

//get post by id
func PostsShow(c*gin.Context) {
	//get id of url
	id := c.Param("id")

	//get the post
	var post models.Post
	initializers.DB.First(&post, id)

	//respond with them
	c.JSON(200, gin.H{
		"post": post,
	})

}

//update post
func PostsUpdate(c*gin.Context) {
	//get id of url
	id := c.Param("id")

	//get data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title : body.Title,
		Body  : body.Body,
	})

	//respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

//delete post
func PostsDelete(c*gin.Context) {
	//get id of url
	id := c.Param("id")

	//delete it
	initializers.DB.Delete(&models.Post{}, id)
		
	//respond
	c.Status(200)
}