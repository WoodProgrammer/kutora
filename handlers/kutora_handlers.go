package handlers

import (
	"fmt"

	runbook "github.com/WoodProgrammer/runbook"
	"github.com/gin-gonic/gin"
)

type KutoraAPI interface {
	GetKutoraRunbooks(c *gin.Context)
}

type KutoraAPIHandler struct {
	DbClient runbook.DB
}

func (k *KutoraAPIHandler) GetKutoraRunbooks(c *gin.Context) {
	result, err := k.DbClient.GetRunBooks()
	if err != nil {

		return
	}
	fmt.Println(result)
}
