package utils

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

var lmcache cache.Cache

func InitCache()  {
	engine := beego.AppConfig.String("cache::engine")
	switch engine {
	case "file":
		fileEngineCache()
	case "redis":
		redisEngineCache()
	}
}

func fileEngineCache()  {

}
// 配置 redis chache
func redisEngineCache() {
	host := beego.AppConfig.String("redis::redis_host")
	var err error
	defer func() {
		if r := recover(); r != nil {
			lmcache = nil
		}
	}()

	lmcache, err = cache.NewCache("redis", `{"conn":"`+ host +`"}`)
	if err != nil {
		LogError("Connect to the redis host " + host + " failed.")
		LogError(err)
	}
}
// 设置缓存数据
func SetCache(key string, value interface{}, timeout int) error {
	data, err := Encode(value)
	if err != nil {
		return err
	}

	if lmcache == nil {
		return errors.New("lmcache is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			LogError(r)
			lmcache = nil
		}
	}()

	timeouts := time.Duration(timeout) * time.Second
	err = lmcache.Put(key, data, timeouts)
	if err != nil {
		LogError(err)
		LogError("SetCache 失败，key: " + key)
	}
	return nil
}
// 获取缓存数据
func GetCache(key string, to interface{}) error {
	if lmcache == nil {
		return errors.New("lmcache is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			LogError(r)
			lmcache = nil
		}
	}()
	data := lmcache.Get(key)
	if data == nil {
		return errors.New("Cache 不存在")
	}
	err := Decode(data.([]byte), to)
	if err != nil {
		LogError(err)
		LogError("GetCache 失败，key: " + key)
	}
	return err
}
// 删除缓存
func DelCache(key string) error {
	if lmcache == nil {
		return errors.New("lmcache is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			// fmt.Println("get cache error caught: %v\n", r)
			lmcache = nil
		}
	}()
	err := lmcache.Delete(key)
	if err != nil {
		return errors.New("Cache 删除失败")
	}
	return nil
}

// 对缓存数据进行编码
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
// 对缓存数据进行解码
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}