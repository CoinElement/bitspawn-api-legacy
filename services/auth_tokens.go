/*

 */

package services

import (
	"github.com/go-redis/redis"
	"time"
)

type AuthTokenStore struct {
	redisClient *redis.Client
}

func NewAuthTokenStore(addr string) *AuthTokenStore {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       8,  // use default DB
	})

	return &AuthTokenStore{
		redisClient: client,
	}
}

func (a *AuthTokenStore) SetToken(token string, username string) error {
	return a.redisClient.Set(token, username, time.Hour*24*30).Err()
}

func (a *AuthTokenStore) GetUser(token string) (string, error) {
	getCmd := a.redisClient.Get(token)

	return getCmd.Result()
}
