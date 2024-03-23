package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trishan9/Gin-CRUD/helpers"
	"github.com/trishan9/Gin-CRUD/initializers"
	"github.com/trishan9/Gin-CRUD/models"
)

type SignUpRequestBody struct {
	Name     string `json:"name" example:"John Doe"`
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"password123"`
	Avatar   string `json:"avatar" example:"avatar.jpg" type:"file"`
}

type SignUpSuccessResponse struct {
	Message string `json:"message" example:"Sign up successful"`
}

// SignUp registers a new user.
// @Tags Auth
// @Summary Register a new user
// @Description Register a new user
// @Accept json, multipart/form-data
// @Produce json
// @Param body body SignUpRequestBody true "User Data"
// @Success 200 {object} SignUpSuccessResponse
// @Router /api/auth/signup [post]
func SignUp(c *gin.Context) {
	var reqBody struct {
		Name     string `json:"name" form:"name" binding:"required"`
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	c.Bind(&reqBody)

	hashedPassword, err := helpers.GenerateHash(reqBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Get the image and upload it on server
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar"})
		return
	}

	// Get the image url after uploading to cloud
	avatarUrl, err := helpers.UploadToCloudinary(file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload avatar"})
		return
	}

	otpCode := helpers.GenerateOtp()
	mailErr := helpers.SendMail(otpCode, reqBody.Name, reqBody.Email)

	if mailErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": mailErr.Error()})
		return
	}

	user := models.User{Name: reqBody.Name, Email: reqBody.Email, Avatar: string(avatarUrl), Password: string(hashedPassword), OtpCode: string(otpCode)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create user",
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sign up successful",
	})
}

type ValidateOtpRequestBody struct {
	Email   string `json:"email" example:"john@example.com"`
	OtpCode string `json:"otp" example:"123456"`
}

type ValidateOtpSuccessResponse struct {
	Message string `json:"message" example:"Email address verified successfully!"`
}

// / ValidateOtp validates the OTP for email verification.
// @Tags Auth
// @Summary Validate email address
// @Description Validate email address using OTP
// @Accept json
// @Produce json
// @Param body body ValidateOtpRequestBody true "Request Body"
// @Success 200 {object} ValidateOtpSuccessResponse
// @Router /api/auth/signup/validate [post]
func ValidateOtp(c *gin.Context) {
	var reqBody struct {
		Email   string `json:"email"`
		OtpCode string `json:"otp"`
	}
	c.Bind(&reqBody)

	var user models.User
	initializers.DB.First(&user, "email = ?", reqBody.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No user found with this email address!",
		})
		return
	}

	if user.IsVerified {
		c.JSON(http.StatusOK, gin.H{
			"message": "User has already been verified!",
		})
		return
	}

	if user.OtpCode != reqBody.OtpCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "OTP is incorrect!",
		})
		return
	}

	initializers.DB.Model(&user).Update("is_verified", true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Email address verified successfully!",
	})
}

type RegenerateOtpRequestBody struct {
	Email string `json:"email" example:"john@example.com"`
}

type RegenerateOtpSuccessResponse struct {
	Message string `json:"message" example:"OTP sent successfully!"`
}

// RegenerateOtp regenerates the OTP for email verification.
// @Tags Auth
// @Summary Regenerate OTP
// @Description Regenerate OTP for email verification
// @Accept json
// @Produce json
// @Param body body RegenerateOtpRequestBody true "Request Body"
// @Success 200 {object} RegenerateOtpSuccessResponse
// @Router /api/auth/signup/regenerate [post]
func RegenerateOtp(c *gin.Context) {
	var reqBody struct {
		Email string `json:"email"`
	}
	c.Bind(&reqBody)

	var user models.User
	initializers.DB.First(&user, "email = ?", reqBody.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No user found with this email address!",
		})
		return
	}

	otpCode := helpers.GenerateOtp()
	mailErr := helpers.SendMail(otpCode, user.Name, user.Email)

	if mailErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": mailErr.Error()})
		return
	}

	initializers.DB.Model(&user).Update("otp_code", otpCode)

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP sent successfully!",
	})
}

type LoginRequestBody struct {
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"password123"`
}

type LoginSuccessResponse struct {
	Message string `json:"message" example:"Logged in successfully!"`
}

// Login logs in a user.
// @Tags Auth
// @Summary User login
// @Description Log in a user
// @Accept json
// @Produce json
// @Param body body LoginRequestBody true "Request Body"
// @Success 200 {object} LoginSuccessResponse
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var reqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.Bind(&reqBody)

	var user models.User
	initializers.DB.First(&user, "email = ?", reqBody.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No user found with this email address!",
		})
		return
	}

	err := helpers.CompareHash(reqBody.Password, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password is incorrect!",
		})
		return
	}

	tokenString, _ := helpers.GenerateToken(user.ID)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully!",
	})
}
