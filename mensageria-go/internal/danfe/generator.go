package danfe

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"mensageria-go/internal/model"

	"github.com/jung-kurt/gofpdf"
)

// removeAcentos remove acentos de strings para evitar erro no PDF
func removeAcentos(s string) string {
	acentos := map[rune]string{
		'á': "a", 'â': "a", 'ã': "a", 'à': "a", 'ä': "a",
		'é': "e", 'ê': "e", 'è': "e", 'ë': "e",
		'í': "i", 'î': "i", 'ì': "i", 'ï': "i",
		'ó': "o", 'ô': "o", 'õ': "o", 'ò': "o", 'ö': "o",
		'ú': "u", 'û': "u", 'ù': "u", 'ü': "u",
		'ç': "c", 'ñ': "n",
		'Á': "A", 'Â': "A", 'Ã': "A", 'À': "A", 'Ä': "A",
		'É': "E", 'Ê': "E", 'È': "E", 'Ë': "E",
		'Í': "I", 'Î': "I", 'Ì': "I", 'Ï': "I",
		'Ó': "O", 'Ô': "O", 'Õ': "O", 'Ò': "O", 'Ö': "O",
		'Ú': "U", 'Û': "U", 'Ù': "U", 'Ü': "U",
		'Ç': "C", 'Ñ': "N",
	}

	var result strings.Builder
	for _, r := range s {
		if val, ok := acentos[r]; ok {
			result.WriteString(val)
		} else {
			result.WriteRune(r)
		}
	}

	texto := result.String()
	texto = strings.ReplaceAll(texto, "á", "a")
	texto = strings.ReplaceAll(texto, "é", "e")
	texto = strings.ReplaceAll(texto, "í", "i")
	texto = strings.ReplaceAll(texto, "ó", "o")
	texto = strings.ReplaceAll(texto, "ú", "u")
	texto = strings.ReplaceAll(texto, "ç", "c")

	return texto
}

// formatarChave formata a chave de acesso em grupos
func formatarChave(chave string) string {
	chaveLimpa := strings.ReplaceAll(chave, " ", "")
	var grupos []string
	for i := 0; i < len(chaveLimpa) && i < 44; i += 4 {
		end := i + 4
		if end > len(chaveLimpa) {
			end = len(chaveLimpa)
		}
		grupos = append(grupos, chaveLimpa[i:end])
	}
	return strings.Join(grupos, " ")
}

// GerarDANFE gera o PDF da nota fiscal
func GerarDANFE(pedido *model.Pedido, protocolo string, chaveAcesso string) (string, error) {
	// Criar pasta
	pdfDir := "notas-fiscais-pdf"
	if err := os.MkdirAll(pdfDir, 0755); err != nil {
		return "", fmt.Errorf("erro ao criar pasta: %v", err)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 10, 10)
	pdf.AddPage()

	// Remover acentos do nome do cliente
	clienteNome := removeAcentos(pedido.Cliente)

	// ==================== CABEÇALHO COM LOGO ====================
	// Verificar se o arquivo da logo existe
	logoPath := "assets/neoshop_logo.png"
	
	if _, err := os.Stat(logoPath); err == nil {
		// Logo no canto superior esquerdo
		pdf.Image(logoPath, 10, 10, 30, 0, false, "", 0, "")
	} else {
		log.Printf("⚠️ Logo não encontrada em: %s", logoPath)
	}

	// if _, err := os.Stat(logoPath); err == nil {
	// 	logoLargura := 40.0
	// 	posX := (210 - logoLargura) / 2 // Centraliza horizontalmente
	// 	pdf.Image(logoPath, posX, 10, logoLargura, 0, false, "", 0, "")
	// } else {
	// 	log.Printf("⚠️ Logo não encontrada em: %s", logoPath)
	// }

	// Título centralizado (ajustado para não sobrepor a logo)
	pdf.SetY(15)
	pdf.SetFont("Arial", "B", 20)
	pdf.SetTextColor(0, 51, 102)
	pdf.CellFormat(0, 12, "NF-e", "", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 6, "NOTA FISCAL ELETRONICA", "", 1, "C", false, 0, "")

	// Linha Superior (Canhoto) - ajustada após o cabeçalho
	pdf.SetY(40)
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(10, 40, 200, 40)
	pdf.SetFont("Arial", "B", 7)
	pdf.SetTextColor(100, 100, 100)
	pdf.CellFormat(0, 4, "DANFE - DOCUMENTO AUXILIAR DA NOTA FISCAL ELETRONICA - CANHOTO", "", 1, "C", false, 0, "")

	// ==================== BARRA SUPERIOR ====================
	pdf.SetFillColor(240, 240, 240)
	pdf.Rect(10, pdf.GetY(), 190, 22, "F")

	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(0, 5, "CHAVE DE ACESSO", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(0, 51, 102)
	chaveFormatada := formatarChave(chaveAcesso)
	pdf.CellFormat(0, 6, chaveFormatada, "", 1, "C", false, 0, "")

	// Código de barras visual
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(0, 0, 0)
	barcodeStr := ""
	if len(chaveAcesso) >= 20 && len(protocolo) >= 8 {
		barcodeStr = chaveAcesso[:20] + protocolo[:8]
	} else {
		barcodeStr = chaveAcesso + protocolo
	}
	pdf.CellFormat(0, 6, barcodeStr, "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// ==================== DADOS DO EMITENTE ====================
	pdf.SetFont("Arial", "B", 9)
	pdf.SetFillColor(0, 51, 102)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 7, "EMITENTE", "", 1, "L", true, 0, "")

	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(255, 255, 255)
	pdf.CellFormat(25, 5, "Razao Social:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 5, "NEOSHOP INC.", "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 5, "CNPJ:", "", 0, "L", false, 0, "")
	pdf.CellFormat(60, 5, "88.115.590/0001-49", "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 5, "IE:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 5, "123.456.789.123", "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 5, "Endereco:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 5, "Rua Joao Cabral, N-2231, Bairro Piraja - Teresina/PI - CEP: 64260-000", "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 5, "Fone:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 5, "(86) 3213-7441", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// ==================== DADOS DO DESTINATÁRIO ====================
	pdf.SetFont("Arial", "B", 9)
	pdf.SetFillColor(0, 51, 102)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 7, "DESTINATARIO / REMETENTE", "", 1, "L", true, 0, "")

	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(25, 5, "Nome:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 5, clienteNome, "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 5, "CNPJ/CPF:", "", 0, "L", false, 0, "")
	pdf.CellFormat(60, 5, pedido.Documento, "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 5, "IE:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 5, "ISENTO", "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 5, "Endereco:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 5, "Rua das Flores, 123 - Apto 42 - Jardim America - Sao Paulo/SP - CEP: 01234-000", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// ==================== IMPOSTOS ====================
	pdf.SetFont("Arial", "B", 9)
	pdf.SetFillColor(0, 51, 102)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 7, "IMPOSTOS INCIDENTES", "", 1, "L", true, 0, "")

	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(0, 0, 0)

	// Calcular impostos
	icms := pedido.ValorTotal * 0.18
	pis := pedido.ValorTotal * 0.0165
	cofins := pedido.ValorTotal * 0.076
	totalImpostos := icms + pis + cofins

	// Tabela de impostos 2x2
	pdf.CellFormat(60, 5, "ICMS (18%):", "1", 0, "L", false, 0, "")
	pdf.CellFormat(35, 5, fmt.Sprintf("R$ %.2f", icms), "1", 0, "R", false, 0, "")
	pdf.CellFormat(60, 5, "PIS (1.65%):", "1", 0, "L", false, 0, "")
	pdf.CellFormat(35, 5, fmt.Sprintf("R$ %.2f", pis), "1", 1, "R", false, 0, "")

	pdf.CellFormat(60, 5, "COFINS (7.6%):", "1", 0, "L", false, 0, "")
	pdf.CellFormat(35, 5, fmt.Sprintf("R$ %.2f", cofins), "1", 0, "R", false, 0, "")
	pdf.CellFormat(60, 5, "TOTAL IMPOSTOS:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(35, 5, fmt.Sprintf("R$ %.2f", totalImpostos), "1", 1, "R", false, 0, "")
	pdf.Ln(2)

	// ==================== PRODUTOS ====================
	pdf.SetFont("Arial", "B", 9)
	pdf.SetFillColor(0, 51, 102)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 7, "ITENS DA NOTA FISCAL", "", 1, "L", true, 0, "")

	// Cabeçalho da tabela
	pdf.SetFont("Arial", "B", 7)
	pdf.SetFillColor(200, 200, 200)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(8, 6, "ITEM", "1", 0, "C", true, 0, "")
	pdf.CellFormat(82, 6, "PRODUTO", "1", 0, "C", true, 0, "")
	pdf.CellFormat(15, 6, "QTD", "1", 0, "C", true, 0, "")
	pdf.CellFormat(25, 6, "UN", "1", 0, "C", true, 0, "")
	pdf.CellFormat(28, 6, "VL UNIT", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 6, "VL TOTAL", "1", 1, "C", true, 0, "")

	// Itens da tabela
	pdf.SetFont("Arial", "", 7)
	for i, p := range pedido.ProdutosList {
		nomeProduto := removeAcentos(p.Nome)

		pdf.CellFormat(8, 5, fmt.Sprintf("%d", i+1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(82, 5, nomeProduto, "1", 0, "L", false, 0, "")
		pdf.CellFormat(15, 5, fmt.Sprintf("%d", p.Quantidade), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 5, "PC", "1", 0, "C", false, 0, "")
		pdf.CellFormat(28, 5, fmt.Sprintf("R$ %.2f", p.Preco), "1", 0, "R", false, 0, "")
		pdf.CellFormat(30, 5, fmt.Sprintf("R$ %.2f", p.Subtotal), "1", 1, "R", false, 0, "")
	}

	// Totais
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(158, 6, "VALOR TOTAL DA NOTA:", "1", 0, "R", false, 0, "")
	pdf.CellFormat(30, 6, fmt.Sprintf("R$ %.2f", pedido.ValorTotal), "1", 1, "R", false, 0, "")
	pdf.Ln(2)

	// ==================== INFORMAÇÕES DA NOTA ====================
	pdf.SetFont("Arial", "B", 9)
	pdf.SetFillColor(0, 51, 102)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 7, "INFORMACOES DA NOTA FISCAL", "", 1, "L", true, 0, "")

	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(0, 0, 0)

	pdf.CellFormat(25, 5, "Numero:", "", 0, "L", false, 0, "")
	idStr := pedido.ID
	if len(idStr) > 8 {
		idStr = idStr[:8]
	}
	pdf.CellFormat(45, 5, idStr, "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 5, "Serie:", "", 0, "L", false, 0, "")
	pdf.CellFormat(25, 5, "1", "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 5, "NF:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 5, pedido.ID[:12], "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 5, "Data Emissao:", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, time.Now().Format("02/01/2006"), "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 5, "Data Saida:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 5, time.Now().Format("02/01/2006"), "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 5, "Protocolo:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 5, protocolo, "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// ==================== PROTOCOLO SEFAZ ====================
	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(0, 139, 0)
	pdf.SetFillColor(220, 255, 220)
	pdf.CellFormat(0, 6, "PROTOCOLO DE AUTORIZACAO DA SEFAZ", "", 1, "L", true, 0, "")

	pdf.SetFont("Arial", "", 7)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 4, "Autorizacao de Uso emitida em "+time.Now().Format("02/01/2006 15:04:05")+" com o protocolo: "+protocolo, "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// ==================== ASSINATURA DIGITAL ====================
	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(220, 220, 220)
	pdf.CellFormat(0, 5, "ASSINATURA DIGITAL", "", 1, "L", true, 0, "")

	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(0, 4, "Assinado digitalmente por: NeoShop Inc. - CNPJ: 88.115.590/0001-49", "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 4, "Autorizacao emitida pela SEFAZ em ambiente de homologacao", "", 1, "L", false, 0, "")

	// Linha de assinatura
	pdf.SetDrawColor(0, 0, 0)
	pdf.Line(10, pdf.GetY()+2, 80, pdf.GetY()+2)
	pdf.SetFont("Arial", "I", 6)
	pdf.CellFormat(0, 5, "Assinatura do Emitente", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// ==================== RODAPÉ ====================
	pdf.SetFont("Arial", "I", 7)
	pdf.SetTextColor(150, 150, 150)
	pdf.CellFormat(0, 4, "Sistema NeoShop - Mensageria com RabbitMQ", "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 4, "DANFE gerado em "+time.Now().Format("02/01/2006 15:04:05"), "", 1, "C", false, 0, "")

	// Salvar PDF
	path := filepath.Join(pdfDir, fmt.Sprintf("%s-danfe.pdf", pedido.ID))
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return "", err
	}

	log.Printf("✅ DANFE profissional gerado: %s", path)
	return path, nil
}