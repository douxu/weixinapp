package common

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

//VerifyToken : func of Verify token
func VerifyToken(username, token string) (result bool) {
	result = false
	rc := RedisClient.Get()
	value, err := redis.String(rc.Do("GET", username))
	if err != nil {
		log.Fatal(err)
	}
	if value == token {
		result = true
	}
	defer rc.Close()
	return result
}
