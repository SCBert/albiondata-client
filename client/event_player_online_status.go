package client

import (
	"github.com/SCBert/albiondata-client/lib"
	"github.com/SCBert/albiondata-client/log"
)

type eventPlayerOnlineStatus struct {
	CharacterID   lib.CharacterID `mapstructure:"0"`
	CharacterName string          `mapstructure:"1"`
	IsOnline      bool            `mapstructure:"2"`
}

func (event eventPlayerOnlineStatus) Process(state *albionState) {
	log.Debug("Got player online status event...")
}
