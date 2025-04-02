package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type AppendLogEntry struct {
	MemberId       string
	LogEntry       LogEntry
	Leader         bool
	CommandChannel chan string
}

type LogEntry struct {
	Command string
	Term    string
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
		switch temp {
		case "BYE":
			fmt.Printf("MEMBER LEFT")
		case "HEART_BEAT":
			node.CommandChannel <- temp
		default:
			node.CommandChannel <- temp
		}

		c.Write([]byte(string("ok")))
		c.Close()
	}

}

func (node *AppendLogEntry) RunServer() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

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
