// internal/service/pagamento.go
package service

import "mensageria-go/internal/model"

func ProcessarPagamento(
	pedido model.Pedido,
	stripeService *StripeService,
) (model.NotaFiscal, error) {
	
	err := stripeService.Pagar(pedido.ValorTotal)
	if err != nil {
		return model.NotaFiscal{}, err
	}
	
	return GerarNota(pedido)
}