package service

import (
	"github.com/JungleMC/login-service/internal/config"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
)

var Get *LoginService

type LoginService struct {
	RDB *redis.Client
}

func Start(rdb *redis.Client) {
	Get = &LoginService{
		RDB: rdb,
	}

	config.Get = &config.Config{}
	if err := env.Parse(config.Get); err != nil {
		panic(err)
	}
}
