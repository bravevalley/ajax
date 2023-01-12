package databases

import (
	"time"

	"github.com/go-redis/redis"
)

var SessionDB *redis.Client

func getData(key string) (string, error) {
	val, err := SessionDB.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, err
}

func setData(key, value string, exp time.Duration) error {
	err := SessionDB.Set(key, value, exp).Err()
	if err != nil {
		return err
	}

	return nil
}
