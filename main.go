package main

import (
	"os"

	serv "github.com/WoodProgrammer/kutora-queue/tcp"
)

var INTERVAL = 1
var MemberList = []string{}

func main() {
	port := os.Getenv("RAFT_PORT")
	host := os.Getenv("RAFT_HOST")
	memberId := os.Getenv("MEMBER_ID")

	raftMember := serv.AppendLogEntry{
		MemberId:       memberId,
		CommandChannel: make(chan string),
	}
	go raftMember.RunServer(host, port)

	select {}
}
