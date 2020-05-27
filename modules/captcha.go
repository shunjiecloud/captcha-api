package modules

import (
	"time"
)

//RedisCaptchaStorePrefix 验证码种子集合前缀
const RedisCaptchaStorePrefix = "captcha:collects:"

//RedisCaptchaMaxNumKey 保有的验证码最大数量
const RedisCaptchaMaxNumKey = "captcha:max"

//RedisCaptchaCurNumKey 当前验证码数量
const RedisCaptchaCurNumKey = "captcha:cur"

type RedisCaptchaStore struct {
	collectNum int
	expiration time.Duration
}

func (store *RedisCaptchaStore) Set(id string, digits []byte) {
	//  检查key数量是否超过上限
	curNum, err := ModuleContext.Redis.Get(RedisCaptchaCurNumKey).Int64()
	if err != nil {

	}
	maxNum, err := ModuleContext.Redis.Get(RedisCaptchaMaxNumKey).Int64()
	if err != nil {

	}
	if curNum >= maxNum {
		//  超数，服务不可用
		return
	}
	//  设置key，计数器加1
	key := RedisCaptchaStorePrefix + id
	_, err = ModuleContext.Redis.Set(key, digits, store.expiration).Result()
	if err != nil {

	}
	_, err = ModuleContext.Redis.Incr(RedisCaptchaCurNumKey).Result()
	if err != nil {

	}
}

func (store *RedisCaptchaStore) Get(id string, clear bool) (digits []byte) {
	digits = make([]byte, 0)
	key := RedisCaptchaStorePrefix + id
	ret, err := ModuleContext.Redis.Get(key).Result()
	if err != nil {
		return nil
	}
	digits = []byte(ret)

	if clear == true {
		//  clear为true，删除id，计数器减1
		_, err = ModuleContext.Redis.Del(key).Result()
		if err != nil {

		}
		_, err = ModuleContext.Redis.Decr(RedisCaptchaCurNumKey).Result()
		if err != nil {

		}
	}
	return digits
}

func NewRedisCaptchaStore(collectNum int, expiration time.Duration) *RedisCaptchaStore {
	store := RedisCaptchaStore{
		collectNum: collectNum,
		expiration: expiration,
	}
	return &store
}
