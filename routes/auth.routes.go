package routes

import (
	"github.com/abhishekmaurya0/2fa/controllers"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.LoginUser)
	router.POST("/otp/enable", rc.authController.Enable2FA)
	//router.POST("/otp/verify", rc.authController.VerifyOTP)
	//router.POST("/otp/validate", rc.authController.ValidateOTP)
	router.POST("/otp/disable", rc.authController.DisableOTP)
}
