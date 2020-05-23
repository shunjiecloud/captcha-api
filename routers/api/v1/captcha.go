package v1

import (
	"fmt"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/shunjiecloud/captcha_api/schemas"
)

//GetCaptcha 获取验证码
func GetCaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp schemas.GetCaptchaResponse
		id := captcha.New()
		resp.CaptchaID = id
		resp.URL = fmt.Sprintf("/captcha/v1/captcha/%v.png", id)
		c.JSON(http.StatusOK, &resp)
	}
}

//CaptchaSrv 验证码srv
func CaptchaSrv() gin.HandlerFunc {
	captchaSrv := captcha.Server(captcha.StdWidth, captcha.StdHeight)
	return func(c *gin.Context) {
		captchaSrv.ServeHTTP(c.Writer, c.Request)
	}
}
