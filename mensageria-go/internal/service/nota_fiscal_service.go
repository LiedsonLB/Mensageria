// internal/service/nota_fiscal.go
package service

import (
	"fmt"
	"time"
	"log"

	"mensageria-go/internal/model"
	"mensageria-go/internal/nf"
)

func GerarNota(pedido model.Pedido) (model.NotaFiscal, error) {
	log.Printf("📄 Gerando nota fiscal para pedido: %s", pedido.ID)
	log.Printf("📦 Produtos recebidos: %d", len(pedido.ProdutosList))
	
	for i, p := range pedido.ProdutosList {
		log.Printf("  Produto %d: %s - Qtd: %d - Preço: %.2f", i+1, p.Nome, p.Quantidade, p.Preco)
	}
	
	if len(pedido.ProdutosList) == 0 {
		return model.NotaFiscal{}, fmt.Errorf("pedido sem produtos para gerar nota fiscal")
	}
	
	path, err := nf.GerarNFeXML(
		pedido.ID,
		pedido.Cliente,
		pedido.Documento,
		pedido.ProdutosList,
		pedido.ValorTotal,
	)
	if err != nil {
		return model.NotaFiscal{}, err
	}
	
	nota := model.NotaFiscal{
		Arquivo: path,
		Data:    time.Now(),
	}
	
	return nota, nil
}