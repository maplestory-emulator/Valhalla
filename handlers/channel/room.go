package channel

import (
	"fmt"

	"github.com/Hucaru/Valhalla/game"
	"github.com/Hucaru/Valhalla/mnet"
	"github.com/Hucaru/Valhalla/mpacket"
)

const (
	roomCreate                = 0
	roomSendInvite            = 2
	roomReject                = 3
	roomAccept                = 4
	roomChat                  = 6
	roomCloseWindwow          = 10
	roomInsertItem            = 13
	roomMesos                 = 14
	roomAcceptTrade           = 16
	roomRequestTie            = 42
	roomRequestTieResult      = 43
	roomRequestGiveUp         = 44
	roomRequestUndo           = 46
	roomRequestUndoResult     = 47
	roomRequestExitDuringGame = 48
	roomReadyButtonPressed    = 50
	roomUnready               = 51
	roomOwnerExpells          = 52
	roomGameStart             = 53
	roomChangeTurn            = 55
	roomPlacePiece            = 56
)

func handleUIWindow(conn mnet.MConnChannel, reader mpacket.Reader) {
	operation := reader.ReadByte()

	player := game.Players[conn]

	switch operation {
	case roomCreate:
		if player.RoomID > 0 {
			return
		}
	case roomSendInvite:
	case roomReject:
	case roomAccept:
	case roomChat:
	case roomCloseWindwow:
	case roomInsertItem:
	case roomMesos:
	case roomAcceptTrade:
	case roomRequestTie:
	case roomRequestTieResult:
	case roomRequestGiveUp:
	case roomRequestUndo:
	case roomRequestUndoResult:
	case roomRequestExitDuringGame:
	case roomReadyButtonPressed:
	case roomUnready:
	case roomOwnerExpells:
	case roomGameStart:
	case roomChangeTurn:
	case roomPlacePiece:
	default:
		fmt.Println("Unknown room operation", operation)
	}
}