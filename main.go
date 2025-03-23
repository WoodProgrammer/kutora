package main

import (
	"fmt"

	handlers "github.com/WoodProgrammer/handlers"
	runbook "github.com/WoodProgrammer/runbook"

	"github.com/gin-gonic/gin"
)

func newDBClient() runbook.DB {
	return &runbook.DBClient{}
}

func newKutoraAPIClient(db runbook.DB) handlers.KutoraAPI {
	return &handlers.KutoraAPIHandler{DbClient: db}
}

func main() {
	dbClient := newDBClient()
	kutoraHandler := newKutoraAPIClient(dbClient)

	fmt.Println("The Kutora API v0.0.1")
	router := gin.Default()
	router.GET("/runbooks", kutoraHandler.GetKutoraRunbooks)
	router.Run()
}
