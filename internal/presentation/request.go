package presentation

type CPFRequest struct {
	Number string `json:"number" binding:"required"`
	Type   string `json:"type" binding:"required,oneof=CPF CNPJ"`
}

type CPFResponse struct {
	ID      uint   `json:"id"`
	Number  string `json:"number"`
	Type    string `json:"type"`
	Blocked bool   `json:"blocked"`
}
