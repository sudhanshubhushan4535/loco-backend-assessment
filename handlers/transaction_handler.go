package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sudhanshubhushan4535/loco-backend-assessment/models"
)

var validate = validator.New()

var transactions = make(map[int64]models.Transaction)
var parentToChildren = make(map[int64][]int64)

func CreateTransaction(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction models.Transaction
	if err := c.ShouldBind(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction.ID = transactionID
	transactions[transactionID] = transaction

	if transaction.ParentID != nil {
		parentToChildren[*transaction.ParentID] = append(parentToChildren[*transaction.ParentID], transactionID)
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetTransaction(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, exists := transactions[transactionID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func GetTransactionsByType(c *gin.Context) {
	transactionsType := c.Param("type")

	var transactionIDs []int64
	for id, transaction := range transactions {
		if transaction.Type == transactionsType {
			transactionIDs = append(transactionIDs, id)
		}
	}
	c.JSON(http.StatusOK, transactionIDs)
}

func GetTransactionSum(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	sum := calculateTransactionSum(transactionID)

	c.JSON(http.StatusOK, sum)
}

func calculateTransactionSum(transactionID int64) float64 {
	transaction, exist := transactions[transactionID]
	if !exist {
		return 0
	}

	sum := transaction.Amount

	if children, hasChildren := parentToChildren[transactionID]; hasChildren {
		for _, id := range children {
			sum += calculateTransactionSum(id)
		}
	}

	return sum
}
