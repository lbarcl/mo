package protocol

import (
	"fmt"
	"helper"
)

func HandleHandshake(state *helper.States, packet Packet) {
	switch packet.id {
	case 0x00:
		// Parse the second VarInt
		protocolversion, err := packet.ReadVarInt()
		if err != nil {
			fmt.Println("Error reading PV VarInt:", err)
			break
		}

		if protocolversion != helper.ProtocolVersion {
			(*packet.sender).Close()
			*state = helper.Closed
			break
		}

		serverAdress := packet.ReadString()

		serverPort, err := packet.ReadShort()
		if err != nil {
			fmt.Println("Error reading sp VarInt:", err)
			break
		}

		nextState, err := packet.ReadVarInt()
		if err != nil {
			fmt.Println("Error reading NS VarInt:", err)
			break
		}

		fmt.Println("[Handshake]", "id:", packet.id, "protocol version:", protocolversion, "server address:", serverAdress, "server port:", serverPort, "next state:", nextState)

		*state = helper.States(nextState)

	default:
		fmt.Println("Unknown packet ID in Handshaking state:", packet.id)
	}
}