package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: tcp_scan <ip> <port_range>")
		os.Exit(1)
	}

	ip := os.Args[1]
	portRange := os.Args[2]

	start, err := strconv.Atoi(portRange[:strings.Index(portRange, "-")])
	if err != nil {
		fmt.Println("Invalid port range:", err)
		os.Exit(1)
	}

	end, err := strconv.Atoi(portRange[strings.Index(portRange, "-")+1:])
	if err != nil {
		fmt.Println("Invalid port range:", err)
		os.Exit(1)
	}

	for port := start; port <= end; port++ {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 5*time.Second)
		if err == nil {
			fmt.Printf("Port %d is open on %s\n", port, ip)
			conn.Close()
		}
	}
}
