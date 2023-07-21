package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/abhishekmaurya0/2fa/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *models.RegisterUserInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	pass, err := GeneratePassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Wrong Password"})
		return
	}
	payload.Password = string(pass)

	newUser := models.User{
		Name:        payload.Name,
		Email:       strings.ToLower(payload.Email),
		Password:    payload.Password,
		Otp_enabled: false,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Email already exist, please use another email address"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Registered successfully, please login"})
}

func (ac *AuthController) LoginUser(ctx *gin.Context) {
	var payload *models.LoginUserInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}
	if !ValidatePass(payload.Password, []byte(user.Password)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Password"})
		return
	}
	userResponse := gin.H{
		"id":          user.ID.String(),
		"name":        user.Name,
		"email":       user.Email,
		"otp_enabled": user.Otp_enabled,
	}
	if userResponse["otp_enabled"] != false {
		otp := generateTOTP(payload.Secret, time.Now())
		flag := validateTOTP(otp, user.Otp_secret, time.Now())
		if !flag {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "OTP verification is unsuccessful"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": userResponse})
}
func (ac *AuthController) Enable2FA(ctx *gin.Context) {
	var payload *models.LoginUserInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	if !ValidatePass(payload.Password, []byte(user.Password)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Password"})
		return
	}
	user.Otp_enabled = true
	user.Otp_secret = generateSecretKey()
	result = ac.DB.Model(user).Updates(user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	userResponse := gin.H{
		"id":          user.ID.String(),
		"name":        user.Name,
		"email":       user.Email,
		"otp_enabled": user.Otp_enabled,
		"otp_secret":  user.Otp_secret,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": userResponse})

}

func (ac *AuthController) DisableOTP(ctx *gin.Context) {
	var payload *models.OTPInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "id = ?", payload.UserId)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "User doesn't exist"})
		return
	}

	user.Otp_enabled = false
	ac.DB.Save(&user)

	userResponse := gin.H{
		"id":          user.ID.String(),
		"name":        user.Name,
		"email":       user.Email,
		"otp_enabled": user.Otp_enabled,
	}
	ctx.JSON(http.StatusOK, gin.H{"otp_disabled": true, "user": userResponse})
}
