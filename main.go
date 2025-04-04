package main

import (
	"os"

	serv "github.com/WoodProgrammer/kutora-queue/tcp"
	"github.com/rs/zerolog/log"
)

var INTERVAL = 1
var MemberList = []string{}

func main() {
	port := os.Getenv("RAFT_PORT")
	host := os.Getenv("RAFT_HOST")
	memberId := os.Getenv("RAFT_MEMBER_ID")
	log.Info().Msgf("The raft member is starting ... %s:%s -- memberId: %s", host, port, memberId)
	raftMember := serv.Node{
		MemberId:       memberId,
		CommandChannel: make(chan string),
	}
	go raftMember.RunServer(host, port)

	select {}
}
