package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	handlers "github.com/WoodProgrammer/handlers"
	runbook "github.com/WoodProgrammer/runbook"

	"github.com/gin-gonic/gin"
)

func newDBClient() runbook.DB {
	db, err := sql.Open("sqlite3", "runbooks.db")
	if err != nil {
		return nil
	}

	return &runbook.DBClient{Dao: db}
}

func newKutoraAPIClient(db runbook.DB) handlers.KutoraAPI {
	return &handlers.KutoraAPIHandler{DbClient: db}
}

func main() {
	dbClient := newDBClient()
	kutoraHandler := newKutoraAPIClient(dbClient)

	router := gin.Default()
	router.GET("/runbooks", kutoraHandler.GetKutoraRunbooks)
	router.Run()
}
