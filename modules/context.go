package modules

import (
	"time"

	"github.com/dchest/captcha"
	"github.com/go-redis/redis/v7"
)

type moduleWrapper struct {
	Redis *redis.Client
}

//ModuleContext 模块上下文
var ModuleContext moduleWrapper

//Setup 初始化Modules
func Setup() {
	//  redis
	ModuleContext.Redis = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := ModuleContext.Redis.Ping().Result()
	if err != nil {
		panic(err)
	}

	//  captcha
	store := NewRedisCaptchaStore(10000, time.Duration(10)*time.Minute)
	captcha.SetCustomStore(store)
}
