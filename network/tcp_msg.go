package network

import (
	"encoding/binary"
	"errors"
	"io"
)

type MsgParser struct {
}

func NewMsgParser() *MsgParser {
	p := new(MsgParser)
	return p
}

func (p *MsgParser) Write(c *TcpClient, data []byte) error {
	msgLen := uint32(len(data))
	buff := make([]byte, msgLen+4)
	binary.BigEndian.PutUint32(buff, msgLen)

	copy(buff[4:], data)
	c.Send(buff)
	return nil
}

func (p *MsgParser) Read(c *TcpClient) ([]byte, error) {
	conn := c.GetConnection()

	var b [4]byte
	bufMsgLen := b[:4]

	// read len
	_, err := io.ReadFull(conn, bufMsgLen)
	if err != nil {
		return nil, err
	}

	// parse len
	var msgLen uint32
	msgLen = binary.BigEndian.Uint32(bufMsgLen)

	// check len
	if msgLen > 8*1024 {
		return nil, errors.New("message too long")
	} else if msgLen < 4 {
		return nil, errors.New("message too short")
	}

	// data
	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(conn, msgData); err != nil {
		return nil, err
	}

	return msgData, nil
}
