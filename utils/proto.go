package utils

import (
	"github.com/golang/protobuf/proto"
	"github.com/wsboyxb/robot/msg"
)

func ActionToData(a, m, cb string, data []byte) ([]byte, error) {
	req := &msg.HMRequest{
		Action:   proto.String(a),
		Method:   proto.String(m),
		Data:     data,
		Callback: proto.String(cb),
		//ChannelId: proto.Int32(int32(conf.ChannelID)),
	}

	msgByte, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	return msgByte, nil
}
