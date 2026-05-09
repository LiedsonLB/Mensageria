package model

type Pedido struct {
	ID          string        `json:"id,omitempty"`
	Cliente   string        `json:"cliente"`
	Documento   string        `json:"documento"`
	Email     string        `json:"email"`
	Produtos  []ProdutoItem `json:"produtos"`
	ValorTotal float64      `json:"valor_total"`
	Status    string        `json:"status"`
}