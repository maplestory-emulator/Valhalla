package game

import (
	"net"

	"github.com/Hucaru/Valhalla/constant/opcode"
	"github.com/Hucaru/Valhalla/mnet"
	"github.com/Hucaru/Valhalla/mpacket"
)

type world struct {
	conn          mnet.Server
	Icon          byte
	Name, Message string
	Ribbon        byte
	Channels      []channel
}

func (w *world) AddChannel(ch channel) bool {
	if len(w.Channels) > 19 {
		return false
	}

	w.Channels = append(w.Channels, ch)
	return true
}

func (w *world) generateInfoPacket() mpacket.Packet {
	p := mpacket.CreateInternal(opcode.WorldInfo)
	p.WriteByte(w.Icon)
	p.WriteString(w.Name)
	p.WriteString(w.Message)
	p.WriteByte(w.Ribbon)
	p.WriteByte(byte(len(w.Channels)))

	for _, v := range w.Channels {
		p.WriteBytes(v.generatePacket())
	}

	return p
}

func (w *world) serialisePacket(reader mpacket.Reader) {
	w.Icon = reader.ReadByte()
	w.Name = reader.ReadString(int(reader.ReadInt16()))
	w.Message = reader.ReadString(int(reader.ReadInt16()))
	w.Ribbon = reader.ReadByte()

	nOfChannels := int(reader.ReadByte())
	w.Channels = make([]channel, nOfChannels)

	for i := 0; i < nOfChannels; i++ {
		w.Channels[i].serialisePacket(reader)
	}
}

type channel struct {
	IP          net.IP
	Port        int16
	MaxPop, Pop int16
}

func (c channel) generatePacket() mpacket.Packet {
	p := mpacket.NewPacket()
	p.WriteBytes(c.IP)
	p.WriteInt16(c.Port)
	p.WriteInt16(c.MaxPop)
	p.WriteInt16(c.Pop)
	return p
}

func (c *channel) serialisePacket(reader mpacket.Reader) {
	c.IP = reader.ReadBytes(4)
	c.Port = reader.ReadInt16()
	c.MaxPop = reader.ReadInt16()
	c.Pop = reader.ReadInt16()
}