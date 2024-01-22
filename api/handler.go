package api

import (
	"Credit-Customer-Golang-Test/models"
	"Credit-Customer-Golang-Test/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo *repositories.Repository
}

func NewHandler(repo *repositories.Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) CreateConsumer(c *gin.Context) {
	var consumer models.Consumer
	if err := c.BindJSON(&consumer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newConsumer, err := h.Repo.CreateConsumer(consumer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newConsumer)
}
func (h *Handler) CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTransaction, err := h.Repo.CreateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newTransaction)
}
