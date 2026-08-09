package sefaz

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)


func gerarChaveAcesso() string {
	ano := time.Now().Format("06")
	mes := time.Now().Format("01")
	dia := time.Now().Format("02")
	return fmt.Sprintf("3524%s%s%s0000000000000010000000011000000010", ano, mes, dia)
}

func gerarProtocolo() string {
	return fmt.Sprintf("135240%s%08d", time.Now().Format("20060102"), time.Now().UnixNano()%100000000)
}

func gerarHashSimulado(content string) string {
	hash := md5.Sum([]byte(content))
	return hex.EncodeToString(hash[:])
}


func BuildSoapEnvelopeCompleto(xmlNFe string, idLote string, cnpjRemetente string) (string, error) {
	timestamp := time.Now().Format("2006-01-02T15:04:05-03:00")
	
	soapEnvelope := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" 
               xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
               xmlns:xsd="http://www.w3.org/2001/XMLSchema">
  <soap:Header>
    <nfeCabecMsg xmlns="http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao">
      <cUF>22</cUF>
      <versaoDados>4.00</versaoDados>
      <token>TOKEN-SEFAZ-SIMULADO-%s</token>
      <CNPJ>%s</CNPJ>
      <dhRequisicao>%s</dhRequisicao>
    </nfeCabecMsg>
    <wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd">
      <wsse:UsernameToken>
        <wsse:Username>%s</wsse:Username>
        <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">%s</wsse:Password>
        <wsse:Nonce>%s</wsse:Nonce>
        <wsse:Created>%s</wsse:Created>
      </wsse:UsernameToken>
    </wsse:Security>
  </soap:Header>
  <soap:Body>
    <nfeDadosMsg xmlns="http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao">
      <enviNFe xmlns="http://www.portalfiscal.inf.br/nfe" versao="4.00">
        <idLote>%s</idLote>
        <indSinc>0</indSinc>
        %s
      </enviNFe>
    </nfeDadosMsg>
  </soap:Body>
</soap:Envelope>`,
		time.Now().Format("20060102150405"),
		cnpjRemetente,
		timestamp,
		cnpjRemetente,
		gerarHashSimulado(cnpjRemetente)[:16],
		gerarHashSimulado(cnpjRemetente+timestamp)[:20],
		timestamp,
		idLote,
		xmlNFe,
	)
	
	return soapEnvelope, nil
}

func BuildSoapEnvelope(xmlNFe string, idLote string) (string, error) {
	envelope := SoapEnvelope{
		XmlnsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsXS:   "http://www.w3.org/2001/XMLSchema-instance",
		XmlnsXD:   "http://www.w3.org/2001/XMLSchema",
		Body: SoapBody{
			EnviNFe: EnviNFeRequest{
				Versao:  "4.00",
				IdLote:  idLote,
				IndSinc: 0,
				NFeXML:  xmlNFe,
			},
		},
	}
	
	output, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		return "", err
	}
	
	return xml.Header + string(output), nil
}

func GerarXMLAssinado(xmlOriginal string) string {
	digestValue := gerarHashSimulado(xmlOriginal)
	
	assinatura := fmt.Sprintf(`
  <Signature xmlns="http://www.w3.org/2000/09/xmldsig#">
    <SignedInfo>
      <CanonicalizationMethod Algorithm="http://www.w3.org/TR/2001/REC-xml-c14n-20010315"/>
      <SignatureMethod Algorithm="http://www.w3.org/2000/09/xmldsig#rsa-sha1"/>
      <Reference URI="">
        <Transforms>
          <Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature"/>
        </Transforms>
        <DigestMethod Algorithm="http://www.w3.org/2000/09/xmldsig#sha1"/>
        <DigestValue>%s</DigestValue>
      </Reference>
    </SignedInfo>
    <SignatureValue>assinatura-digital-simulada-%d</SignatureValue>
    <KeyInfo>
      <X509Data>
        <X509Certificate>NeoShop-Certificado-SEFAZ-Simulado</X509Certificate>
      </X509Data>
    </KeyInfo>
  </Signature>`, digestValue, time.Now().UnixNano())
	
	pos := strings.LastIndex(xmlOriginal, "</NFe>")
	if pos != -1 {
		return xmlOriginal[:pos] + assinatura + xmlOriginal[pos:]
	}
	
	return xmlOriginal
}


type SefazClient struct {
	Ambiente    string
	Timeout     time.Duration
	CNPJ        string
	Certificado string
}

func NewSefazClient(ambiente string) *SefazClient {
	return &SefazClient{
		Ambiente:    ambiente,
		Timeout:     30 * time.Second,
		CNPJ:        "88115590000149",
		Certificado: "certificado-a1-simulado",
	}
}

func (c *SefazClient) EnviarLote(soapEnvelope string) (*RespostaEnvioLote, error) {
	log.Println("   📡 Enviando lote para SEFAZ via SOAP...")
	log.Println("   🔐 Autenticacao via Token e Certificado Digital")
	time.Sleep(2 * time.Second)
	
	recibo := fmt.Sprintf("%d%08d", time.Now().UnixNano()%1000000, time.Now().UnixNano()%100000000)
	
	return &RespostaEnvioLote{
		Versao:  "4.00",
		Recibo:  recibo,
		CStat:   STATUS_LOTE_RECEBIDO,
		XMotivo: "Lote recebido com sucesso",
	}, nil
}

func (c *SefazClient) ConsultarRecibo(recibo string) (*RespostaConsultaRecibo, error) {
	log.Printf("   🔎 Consultando recibo: %s...", recibo)
	time.Sleep(3 * time.Second)
	
	chaveAcesso := gerarChaveAcesso()
	protocolo := gerarProtocolo()
	
	return &RespostaConsultaRecibo{
		Versao:  "4.00",
		CStat:   STATUS_AUTORIZADA,
		XMotivo: "NFe autorizada",
		ProtNFe: ProtNFe{
			Versao: "4.00",
			InfProt: InfProt{
				ChNFe:    chaveAcesso,
				DhRecbto: time.Now().Format("2006-01-02T15:04:05-03:00"),
				NProt:    protocolo,
				DigVal:   gerarHashSimulado(chaveAcesso),
				CStat:    STATUS_AUTORIZADA,
				XMotivo:  "Autorizado o uso da NF-e",
			},
		},
	}, nil
}


func salvarXMLAutorizado(xmlConteudo string, protocolo string, chaveAcesso string, pedidoID string) (string, error) {
	dir := "notas-fiscais-autorizadas"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("erro ao criar diretório: %v", err)
	}
	
	xmlAutorizado := xmlConteudo
	
	protTag := fmt.Sprintf(`
    <protNFe versao="4.00">
      <infProt>
        <chNFe>%s</chNFe>
        <dhRecbto>%s</dhRecbto>
        <nProt>%s</nProt>
        <digVal>%s</digVal>
        <cStat>100</cStat>
        <xMotivo>Autorizado o uso da NF-e</xMotivo>
      </infProt>
    </protNFe>`, 
		chaveAcesso, 
		time.Now().Format("2006-01-02T15:04:05-03:00"), 
		protocolo, 
		gerarHashSimulado(chaveAcesso))
	
	if strings.Contains(xmlAutorizado, "</infNFe>") {
		xmlAutorizado = strings.Replace(xmlAutorizado, "</infNFe>", protTag+"\n  </infNFe>", 1)
	}
	
	path := filepath.Join(dir, fmt.Sprintf("%s-autorizada.xml", pedidoID))
	err := os.WriteFile(path, []byte(xmlAutorizado), 0644)
	if err != nil {
		return "", fmt.Errorf("erro ao salvar XML autorizado: %v", err)
	}
	
	log.Printf("📄 XML autorizado salvo em: %s", path)
	return path, nil
}


type SefazService struct {
	Client   *SefazClient
	Ambiente string
}

func NewSefazService(ambiente string) *SefazService {
	return &SefazService{
		Client:   NewSefazClient(ambiente),
		Ambiente: ambiente,
	}
}

func (s *SefazService) EnviarNotaFiscal(xmlPath string) (*RespostaSEFAZ, error) {
	log.Println("🏛️ ========== SEFAZ - WEBSERVICE ==========")
	log.Printf("📄 XML: %s", xmlPath)
	log.Printf("🌍 Ambiente: %s", s.Ambiente)
	log.Println("📡 Protocolo: SOAP 1.2 + WS-Security")
	
	xmlData, err := os.ReadFile(xmlPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler XML: %v", err)
	}
	
	log.Println("🔐 Assinando digitalmente a nota fiscal...")
	log.Println("   📜 Certificado A1: NeoShop-Certificado-SEFAZ")
	xmlAssinado := GerarXMLAssinado(string(xmlData))
	
	idLote := fmt.Sprintf("%d", time.Now().UnixNano())
	log.Printf("📦 ID do Lote: %s", idLote)
	
	soapEnvelope, err := BuildSoapEnvelopeCompleto(xmlAssinado, idLote, s.Client.CNPJ)
	if err != nil {
		return nil, fmt.Errorf("erro ao construir SOAP: %v", err)
	}
	
	log.Println("📡 SOAP Envelope construído com Header de segurança")
	log.Println("   🔐 WS-Security: UsernameToken + Nonce + Created")
	log.Println("   🔐 nfeCabecMsg: cUF, versaoDados, token, CNPJ")
	
	soapDir := "soap-envelopes"
	os.MkdirAll(soapDir, 0755)
	soapPath := filepath.Join(soapDir, fmt.Sprintf("%s-soap.xml", idLote))
	os.WriteFile(soapPath, []byte(soapEnvelope), 0644)
	log.Printf("📡 SOAP Envelope salvo: %s", soapPath)
	
	respEnvio, err := s.Client.EnviarLote(soapEnvelope)
	if err != nil {
		return nil, err
	}
	
	log.Printf("📋 Recibo recebido: %s", respEnvio.Recibo)
	
	respConsulta, err := s.Client.ConsultarRecibo(respEnvio.Recibo)
	if err != nil {
		return nil, err
	}
	
	chaveAcesso := respConsulta.ProtNFe.InfProt.ChNFe
	protocolo := respConsulta.ProtNFe.InfProt.NProt
	dataAutorizacao := respConsulta.ProtNFe.InfProt.DhRecbto
	
	pedidoID := filepath.Base(xmlPath)
	pedidoID = strings.TrimSuffix(pedidoID, ".xml")
	
	xmlAutorizadoPath, err := salvarXMLAutorizado(string(xmlData), protocolo, chaveAcesso, pedidoID)
	if err != nil {
		log.Printf("⚠️ Erro ao salvar XML autorizado: %v", err)
	}
	
	resposta := &RespostaSEFAZ{
		Status:           "AUTORIZADA",
		Protocolo:        protocolo,
		ChaveAcesso:      chaveAcesso,
		DataAutorizacao:  dataAutorizacao,
		Mensagem:         respConsulta.XMotivo,
		Recibo:           respEnvio.Recibo,
		NumeroLote:       idLote,
		XmlAutorizado:    xmlAssinado,
		XmlAutorizadoPath: xmlAutorizadoPath,
	}
	
	log.Println("✅ ========== NOTA AUTORIZADA ==========")
	log.Printf("   Protocolo: %s", protocolo)
	log.Printf("   Chave: %s", chaveAcesso)
	log.Printf("   XML Autorizado: %s", xmlAutorizadoPath)
	
	return resposta, nil
}

func EnviarNFe(xmlPath string, pedido interface{}) (RespostaSEFAZ, error) {
	service := NewSefazService("homologacao")
	resp, err := service.EnviarNotaFiscal(xmlPath)
	if err != nil {
		return RespostaSEFAZ{}, err
	}
	return *resp, nil
}