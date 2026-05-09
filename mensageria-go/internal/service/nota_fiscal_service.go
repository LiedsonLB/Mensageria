// internal/service/nota_fiscal.go
package service

import (
	"time"

	"mensageria-go/internal/model"
	"mensageria-go/internal/nf"
)

func GerarNota(pedido model.Pedido) (model.NotaFiscal, error) {
	
	path, err := nf.GerarNFeXML(
		pedido.ID,
		pedido.Cliente,
		pedido.Documento,
		pedido.Produtos,
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