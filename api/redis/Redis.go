package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
)

var Pool redis.Pool

func SetPool(redisPool *redis.Pool) {
	Pool = *redisPool
}

func Get(key string) *string {
	if &Pool != nil {
		conn := Pool.Get()
		defer conn.Close()
		var value string
		value, err := redis.String(conn.Do("GET", key))
		if err != nil {
			logrus.Error(err)
			return nil
		}
		return &value
	}
	return nil
}

func Set(key string, value string) {
	if &Pool != nil {
		conn := Pool.Get()
		defer conn.Close()
		_, err := conn.Do("SET", key, value)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func SetEx(key string, ex int, value string) {
	if &Pool != nil {
		conn := Pool.Get()
		defer conn.Close()
		_, err := conn.Do("SETEX", key, ex, value)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func Exists(key string) bool {
	if &Pool != nil {
		conn := Pool.Get()
		defer conn.Close()
		var exist bool
		exist, err := redis.Bool(conn.Do("EXISTS", key))
		if err != nil {
			logrus.Error(err)
			return false
		}
		return exist
	}
	return false
}
