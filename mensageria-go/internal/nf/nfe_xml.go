// internal/nf/nfe_xml.go
package nf

import (
	"fmt"
	"os"
	"strings"
	"time"

	"mensageria-go/internal/model"
)

var empresaConfig = struct {
	Nome string
	CNPJ string
	UF   string
	CodigoMunicipio string
}{
	Nome: "NeoShop Inc",
	CNPJ: "88.115.590/0001-49",
	UF:   "22",
	CodigoMunicipio: "2208403",
}

func GerarNFeXML(pedidoID string, cliente string, documento string, produtos []model.ProdutoItem, valorTotal float64) (string, error) {
	
	if len(produtos) == 0 {
		return "", fmt.Errorf("nenhum produto para gerar nota fiscal")
	}
	
	fmt.Printf("📄 Gerando XML para pedido: %s\n", pedidoID)
	fmt.Printf("📦 Produtos: %d itens\n", len(produtos))
	for i, p := range produtos {
		fmt.Printf("  %d - %s (Qtd: %d, Preço: R$ %.2f)\n", i+1, p.Nome, p.Quantidade, p.Preco)
	}
	
	data := time.Now().Format("2006-01-02T15:04:05-03:00")
	
	docLimpo := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(documento, ".", ""), "-", ""), "/", ""), " ", "")
	
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
	
	itensXML := ""
	var totalProdutos float64 = 0
	
	for i, produto := range produtos {
		quantidade := float64(produto.Quantidade)
		subtotal := produto.Subtotal
		if subtotal == 0 {
			subtotal = produto.Preco * quantidade
		}
		totalProdutos += subtotal
		
		itensXML += fmt.Sprintf(`
    <det nItem="%d">
      <prod>
        <cProd>%03d</cProd>
        <cEAN>SEM GTIN</cEAN>
        <xProd>%s</xProd>
        <NCM>84713000</NCM>
        <CFOP>5102</CFOP>
        <uCom>UN</uCom>
        <qCom>%.4f</qCom>
        <vUnCom>%.2f</vUnCom>
        <vProd>%.2f</vProd>
        <cEANTrib>SEM GTIN</cEANTrib>
        <uTrib>UN</uTrib>
        <qTrib>%.4f</qTrib>
        <vUnTrib>%.2f</vUnTrib>
        <indTot>1</indTot>
      </prod>
    </det>`, 
			i+1, 
			i+1, 
			escapeXML(produto.Nome), 
			quantidade, 
			produto.Preco, 
			subtotal,
			quantidade,
			produto.Preco)
	}
	
	if valorTotal == 0 {
		valorTotal = totalProdutos
	}
	
	xml := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<NFe xmlns="http://www.portalfiscal.inf.br/nfe">
  <infNFe versao="4.00" Id="NFe%s">
    <ide>
      <cUF>%s</cUF>
      <natOp>Venda de Produtos</natOp>
      <mod>55</mod>
      <serie>1</serie>
      <nNF>%s</nNF>
      <dhEmi>%s</dhEmi>
      <tpNF>1</tpNF>
      <idDest>1</idDest>
      <cMunFG>%s</cMunFG>
    </ide>
    <emit>
      <CNPJ>%s</CNPJ>
      <xNome>%s</xNome>
      <xFant>%s</xFant>
      <enderEmit>
        <xLgr>Rua da Tecnologia, 1000</xLgr>
        <nro>1000</nro>
        <xBairro>Centro</xBairro>
        <cMun>%s</cMun>
        <xMun>Teresina</xMun>
        <UF>PI</UF>
        <CEP>64000000</CEP>
      </enderEmit>
    </emit>
    <dest>
      <xNome>%s</xNome>
      <indIEDest>%s</indIEDest>
      %s
      <enderDest>
        <xLgr>Rua do Cliente</xLgr>
        <nro>123</nro>
        <xBairro>Centro</xBairro>
        <cMun>%s</cMun>
        <xMun>Teresina</xMun>
        <UF>PI</UF>
      </enderDest>
    </dest>%s
    <total>
      <ICMSTot>
        <vBC>0.00</vBC>
        <vICMS>0.00</vICMS>
        <vICMSDeson>0.00</vICMSDeson>
        <vFCP>0.00</vFCP>
        <vBCST>0.00</vBCST>
        <vST>0.00</vST>
        <vFCPST>0.00</vFCPST>
        <vFCPSTRet>0.00</vFCPSTRet>
        <vProd>%.2f</vProd>
        <vFrete>0.00</vFrete>
        <vSeg>0.00</vSeg>
        <vDesc>0.00</vDesc>
        <vII>0.00</vII>
        <vIPI>0.00</vIPI>
        <vIPIDevol>0.00</vIPIDevol>
        <vPIS>0.00</vPIS>
        <vCOFINS>0.00</vCOFINS>
        <vOutro>0.00</vOutro>
        <vNF>%.2f</vNF>
      </ICMSTot>
    </total>
    <transp>
      <modFrete>0</modFrete>
    </transp>
  </infNFe>
</NFe>`, 
		pedidoID,
		empresaConfig.UF,
		pedidoID,
		data,
		empresaConfig.CodigoMunicipio,
		empresaConfig.CNPJ,
		empresaConfig.Nome,
		empresaConfig.Nome,
		empresaConfig.CodigoMunicipio,
		escapeXML(cliente),
		indIEDest,
		tagDocumento,
		empresaConfig.CodigoMunicipio,
		itensXML,
		valorTotal,
		valorTotal,
	)
	
	err := os.MkdirAll("notas-fiscais", 0755)
	if err != nil {
		return "", fmt.Errorf("erro ao criar diretório: %v", err)
	}
	
	path := fmt.Sprintf("notas-fiscais/%s.xml", pedidoID)
	err = os.WriteFile(path, []byte(xml), 0644)
	if err != nil {
		return "", fmt.Errorf("erro ao salvar XML: %v", err)
	}
	
	fmt.Printf("✅ Nota fiscal gerada: %s\n", path)
	return path, nil
}

func escapeXML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "'", "&apos;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	return s
}