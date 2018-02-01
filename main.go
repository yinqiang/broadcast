package main

import (
	"flag"
	"fmt"
	"net"
)

func getIp() net.IP {
	addrs, _ := net.InterfaceAddrs()
	var ip net.IP = nil
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP
			}
		}
	}
	return ip
}

func broadcast(msg string, addr *net.UDPAddr) {
	laddr := net.UDPAddr{
		IP: getIp(),
	}
	conn, err := net.DialUDP("udp", &laddr, addr)
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close()
	conn.Write([]byte(msg))
}

func main() {
	msg := flag.String("m", "", "Message to send")
	addr := flag.String("addr", "255.255.255.255", "IP Addr")
	port := flag.Int("port", 8703, "Port")
	flag.Parse()

	if len(*msg) == 0 {
		flag.Usage()
		return
	}
	raddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", *addr, *port))
	if err != nil {
		panic(err)
		return
	}
	broadcast(*msg, raddr)
}
