package service

import (
	"context"
	"errors"
	"github.com/JungleMC/login-service/internal/config"
	"github.com/JungleMC/sdk/pkg/redis/messages"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
	"log"
)

var Get *LoginService

type LoginService struct {
	rdb  *redis.Client
	channel <-chan *redis.Message
}

func Start(rdb *redis.Client) {
	Get = &LoginService{
		rdb: rdb,
	}

	config.Get = &config.Config{}
	if err := env.Parse(config.Get); err != nil {
		panic(err)
	}

	Get.Bootstrap()
}

func (s *LoginService) Bootstrap() {
	s.channel = s.rdb.PSubscribe(context.Background(), "login.*").Channel()
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
	case "login.begin":
		return s.onLoginBegin(msg)
	}
	return errors.New("not implemented: " + msg.Channel)
}

func (s *LoginService) onLoginBegin(m *redis.Message) error {
	msg := &messages.LoginBegin{}
	err := msg.UnmarshalBinary([]byte(m.Payload))
	if err != nil {
		return err
	}
	log.Println(msg.Username)
	return nil
}
