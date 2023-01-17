package databases

import (
	"time"

	"github.com/go-redis/redis"
)

var SessionDB *redis.Client

func GetData(key string) (string, error) {
	val, err := SessionDB.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, err
}

func SetData(key, value string, exp time.Duration) error {
	err := SessionDB.Set(key, value, exp).Err()
	if err != nil {
		return err
	}

	return nil
}

func CheckData(key string) bool {
	v, err := SessionDB.Exists(key).Result()
	if err != nil || v == 0 {
		return false
	}
	return true
}

func RemoveData(key string) error {
	return SessionDB.Del(key).Err()
}
