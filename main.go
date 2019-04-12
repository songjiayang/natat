package main

import (
	"flag"
	gostun "github.com/gortc/stun"
	"github.com/songjiayang/natat/assert"
	"github.com/songjiayang/natat/stun"
	"log"
	"net"
	"strings"
)

var (
	bind, studs string
)

func init() {
	flag.StringVar(&bind, "bind", "0.0.0.0:3489", "ping with local address bind.")
	flag.StringVar(&studs, "studs", "stun.l.google.com:19302,stun1.l.google.com:19302", "stun servers for ping.")
}

func main() {
	flag.Parse()

	fields := strings.Split(studs, ",")
	if len(fields) < 2 {
		log.Panic("missing stund server configration.")
	}

	// resolve UDP address
	bindAddr := resolveUDPAddr(bind)
	stun1Addr := resolveUDPAddr(fields[0])
	stun2Addr := resolveUDPAddr(fields[1])

	log.Println("start stun server ping...")
	// send STUN ping request
	xorAddr := ping(bindAddr, stun1Addr)
	xorAddr2 := ping(bindAddr, stun2Addr)

	log.Printf("%s mapped: %s -> %s\n", fields[0], bindAddr.String(), xorAddr.String())
	log.Printf("%s mapped: %s -> %s\n", fields[1], bindAddr.String(), xorAddr2.String())

	log.Println("start NAT type assert...")
	natAssert(bindAddr, xorAddr, xorAddr2)
}

func resolveUDPAddr(addr string) *net.UDPAddr {
	uAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Panicf("%s resolved failed: %s", addr, err.Error())
	}

	return uAddr
}

func ping(bindAddr, stunAddr *net.UDPAddr) *gostun.XORMappedAddress {
	xorAddr, err := stun.Ping(bindAddr, stunAddr)

	if err != nil {
		log.Panicf("stun.Ping(%s, %s): %s", bindAddr.String(), stunAddr.String(), err.Error())
	}

	return xorAddr
}

func natAssert(bindAddr *net.UDPAddr, xorAddr, xorAddr2 *gostun.XORMappedAddress) {
	var natType = "Symmetric"

	switch {
	case assert.IsSymmetric(xorAddr, xorAddr2):

	default:
		natType = "Cone"
	}

	log.Printf("It's %s NAT\n", natType)
}
