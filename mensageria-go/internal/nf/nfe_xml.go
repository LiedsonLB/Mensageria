// internal/nf/nfe_xml.go
package nf

import (
	"fmt"
	"os"
	"strings"
	"time"

	"mensageria-go/internal/model"
)

// Configuração da empresa emissora
var empresaConfig = struct {
	Nome string
	CNPJ string
	UF   string
}{
	Nome: "NeoShop Inc",
	CNPJ: "88.115.590/0001-49",
	UF:   "35", // SP
}

func GerarNFeXML(pedidoID string, cliente string, documento string, produtos []model.ProdutoItem, valorTotal float64) (string, error) {
	
	data := time.Now().Format("2006-01-02T15:04:05-07:00")
	
	// Limpar documento
	docLimpo := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(documento, ".", ""), "-", ""), "/", ""), " ", "")
	
	// Determinar o tipo de documento
	var tagDocumento string
	var indIEDest string
	
	if len(docLimpo) == 11 {
		tagDocumento = fmt.Sprintf("<CPF>%s</CPF>", docLimpo)
		indIEDest = "9"
	} else if len(docLimpo) == 14 {
		tagDocumento = fmt.Sprintf("<CNPJ>%s</CNPJ>", docLimpo)
		indIEDest = "1"
	} else {
		tagDocumento = fmt.Sprintf("<xNome>%s</xNome>", cliente)
		indIEDest = "9"
	}
	
	// Construir os itens da nota fiscal usando o subtotal
	itensXML := ""
	for i, produto := range produtos {
		quantidade := float64(produto.Quantidade)
		// Usar o subtotal que já foi calculado
		subtotal := produto.Subtotal
		if subtotal == 0 {
			// Fallback caso subtotal não tenha sido calculado
			subtotal = produto.Preco * quantidade
		}
		
		itensXML += fmt.Sprintf(`
    <det nItem="%d">
      <prod>
        <cProd>%d</cProd>
        <xProd>%s</xProd>
        <qCom>%.2f</qCom>
        <vUnCom>%.2f</vUnCom>
        <vProd>%.2f</vProd>
      </prod>
    </det>`, i+1, i+1, produto.Nome, quantidade, produto.Preco, subtotal)
	}
	
	// Construir XML completo
	xml := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<NFe xmlns="http://www.portalfiscal.inf.br/nfe">
  <infNFe versao="4.00" Id="NFe%s">
    <ide>
      <cUF>%s</cUF>
      <natOp>Venda</natOp>
      <mod>55</mod>
      <serie>1</serie>
      <nNF>%s</nNF>
      <dhEmi>%s</dhEmi>
      <tpNF>1</tpNF>
    </ide>
    <emit>
      <xNome>%s</xNome>
      <CNPJ>%s</CNPJ>
    </emit>
    <dest>
      <xNome>%s</xNome>
      <indIEDest>%s</indIEDest>
      %s
    </dest>
    %s
    <total>
      <ICMSTot>
        <vProd>%.2f</vProd>
        <vNF>%.2f</vNF>
      </ICMSTot>
    </total>
  </infNFe>
</NFe>`, 
		pedidoID,
		empresaConfig.UF,
		pedidoID,
		data,
		empresaConfig.Nome,
		empresaConfig.CNPJ,
		cliente,
		indIEDest,
		tagDocumento,
		itensXML,
		valorTotal,
		valorTotal,
	)
	
	// Salvar arquivo
	path := fmt.Sprintf("notas-fiscais/%s.xml", pedidoID)
	err := os.WriteFile(path, []byte(xml), 0644)
	if err != nil {
		return "", err
	}
	
	return path, nil
}