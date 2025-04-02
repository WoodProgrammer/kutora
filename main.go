package main

import (
	"time"

	serv "github.com/WoodProgrammer/kutora-queue/tcp"
	"github.com/rs/zerolog/log"
)

var INTERVAL = 10
var MemberList = []string{}

func voteMember(memberList *[]string) {

}

func main() {

	raftMember := serv.AppendLogEntry{
		MemberId: "member-1",
	}
	go raftMember.RunServer()
	for {
		go func() {
			command := <-raftMember.CommandChannel
			log.Info().Msgf("The term count is %s", command)
		}()
		log.Info().Msgf("The member : %s is available", raftMember.MemberId)
		time.Sleep(time.Second * time.Duration(INTERVAL))
	}

}
