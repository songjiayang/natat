package stun

import (
	"github.com/gortc/stun"
	"net"
)

func Ping(bind, stud *net.UDPAddr) (xorAddr *stun.XORMappedAddress, err error) {
	conn, err := net.DialUDP("udp", bind, stud)
	if err != nil {
		return
	}

	c, err := stun.NewClient(conn)
	if err != nil {
		return
	}

	defer c.Close()

	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	err = c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			err = res.Error
			return
		}

		xorAddr = &stun.XORMappedAddress{}
		err = xorAddr.GetFrom(res.Message)
	})

	return
}
