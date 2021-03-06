package handlers

import (
	"log"

	"github.com/Hucaru/Valhalla/connection"
	"github.com/Hucaru/Valhalla/consts/opcodes"
	"github.com/Hucaru/Valhalla/maplepacket"
)

// HandleLoginPacket -
func HandleLoginPacket(conn *connection.Login, reader maplepacket.Reader) {
	switch reader.ReadByte() {
	case opcodes.Recv.ReturnToLoginScreen:
		handleReturnToLoginScreen(conn, reader)

	case opcodes.Recv.LoginRequest:
		handleLoginRequest(conn, reader)

	case opcodes.Recv.LoginCheckLogin:
		handleGoodLogin(conn, reader)

	case opcodes.Recv.LoginWorldSelect:
		handleWorldSelect(conn, reader)

	case opcodes.Recv.LoginChannelSelect:
		handleChannelSelect(conn, reader)

	case opcodes.Recv.LoginNameCheck:
		handleNameCheck(conn, reader)

	case opcodes.Recv.LoginNewCharacter:
		handleNewCharacter(conn, reader)

	case opcodes.Recv.LoginDeleteChar:
		handleDeleteCharacter(conn, reader)

	case opcodes.Recv.LoginSelectCharacter:
		handleSelectCharacter(conn, reader)

	default:
		log.Println("UNKNOWN LOGIN PACKET:", reader)
	}

}
