package service

import (
	"github.com/JungleMC/sdk/pkg/events"
	"github.com/google/uuid"
	"log"
)

func (s *LoginService) onPlayerLoginEvent(event *events.PlayerLoginEvent) error {
	networkId, _ := uuid.FromBytes(event.NetworkId)
	profileId, _ := uuid.FromBytes(event.ProfileId)
	log.Printf("[Network ID: %s] [Profile ID: %s]: %s joined the game\n", networkId.String(), profileId.String(), event.Username)
	return nil
}
