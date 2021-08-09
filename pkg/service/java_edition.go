package service

import (
	"github.com/JungleMC/sdk/pkg/events"
	"github.com/google/uuid"
	"log"
)

func (s *LoginService) onPlayerLoginEvent(event *events.PlayerLoginEvent) error {
	profileId, _ := uuid.FromBytes(event.ProfileId)
	log.Printf("[Profile ID: %v]: %v joined the game\n", profileId.String(), event.Username)
	return nil
}
