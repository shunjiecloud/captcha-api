package schemas

//GetCaptchaResponse 获取验证码返回结构
type GetCaptchaResponse struct {
	CaptchaID string `json:"captcha_id"`
	URL       string `json:"captcha_url"`
}

//PostCaptchaVerfifyRequest 校验验证码请求结构
type PostCaptchaVerfifyRequest struct {
	CaptchaID string `json:"captcha_id" binding:"required"`
	Solution  string `json:"solution" binding:"required"`
}
