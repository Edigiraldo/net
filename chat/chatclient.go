package chat

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func CreateClient() {
	wg := &sync.WaitGroup{}
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go outputToStdout(conn, wg)
	go inputToStdin(conn, wg)
	wg.Add(2)

	wg.Wait()
	conn.Close()
}

func outputToStdout(conn net.Conn, wg *sync.WaitGroup) {
	io.Copy(os.Stdout, conn)
	wg.Done()
}

func inputToStdin(conn net.Conn, wg *sync.WaitGroup) {
	io.Copy(conn, os.Stdin)
	wg.Done()
}
