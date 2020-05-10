package v1

import (
	"fmt"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/shunjiecloud/captcha_api/schemas"
	//api "github.com/shunjiecloud/captcha_api/client"
	//"github.com/shunjiecloud/errors"
	//"github.com/shunjiecloud/pkg/app"
	//ec "github.com/shunjiecloud/pkg/errcode"
	//"go.uber.org/zap"
)

//GetCaptcha 获取验证码
func GetCaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp schemas.GetCaptchaResponse
		id := captcha.New()
		resp.CaptchaID = id
		resp.URL = fmt.Sprintf("/captcha/v1/api/captcha/%v.png", id)
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

// func PostCaptchaVerfify() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		gApp := app.Gin{C: c}
// 		var request api.PostCaptchaVerfifyRequest
// 		err := gApp.C.ShouldBindJSON(&request)
// 		if err != nil {
// 			gApp.WriteError(err)
// 			return
// 		}
// 		isOk := captcha.VerifyString(request.CaptchaId, request.Solution)
// 		if isOk == true {
// 			c.JSON(http.StatusOK, &app.Ok)
// 		} else {
// 			gApp.WriteError(errors.New(code.ErrCaptchaVerfifyFailed, zap.String("solution", request.Solution)))
// 		}
// 	}
// }
