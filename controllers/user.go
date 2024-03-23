package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/trishan9/Gin-CRUD/helpers"
	"github.com/trishan9/Gin-CRUD/initializers"
	"github.com/trishan9/Gin-CRUD/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Get the email and password from Request Body
	var reqBody struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.Bind(&reqBody)

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)
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

	// Upload the image on the cloud, get the image url and remove from server
	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	var ctx = context.Background()
	resp, err := cld.Upload.Upload(ctx, "assets/uploads/"+file.Filename, uploader.UploadParams{PublicID: "my_image"})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload avatar"})
		return
	}

	log.Println(resp.SecureURL)

	defer func() {
		os.Remove("assets/uploads/" + file.Filename)
	}()

	// Create the OTP and send to mail
	otpCode := uuid.New().String()
	helpers.SendMail(otpCode[:8], reqBody.Name, reqBody.Email)

	user := models.User{Name: reqBody.Name, Email: reqBody.Email, Avatar: string(resp.SecureURL), Password: string(hashedPassword), OtpCode: string(otpCode[:8])}
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

	otpCode := uuid.New().String()
	helpers.SendMail(otpCode[:8], user.Name, user.Email)

	initializers.DB.Model(&user).Update("otp_code", otpCode[:8])

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP sent successfully!",
	})
}

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

	// Compare the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password is incorrect!",
		})
		return
	}

	// Generate and Set JWT Token on Cookie
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully!",
	})
}
