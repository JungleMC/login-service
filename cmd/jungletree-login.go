package main

import (
	"github.com/JungleMC/login-service/pkg/service"
	"github.com/JungleMC/sdk/pkg/redis"
)

func main() {
	service.Start(redis.NewClient())
}
