package main

import (
	"flag"
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

func broadcast(msg string) {
	laddr := net.UDPAddr{
		IP:   getIp(),
		Port: 8703,
	}
	raddr := net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 8703,
	}
	conn, err := net.DialUDP("udp", &laddr, &raddr)
	if err != nil {
		panic(err)
		return
	}
	conn.Write([]byte(msg))
	conn.Close()
}

func main() {
	msg := flag.String("m", "", "Message to send")
	flag.Parse()

	if len(*msg) == 0 {
		flag.Usage()
		return
	}
	broadcast(*msg)
}
