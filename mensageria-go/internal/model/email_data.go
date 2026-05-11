package model

type EmailData struct {
	PedidoID string
	Cliente  string
	Produtos   []ProdutoItem
	ValorTotal    string
	Data     string
	NotaURL    string
	DANFE_URL  string
	ChaveAcesso string
	Protocolo   string
	Ano         int
	LogoBase64  string
}