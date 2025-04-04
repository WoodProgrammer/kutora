package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type AppendLogEntry struct {
	MemberId       string
	LogEntry       LogEntry
	Leader         bool
	CommandChannel chan string
	mtex           sync.Mutex
}

type LogEntry struct {
	Command []string
	Term    int
}

func (node *AppendLogEntry) handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		fmt.Printf("command is %s\n", temp)
		if temp == "SETX" {
			break
		}

		node.CommandChannel <- temp
		c.Write([]byte(string("ok")))
	}

	c.Close()
}

func (node *AppendLogEntry) processRequests() {

	for cmd := range node.CommandChannel {
		node.mtex.Lock()
		node.LogEntry.Term = node.LogEntry.Term + 1
		node.LogEntry.Command = append(node.LogEntry.Command, cmd)
		node.mtex.Unlock()
	}
}

func (node *AppendLogEntry) RunServer(host, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")
	go node.processRequests()
	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		// Handle client connection in a goroutine
		go node.handleConnection(conn)
	}
}
