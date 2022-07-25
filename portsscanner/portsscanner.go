package portsscanner

import (
	"fmt"
	"net"
	"sync"
)

func portScanner(wg *sync.WaitGroup, port int) {
	defer wg.Done()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", port))
	if err != nil {
		return
	}
	conn.Close()
	fmt.Printf("port %d is open\n", port)
}

func RunExample() {
	wg := &sync.WaitGroup{}
	for port := 0; port < 100; port++ {
		wg.Add(1)
		go portScanner(wg, port)
	}
	wg.Wait()
}
