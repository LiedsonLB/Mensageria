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

// removeAcentos e normaliza texto para evitar problemas no PDF
func normalizarTexto(s string) string {
	// Mapeamento de caracteres acentuados para seus equivalentes sem acento
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
	return result.String()
}

// formatarChave formata a chave de acesso em grupos de 4
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

	// Adicionar fonte que suporta acentos (usar UTF-8)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 10, 10)
	pdf.AddPage()

	// Usar fonte que suporta caracteres especiais
	pdf.SetFont("Arial", "", 10)

	// Normalizar textos para evitar problemas de acentuação
	clienteNome := normalizarTexto(pedido.Cliente)
	clienteNome = strings.ReplaceAll(clienteNome, "Ã©", "e")
	clienteNome = strings.ReplaceAll(clienteNome, "Ã£", "a")
	clienteNome = strings.ReplaceAll(clienteNome, "Ãµ", "o")
	clienteNome = strings.ReplaceAll(clienteNome, "Ã§", "c")
	clienteNome = strings.ReplaceAll(clienteNome, "Ã¡", "a")
	clienteNome = strings.ReplaceAll(clienteNome, "Ã­", "i")
	clienteNome = strings.ReplaceAll(clienteNome, "Ã³", "o")
	clienteNome = strings.ReplaceAll(clienteNome, "Ãº", "u")

	// Limpar chave de acesso
	chaveAcesso = strings.TrimSpace(chaveAcesso)
	if len(chaveAcesso) < 44 {
		chaveAcesso = chaveAcesso + strings.Repeat("0", 44-len(chaveAcesso))
	}
	if len(chaveAcesso) > 44 {
		chaveAcesso = chaveAcesso[:44]
	}

	// Limpar protocolo
	protocolo = strings.TrimSpace(protocolo)

	// Garantir que produtos estão carregados
	produtosList := pedido.ProdutosList
	if len(produtosList) == 0 {
		produtosList, _ = pedido.GetProdutosList()
	}

	// Calcular subtotais se necessário
	for i := range produtosList {
		if produtosList[i].Subtotal == 0 {
			produtosList[i].Subtotal = produtosList[i].Preco * float64(produtosList[i].Quantidade)
		}
		// Normalizar nome do produto
		produtosList[i].Nome = normalizarTexto(produtosList[i].Nome)
	}

	log.Printf("📄 Gerando DANFE para pedido: %s", pedido.ID)
	log.Printf("   Cliente: %s", clienteNome)
	log.Printf("   Produtos: %d", len(produtosList))

	// ==================== CABEÇALHO COM LOGO (CANTO SUPERIOR ESQUERDO) ====================
	logoPath := "assets/neoshop_logo.png"

	if _, err := os.Stat(logoPath); err == nil {
		// Logo no canto superior esquerdo
		pdf.Image(logoPath, 10, 10, 30, 0, false, "", 0, "")
		pdf.SetY(20)
	} else {
		log.Printf("⚠️ Logo não encontrada em: %s", logoPath)
		pdf.SetY(15)
	}

	// Título NF-e (centralizado)
	pdf.SetFont("Arial", "B", 22)
	pdf.SetTextColor(0, 51, 102)
	pdf.CellFormat(0, 10, "NF-e", "", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.CellFormat(0, 6, "NOTA FISCAL ELETRONICA", "", 1, "C", false, 0, "")

	// Ambiente de homologação
	// pdf.SetFont("Arial", "B", 8)
	// pdf.SetTextColor(255, 0, 0)
	// pdf.SetFillColor(255, 240, 240)
	// pdf.CellFormat(0, 5, ">>> AMBIENTE DE HOMOLOGACAO - SEM VALOR FISCAL <<<", "", 1, "C", true, 0, "")
	// pdf.SetTextColor(0, 0, 0)

	// Linha Canhoto
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
	pdf.SetFont("Arial", "B", 6)
	pdf.SetTextColor(100, 100, 100)
	pdf.CellFormat(0, 4, "DANFE - DOCUMENTO AUXILIAR DA NOTA FISCAL ELETRONICA - CANHOTO", "", 1, "C", false, 0, "")

	// ==================== CHAVE DE ACESSO ====================
	pdf.SetFillColor(240, 240, 245)
	pdf.Rect(10, pdf.GetY(), 190, 22, "F")

	pdf.SetFont("Arial", "B", 7)
	pdf.SetTextColor(0, 51, 102)
	pdf.CellFormat(0, 5, "CHAVE DE ACESSO", "", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(0, 0, 0)
	chaveFormatada := formatarChave(chaveAcesso)
	pdf.CellFormat(0, 5, chaveFormatada, "", 1, "C", false, 0, "")

	// Código de barras visual
	pdf.SetFont("Arial", "", 9)
	barcodeStr := chaveAcesso[:20]
	if len(protocolo) >= 8 {
		barcodeStr += protocolo[:8]
	}
	pdf.CellFormat(0, 5, barcodeStr, "", 1, "C", false, 0, "")
	pdf.Ln(4)

	// ==================== DADOS DO EMITENTE ====================
	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(0, 51, 102)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 6, "EMITENTE", "", 1, "L", true, 0, "")

	pdf.SetFont("Arial", "", 7)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(255, 255, 255)

	pdf.CellFormat(25, 4, "Razao Social:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, "NEOSHOP INC.", "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 4, "CNPJ:", "", 0, "L", false, 0, "")
	pdf.CellFormat(60, 4, "88.115.590/0001-49", "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 4, "IE:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, "123.456.789.123", "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 4, "Endereco:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, " BR-343, Bairro Petecas - Piripiri/PI", "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 4, "E-mail:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, "neoshop@gmail.com", "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 4, "Telefone:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, "(86) 3276-2206", "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 4, "Site:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, "www.neoshop.com.br", "", 1, "L", false, 0, "")

	pdf.Ln(2)

	// ==================== DADOS DO DESTINATÁRIO ====================
	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(0, 51, 102)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 6, "DESTINATARIO", "", 1, "L", true, 0, "")

	pdf.SetFont("Arial", "", 7)
	pdf.SetTextColor(0, 0, 0)

	pdf.CellFormat(25, 4, "Nome:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, clienteNome, "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 4, "CPF/CNPJ:", "", 0, "L", false, 0, "")
	pdf.CellFormat(60, 4, pedido.Documento, "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 4, "IE:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, "ISENTO", "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 4, "E-mail:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, pedido.Email, "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 4, "Telefone:", "", 0, "L", false, 0, "")
	// mecha no cpf para usa-lo como telefone de contato
	telefoneContato := strings.ReplaceAll(pedido.Documento, ".", "")
	telefoneContato = strings.ReplaceAll(telefoneContato, "-", "")
	telefoneContato = telefoneContato[len(telefoneContato)-10:]	
	pdf.CellFormat(0, 4, telefoneContato, "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 4, "Endereco:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, "Rua das Flores, 123 - Bloco B - Apto 5 - Jardim Paulista  - Sao Paulo - SP", "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// ==================== IMPOSTOS ====================
	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(0, 51, 102)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 6, "IMPOSTOS INCIDENTES", "", 1, "L", true, 0, "")

	pdf.SetFont("Arial", "", 7)
	pdf.SetTextColor(0, 0, 0)

	icms := pedido.ValorTotal * 0.18
	pis := pedido.ValorTotal * 0.0165
	cofins := pedido.ValorTotal * 0.076
	totalImpostos := icms + pis + cofins

	pdf.CellFormat(60, 4, "ICMS (18%):", "1", 0, "L", false, 0, "")
	pdf.CellFormat(35, 4, fmt.Sprintf("R$ %.2f", icms), "1", 0, "R", false, 0, "")
	pdf.CellFormat(60, 4, "PIS (1.65%):", "1", 0, "L", false, 0, "")
	pdf.CellFormat(35, 4, fmt.Sprintf("R$ %.2f", pis), "1", 1, "R", false, 0, "")

	pdf.CellFormat(60, 4, "COFINS (7.6%):", "1", 0, "L", false, 0, "")
	pdf.CellFormat(35, 4, fmt.Sprintf("R$ %.2f", cofins), "1", 0, "R", false, 0, "")
	pdf.CellFormat(60, 4, "TOTAL IMPOSTOS:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(35, 4, fmt.Sprintf("R$ %.2f", totalImpostos), "1", 1, "R", false, 0, "")
	pdf.Ln(2)

	// ==================== PRODUTOS ====================
	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(0, 51, 102)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 6, "ITENS DA NOTA FISCAL", "", 1, "L", true, 0, "")

	// Cabeçalho da tabela
	pdf.SetFont("Arial", "B", 6)
	pdf.SetFillColor(200, 200, 200)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(8, 5, "ITEM", "1", 0, "C", true, 0, "")
	pdf.CellFormat(82, 5, "PRODUTO", "1", 0, "C", true, 0, "")
	pdf.CellFormat(15, 5, "QTD", "1", 0, "C", true, 0, "")
	pdf.CellFormat(25, 5, "UN", "1", 0, "C", true, 0, "")
	pdf.CellFormat(28, 5, "VL UNIT", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 5, "VL TOTAL", "1", 1, "C", true, 0, "")

	// Itens da tabela
	pdf.SetFont("Arial", "", 6)
	for i, p := range produtosList {
		nomeProduto := normalizarTexto(p.Nome)
		subtotal := p.Subtotal

		pdf.CellFormat(8, 4, fmt.Sprintf("%d", i+1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(82, 4, nomeProduto, "1", 0, "L", false, 0, "")
		pdf.CellFormat(15, 4, fmt.Sprintf("%d", p.Quantidade), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 4, "PC", "1", 0, "C", false, 0, "")
		pdf.CellFormat(28, 4, fmt.Sprintf("R$ %.2f", p.Preco), "1", 0, "R", false, 0, "")
		pdf.CellFormat(30, 4, fmt.Sprintf("R$ %.2f", subtotal), "1", 1, "R", false, 0, "")
	}

	// Totais
	pdf.SetFont("Arial", "B", 7)
	pdf.CellFormat(158, 5, "VALOR TOTAL DA NOTA:", "1", 0, "R", false, 0, "")
	pdf.CellFormat(30, 5, fmt.Sprintf("R$ %.2f", pedido.ValorTotal), "1", 1, "R", false, 0, "")
	pdf.Ln(2)

	// ==================== INFORMAÇÕES DA NOTA ====================
	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(0, 51, 102)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 6, "INFORMACOES DA NOTA FISCAL", "", 1, "L", true, 0, "")

	pdf.SetFont("Arial", "", 7)
	pdf.SetTextColor(0, 0, 0)

	pdf.CellFormat(25, 4, "Numero:", "", 0, "L", false, 0, "")
	idStr := pedido.ID
	if len(idStr) > 8 {
		idStr = idStr[:8]
	}
	pdf.CellFormat(45, 4, idStr, "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 4, "Serie:", "", 0, "L", false, 0, "")
	pdf.CellFormat(25, 4, "1", "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 4, "NF:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, pedido.ID[:12], "", 1, "L", false, 0, "")

	pdf.CellFormat(25, 4, "Data Emissao:", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 4, time.Now().Format("02/01/2006 15:04:05"), "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 4, "Protocolo:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, protocolo, "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// ==================== PROTOCOLO SEFAZ ====================
	pdf.SetFont("Arial", "B", 7)
	pdf.SetTextColor(0, 139, 0)
	pdf.SetFillColor(220, 255, 220)
	pdf.CellFormat(0, 5, "PROTOCOLO DE AUTORIZACAO DA SEFAZ", "", 1, "L", true, 0, "")
	pdf.SetFont("Arial", "", 6)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 4, "Autorizacao de Uso emitida em "+time.Now().Format("02/01/2006 15:04:05")+" com o protocolo: "+protocolo, "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// ==================== ASSINATURA DIGITAL ====================
	// pdf.SetFont("Arial", "B", 8)
	// pdf.SetTextColor(0, 100, 0)
	// pdf.SetFillColor(220, 255, 220)
	// pdf.CellFormat(0, 6, "PROTOCOLO DE AUTORIZACAO DA SEFAZ", "", 1, "L", true, 0, "")

	// pdf.SetFont("Arial", "", 7)
	// pdf.SetTextColor(0, 0, 0)

	// Dados da autorização
	dataAutorizacaoFormatada := time.Now().Format("02/01/2006 15:04:05")

	// Caixa com informações da autorização
	pdf.SetFillColor(245, 255, 245)
	pdf.Rect(10, pdf.GetY(), 190, 28, "F")

	pdf.CellFormat(40, 4, "Data Autorizacao:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, dataAutorizacaoFormatada, "", 1, "L", false, 0, "")

	pdf.CellFormat(40, 4, "Chave de Acesso:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, chaveAcesso, "", 1, "L", false, 0, "")

	pdf.CellFormat(40, 4, "Protocolo:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, protocolo, "", 1, "L", false, 0, "")

	pdf.CellFormat(40, 4, "Status:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, "100 - AUTORIZADA", "", 1, "L", false, 0, "")

	pdf.CellFormat(40, 4, "Motivo:", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 4, "NFe autorizada pela SEFAZ", "", 1, "L", false, 0, "")

	pdf.Ln(4)

	// ==================== RODAPÉ ====================
	pdf.SetFont("Arial", "I", 6)
	pdf.SetTextColor(150, 150, 150)
	pdf.CellFormat(0, 4, "Documento emitido por Sistema NeoShop", "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 4, "Sistema NeoShop - Mensageria com RabbitMQ", "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 4, fmt.Sprintf("DANFE gerado em %s", time.Now().Format("02/01/2006 15:04:05")), "", 1, "C", false, 0, "")

	// Salvar PDF
	path := filepath.Join(pdfDir, fmt.Sprintf("%s-danfe.pdf", pedido.ID))
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return "", err
	}

	log.Printf("✅ DANFE gerado com sucesso: %s", path)
	return path, nil
}
