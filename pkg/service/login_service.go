package service

import (
	"context"
	"errors"
	"github.com/JungleMC/login-service/internal/config"
	"github.com/JungleMC/sdk/pkg/events"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
	"log"
)

var Get *LoginService

type LoginService struct {
	rdb     *redis.Client
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
		err := unmarshalMessage(msg, proto.Message(event))
		if err != nil {
			return err
		}
		return s.onPlayerLoginEvent(event)
	}
	return errors.New("not implemented: " + msg.Channel)
}

func unmarshalMessage(in *redis.Message, out proto.Message) error {
	err := proto.Unmarshal([]byte(in.Payload), out)
	if err != nil {
		return err
	}
	return nil
}
