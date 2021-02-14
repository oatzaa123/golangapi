package api

import (
	"fmt"
	"main/db"
	"main/middleware"
	"main/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionResult struct {
	ID            uint
	Total         float64
	Paid          float64
	Change        float64
	PaymentType   string
	PaymentDetail string
	OrderList     string
	Staff         string
	CreateAt      time.Time
}

//SetupTransaction ...
func SetupTransaction(router *gin.Engine) {
	TransactionAPI := router.Group("api/v2")
	{
		TransactionAPI.GET("/transaction", middleware.JwtVerify, getTransaction)
		TransactionAPI.POST("/transaction", middleware.JwtVerify, createTransaction)
	}
}

//Join Table
func getTransaction(c *gin.Context) {
	var result []TransactionResult
	db.GetDB().Debug().Raw("SELECT transactions.id, total, paid, change, payment_type, payment_detail, order_list, users.username as Staff, transactions.create_at FROM transactions join users on transactions.staff_id = users.id", nil).Scan(&result)
	c.JSON(200, gin.H{"result": result})
}

func createTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBind(&transaction); err == nil {
		transaction.StaffID = c.GetString("jwt_staff_id")
		fmt.Println("xxx", c.GetString("jwt_staff_id"))
		transaction.CreateAt = time.Now()
		db.GetDB().Create(&transaction)
		c.JSON(http.StatusOK, gin.H{"result": "ok", "data": transaction})
	} else {
		c.JSON(404, gin.H{"result": "nok", "error": err})
	}
}
