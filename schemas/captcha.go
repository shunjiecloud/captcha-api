package schemas

type GetCaptchaResponse struct {
	CaptchaId string `json:"captcha_id"`
	Url       string `json:"captcha_url"`
}

type PostCaptchaVerfifyRequest struct {
	CaptchaId string `json:"captcha_id" binding:"required"`
	Solution  string `json:"solution" binding:"required"`
}
