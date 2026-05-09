package model

type ProdutoItem struct {
	Nome     string  `json:"nome"`
	Preco    float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
	Subtotal   float64 `json:"subtotal,omitempty"`
}
