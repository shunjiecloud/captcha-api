package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shunjiecloud/captcha_api/routers/api"
	"github.com/shunjiecloud/pkg/middlewares"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("captcha/v1")

	//apiv1.GET("captcha", api.GetCaptcha())
	apiv1.GET("captcha/:filename", api.CaptchaSrv())
	//apiv1.POST("captcha/verfify", api.PostCaptchaVerfify())
	return r
}
