package api

import (
	"main/db"

	"github.com/gin-gonic/gin"
)

//Setup ...
func Setup(router *gin.Engine) {
	db.SetupDB()
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransaction(router)
}
