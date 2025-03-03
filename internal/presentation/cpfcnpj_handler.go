package presentation

import (
	"cpf-cnpj-api/internal/domain/services"
	"net/http"
	"strconv"

	"cpf-cnpj-api/internal/infrastructure"
	"log"

	"github.com/gin-gonic/gin"
)

type CPFHandler struct {
	CpfCnpjService *services.CpfCnpjService
	RabbitMQ       *infrastructure.RabbitMQ
}

func NewCPFHandler(cpfCnpjService *services.CpfCnpjService, rabbitMQ *infrastructure.RabbitMQ) *CPFHandler {
	return &CPFHandler{CpfCnpjService: cpfCnpjService, RabbitMQ: rabbitMQ}
}

func (h *CPFHandler) CreateCPF(c *gin.Context) {
	var req CPFRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data."})
		return
	}

	cpfCnpj, err := h.CpfCnpjService.CreateCpfCnpj(req.Number, req.Type)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := CPFResponse{
		ID:      cpfCnpj.ID,
		Number:  cpfCnpj.Number,
		Type:    string(cpfCnpj.Type),
		Blocked: cpfCnpj.Blocked,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *CPFHandler) GetAllCPF(c *gin.Context) {
	cpfs, err := h.CpfCnpjService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cpfs)
}

func (h *CPFHandler) BlockCPF(c *gin.Context) {
	var input struct {
		ID uint `json:"id"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	message := strconv.FormatUint(uint64(input.ID), 10)
	err := h.RabbitMQ.PublishMessage("block_cpf_queue", message)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue block request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Block request queued"})
}
