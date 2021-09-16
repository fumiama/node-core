package link

import (
	"fmt"

	"github.com/IPoWS/node-core/ip64"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func Send(to uint64, data *[]byte, proto uint16, srcport uint16, destport uint16) ([]byte, error) {
	wsn, ok := readMap(to)
	if ok {
		var ip ip64.Ip64
		ip.Pack(Mywsip, to, data, proto, srcport, destport)
		logrus.Infof("[Send] link send %d bytes to %x.", len(*data), to)
		d, err := ip.Send(wsn, websocket.BinaryMessage, nil)
		if err != nil {
			DelConn(to)
		}
		return d, err
	}
	return nil, fmt.Errorf("dest %x unreachable.", to)
}

func Forward(to uint64, ip *ip64.Ip64) ([]byte, error) {
	wsn, ok := readMap(to)
	if ok {
		d, err := ip.Send(wsn, websocket.BinaryMessage, nil)
		if err != nil {
			DelConn(to)
		}
		return d, err
	}
	return nil, fmt.Errorf("dest %x unreachable.", to)
}
