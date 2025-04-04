package tcp

import (
	"bufio"
	"fmt"
	"net"
)

func (node *Node) TcpClient(msg, host string) {

	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}

	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Print("Server reply: ", reply)
}
