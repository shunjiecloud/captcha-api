package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/shunjiecloud-proto/captcha/proto"
	"github.com/shunjiecloud/captcha-api/modules"
	"github.com/shunjiecloud/captcha-api/schemas"
	"github.com/shunjiecloud/pkg/app"
)

//
// @Summary 获取验证码地址
// @tags 验证码
// @Description 获取验证码地址。
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /captcha/v1/captcha [get]
func GetCaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		appCtx := app.AppContext{
			GinCtx: c,
		}
		var resp schemas.GetCaptchaResponse
		//  captcha-srv CaptchaId
		getCaptchaIdResp, err := modules.ModuleContext.CaptchaSrvClient.CaptchaId(context.Background(), &proto.CaptchaIdRequest{})
		if err != nil {
			appCtx.WriteError(err)
			return
		}
		resp.CaptchaID = getCaptchaIdResp.CaptchaId
		resp.URL = fmt.Sprintf("/captcha/v1/captcha/%v.png", resp.CaptchaID)
		c.JSON(http.StatusOK, &resp)
	}
}

//
// @Summary 验证码图片
// @tags 验证码
// @Description 根据验证码地址，获取验证码图片。
// @Accept  json
// @Produce  image/png
// @Success 200 "验证码图片"
// @failure 400 {string} string	"404 page not found"
// @Param filename path string true "验证码图片名，例如：o5CaUbWHwjRUg6tyYrBW.png"
// @Router /captcha/v1/captcha/{filename} [get]
func CaptchaSrv() gin.HandlerFunc {
	captchaSrv := captcha.Server(captcha.StdWidth, captcha.StdHeight)
	return func(c *gin.Context) {
		captchaSrv.ServeHTTP(c.Writer, c.Request)
	}
}
