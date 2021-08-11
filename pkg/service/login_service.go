package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/JungleMC/login-service/internal/config"
	"github.com/JungleMC/sdk/pkg/events"
	"github.com/JungleMC/sdk/pkg/messages"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
	"log"
)

var Instance *LoginService

type LoginService struct {
	rdb     *redis.Client
	channel <-chan *redis.Message
}

func Start() {
	config.Get = &config.Config{}
	err := env.Parse(config.Get)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Get.RedisHost, config.Get.RedisPort),
		Password: config.Get.RedisPassword,
		DB:       config.Get.RedisDatabase,
	})
	defer rdb.Close()

	Instance = &LoginService{
		rdb: rdb,
	}

	Instance.Bootstrap()
}

func (s *LoginService) Bootstrap() {
	s.channel = s.rdb.PSubscribe(context.Background(), "event.*").Channel()
	for msg := range s.channel {
		err := s.onMessage(msg)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func (s *LoginService) onMessage(msg *redis.Message) error {
	switch msg.Channel {
	case "event.login":
		event := &events.PlayerLoginEvent{}
		err := messages.UnmarshalMessage(msg, proto.Message(event))
		if err != nil {
			return err
		}
		return s.onPlayerLoginEvent(event)
	case "event.login.response":
		return nil // don't care
	}
	return errors.New("not implemented: " + msg.Channel)
}
