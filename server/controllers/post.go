package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/trishan9/Gin-CRUD/initializers"
	"github.com/trishan9/Gin-CRUD/models"
)

type CreatePostRequestBody struct {
	Title string `json:"title" example:"Post Title"`
	Body  string `json:"body,omitempty" example:"Post Body"`
}

type CreatePostResponse struct {
	Post    models.Post `json:"post"`
	Message string      `json:"message" example:"Post created"`
}

// CreatePost creates a new post.
// @Summary Create a new post
// @Description Create a new post
// @Tags Posts
// @Accept json
// @Produce json
// @Param body body CreatePostRequestBody true "Post Data"
// @Security ApiKeyAuth
// @Success 201 {object} CreatePostResponse
// @Router /api/posts/ [post]
func CreatePost(c *gin.Context) {
	user, _ := c.Get("user")

	var reqBody struct {
		Title string `json:"title"`
		Body  string `json:"body,omitempty"`
	}
	c.Bind(&reqBody)

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

type GetPostsResponse struct {
	Posts   []models.Post `json:"posts"`
	Message string        `json:"message" example:"Posts fetched successfully"`
}

// GetPosts gets all posts of the authenticated user.
// @Summary Get all posts
// @Description Get all posts of the authenticated user
// @Tags Posts
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} GetPostsResponse
// @Router /api/posts/ [get]
func GetPosts(c *gin.Context) {
	user, _ := c.Get("user")

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

type GetPostResponse struct {
	Post    models.Post `json:"post"`
	Message string      `json:"message" example:"Post fetched successfully"`
}

// GetPostById gets a post by its ID.
// @Summary Get post by ID
// @Description Get a post by its ID
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security ApiKeyAuth
// @Success 200 {object} GetPostResponse
// @Router /api/posts/{id} [get]
func GetPostById(c *gin.Context) {
	postId := c.Param("id")
	// c.GetQuery("type")

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

type UpdatePostRequestBody struct {
	Title string `json:"title" example:"Updated Post Title"`
	Body  string `json:"body,omitempty" example:"Updated Post Body"`
}

type UpdatePostResponse struct {
	Post    models.Post `json:"post"`
	Message string      `json:"message" example:"Post updated successfully"`
}

// UpdatePost updates a post by its ID.
// @Summary Update post by ID
// @Description Update a post by its ID
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param body body UpdatePostRequestBody true "Updated Post Data"
// @Security ApiKeyAuth
// @Success 200 {object} UpdatePostResponse
// @Router /api/posts/{id} [patch]
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

type DeleteSuccessResponse struct {
	Message string `json:"message" example:"Post deleted successfully"`
}

// DeletePost deletes a post by its ID.
// @Summary Delete post by ID
// @Description Delete a post by its ID
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security ApiKeyAuth
// @Success 200 {object} DeleteSuccessResponse
// @Router /api/posts/{id} [delete]
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
