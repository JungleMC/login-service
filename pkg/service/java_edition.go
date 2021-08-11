package service

import (
	"context"
	"github.com/JungleMC/sdk/pkg/events"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"log"
)

func (s *LoginService) onPlayerLoginEvent(event *events.PlayerLoginEvent) error {
	profileId, _ := uuid.FromBytes(event.ProfileId)

	msg := &events.PlayerLoginResponse{
		ClientType:        event.GetClientType(),
		NetworkId:         event.GetNetworkId(),
		ProfileId:         event.GetProfileId(),
		Username:          event.GetUsername(),
		Success:           true, // TODO: for bans
		DisconnectMessage: nil,
	}

	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	pub := s.rdb.Publish(context.Background(), "event.login.response", msgBytes)
	if pub.Err() != nil {
		return pub.Err()
	}

	log.Printf("[Profile ID: %v]: %v joined the game\n", profileId.String(), event.Username)
	return nil
}
