package presentation

import (
	"cpf-cnpj-api/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CPFHandler struct {
	service *services.CpfCnpjService
}

func NewCPFHandler(service *services.CpfCnpjService) *CPFHandler {
	return &CPFHandler{service: service}
}

func (h *CPFHandler) CreateCPF(c *gin.Context) {
	var req CPFRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data."})
		return
	}

	cpfCnpj, err := h.service.CreateCpfCnpj(req.Number, req.Type)
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
	cpfs, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cpfs)
}
