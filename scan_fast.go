package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	_ "sync"
	"time"
	_ "time"
)

func worker(ports, results chan int, ip string) {
	for p := range ports {
		fmt.Sprintf("scanme.nmap.org:%d", ip)
		//conn, err := net.Dial("tcp", address)
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, p), 5*time.Second)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

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
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for port := start; port <= end; port++ {

		go worker(ports, results, ip)

	}
	go func() {
		for i := start; i <= end; i++ {
			ports <- i
		}
	}()

	for i := 0; i < end; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

}
