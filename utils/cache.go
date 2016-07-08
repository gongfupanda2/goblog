package utils

import (
	"bytes"
	"encoding/gob"
	"errors"
	"gopkg.in/redis.v2"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/memcache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/gogather/com/log"
	"time"
)

var cc cache.Cache
var client *redis.Client
func InitCache() {
	cacheConfig := beego.AppConfig.String("cache")
	cc = nil

	if "redis" == cacheConfig {
		initRedis()
	} else {
		initMemcache()
	}

	log.Greenln("[cache] use", cacheConfig)
}

func initMemcache() {
	var err error
	cc, err = cache.NewCache("memcache", `{"conn":"`+beego.AppConfig.String("memcache_host")+`"}`)

	if err != nil {
		beego.Info(err)
	}

}

func initRedis() {
	// cc = &cache.Cache{}
	//var err error
	client = redis.NewClient(&redis.Options{
            Network : beego.AppConfig.String("Network"),
            Addr:     beego.AppConfig.String("RedisAddr"),
            DB:       0,
    })
}

func SetCache(key string, value interface{}, timeout time.Duration) error {
	data, err := Encode(value)
	if err != nil {
		return err
	}
	err1:=client.Set(key, string(data))
	//log.Redln(string(data))
	
	
	if err1 != nil {
		log.Warnln("Cache失败，key:", key)
		return err
	} else {
		log.Blueln("Cache成功，key:", key)
		
		return nil
	}
}

func GetCache(key string, to interface{}) error {
	 data ,err:= client.Get(key).Result()
	//data := cc.Get(key)
	if data == "" {
		return errors.New("Cache不存在")
	}
	// log.Pinkln(data)
	err = Decode([]byte(data), to)
	if err != nil {
		log.Warnln("获取Cache失败", key, err)
	} else {
		log.Greenln("获取Cache成功", key)
		log.Greenln("获取Cache成功", to)
	}

	return err
}

func DelCache(key string) error {
	

	_,err := client.Del(key).Result()
	if err != nil {
		return errors.New("Cache删除失败")
	} else {
		log.Pinkln("删除Cache成功 " + key)
		return nil
	}
}

// --------------------
// Encode
// 用gob进行数据编码
//
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// -------------------
// Decode
// 用gob进行数据解码
//
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
