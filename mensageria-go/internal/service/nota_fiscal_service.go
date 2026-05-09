// internal/service/nota_fiscal.go
package service

import (
	"time"

	"mensageria-go/internal/model"
	"mensageria-go/internal/nf"
)

func GerarNota(pedido model.Pedido) (model.NotaFiscal, error) {

	produtos, _ := pedido.GetProdutosList()
	
	path, err := nf.GerarNFeXML(
		pedido.ID,
		pedido.Cliente,
		pedido.Documento,
		produtos,
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