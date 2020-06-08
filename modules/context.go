package modules

import (
	"fmt"
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
	var redisConfig RedisConfig
	if err := config.Get("config", "redis").Scan(&redisConfig); err != nil {
		panic(err)
	}
	ModuleContext.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", redisConfig.Address, redisConfig.Port),
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	_, err := ModuleContext.Redis.Ping().Result()
	if err != nil {
		panic(err)
	}

	//  captcha-store
	var captchaConfig CaptchaConfig
	if err := config.Get("config", "captcha").Scan(&captchaConfig); err != nil {
		panic(err)
	}
	store := store.NewRedisCaptchaStore(ModuleContext.Redis, captchaConfig.MaxCollectNum, time.Duration(10)*time.Minute)
	captcha.SetCustomStore(store)

	//  captcha-srv
	m_service := micro.NewService()
	ModuleContext.CaptchaSrvClient = proto.NewCaptchaService("go.micro.srv.captcha", m_service.Client())
	if ModuleContext.CaptchaSrvClient == nil {
		panic("captcha-srv client init failed")
	}
}
