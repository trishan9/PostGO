package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/trishan9/Gin-CRUD/initializers"
	"github.com/trishan9/Gin-CRUD/models"
)

func CreatePosts(c *gin.Context) {
	// Get User
	user, _ := c.Get("user")

	// Get Post Data from Request Body
	var reqBody struct {
		Title string `json:"title"`
		Body  string `json:"body,omitempty"`
	}
	c.Bind(&reqBody)

	// Create Post in DB
	post := models.Post{Title: reqBody.Title, Body: reqBody.Body, UserID: user.(models.User).ID}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Can't create post",
		})
		return
	}

	// Return Post
	c.JSON(201, gin.H{
		"post":    post,
		"message": "Post created",
	})
}

func GetPosts(c *gin.Context) {
	user, _ := c.Get("user")

	// Get All Posts
	var posts []models.Post
	result := initializers.DB.Find(&posts, "user_id = ?", user.(models.User).ID)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Can't get posts",
		})
		return
	}

	c.JSON(200, gin.H{
		"posts":   posts,
		"message": "Posts fetched successfully",
	})
}

func GetPostById(c *gin.Context) {
	// Get Post ID
	postId := c.Param("id")
	// c.GetQuery("type")

	// Get Post
	var post models.Post
	result := initializers.DB.First(&post, postId)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Can't get post",
		})
		return
	}

	c.JSON(200, gin.H{
		"post":    post,
		"message": "Post fetched successfully",
	})
}

func UpdatePost(c *gin.Context) {
	postId := c.Param("id")

	var reqBody struct {
		Title string `json:"title"`
		Body  string `json:"body,omitempty"`
	}
	c.Bind(&reqBody)

	var post models.Post
	result := initializers.DB.First(&post, postId)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Can't update the post",
		})
		return
	}

	updateResult := initializers.DB.Model(&post).Updates(models.Post{Title: reqBody.Title, Body: reqBody.Body})

	if updateResult.Error != nil {
		c.JSON(400, gin.H{
			"message": "Can't update the post",
		})
		return
	}

	c.JSON(200, gin.H{
		"post":    post,
		"message": "Post updated successfully",
	})
}

func DeletePost(c *gin.Context) {
	postId := c.Param("id")

	result := initializers.DB.Delete(&models.Post{}, postId)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Can't delete the post",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}
