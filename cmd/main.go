package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sudhanshubhushan4535/loco-backend-assessment/handlers"
)

func main() {
	r := gin.Default()

	// Define API routes
	r.PUT("/transactionservice/transaction/:transaction_id", handlers.CreateTransaction)
	r.GET("/transactionservice/transaction/:transaction_id", handlers.GetTransaction)
	r.GET("/transactionservice/types/:type", handlers.GetTransactionsByType)
	r.GET("/transactionservice/sum/:transaction_id", handlers.GetTransactionSum)

	r.Run(":8080")
}
