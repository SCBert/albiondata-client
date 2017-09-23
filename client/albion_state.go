package client

import (
	"github.com/SCBert/albiondata-client/lib"
)

type albionState struct {
	LocationId           int
	CharacterId          lib.CharacterID
	CharacterName        string
	ContainerItemsToSend map[int64]lib.ItemContainer
}
