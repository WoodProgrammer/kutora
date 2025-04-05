package tcp

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type Node struct {
	MemberId       string
	LogEntry       LogEntry
	State          string
	CommandChannel chan string
	MemberList     []string
	mtex           sync.Mutex
}

type LogEntry struct {
	Command []string
	Term    int
}

func (node *Node) handleConnection(c net.Conn) {

	defer c.Close()

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		n, err := c.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		temp := strings.TrimSpace(string(buffer[:n]))
		node.CommandChannel <- temp
		// Process and use the data (here, we'll just print it)
		fmt.Printf("Received: %s\n", buffer[:n])
	}

}

func (node *Node) processRequests() {

	for cmd := range node.CommandChannel {
		node.mtex.Lock()
		node.LogEntry.Term = node.LogEntry.Term + 1
		node.LogEntry.Command = append(node.LogEntry.Command, cmd)
		node.mtex.Unlock()
		if node.State == "leader" {
			fmt.Println("replicating", cmd)
			node.TcpClient(cmd, "localhost:8081")
		} else {
			fmt.Println("ignore", cmd)
		}
	}
}

func (node *Node) RunServer(host, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")
	go node.processRequests()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go node.handleConnection(conn)

	}
}
