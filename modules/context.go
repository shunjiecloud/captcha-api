package modules

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dchest/captcha"
	"github.com/go-redis/redis/v7"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/shunjiecloud-proto/captcha/proto"
	"github.com/shunjiecloud/captcha-srv/store"
)

type moduleWrapper struct {
	Redis            *redis.Client
	CaptchaSrvClient proto.CaptchaService
}

//ModuleContext 模块上下文
var ModuleContext moduleWrapper

//Setup 初始化Modules
func Setup() {
	//  redis
	var host Host
	if err := config.Get("hosts", "redis").Scan(&host); err != nil {
		panic(err)
	}
	ModuleContext.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", host.Address, host.Port),
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	_, err := ModuleContext.Redis.Ping().Result()
	if err != nil {
		panic(err)
	}

	//  captcha-store
	strMaxNum := os.Getenv("MAX_CAPTCHA_COLLECT_NUM")
	if len(strMaxNum) == 0 {
		panic("MAX_CAPTCHA_COLLECT_NUM not set")
	}
	maxNum, err := strconv.ParseUint(strMaxNum, 10, 64)
	if err != nil {
		panic("MAX_CAPTCHA_COLLECT_NUM is not a number")
	}
	store := store.NewRedisCaptchaStore(ModuleContext.Redis, maxNum, time.Duration(10)*time.Minute)
	captcha.SetCustomStore(store)

	//  captcha-srv
	m_service := micro.NewService()
	ModuleContext.CaptchaSrvClient = proto.NewCaptchaService("go.micro.srv.captcha", m_service.Client())
	if ModuleContext.CaptchaSrvClient == nil {
		panic("captcha-srv client init failed")
	}
}
