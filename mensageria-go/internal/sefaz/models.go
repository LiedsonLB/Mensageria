package sefaz

import (
	"encoding/xml"
)

// ==================== SOAP ENVELOPE (SEM HEADER) ====================

type SoapEnvelope struct {
	XMLName   xml.Name `xml:"soap:Envelope"`
	XmlnsSoap string   `xml:"xmlns:soap,attr"`
	XmlnsXS   string   `xml:"xmlns:xsi,attr"`
	XmlnsXD   string   `xml:"xmlns:xsd,attr"`
	Body      SoapBody `xml:"soap:Body"`
}

type SoapBody struct {
	EnviNFe     EnviNFeRequest     `xml:"enviNFe"`
	ConsReciNFe ConsReciNFeRequest `xml:"consReciNFe"`
}

// ==================== CABEÇALHO SOAP (AUTENTICAÇÃO) ====================

type SoapHeader struct {
	XMLName    xml.Name   `xml:"soap:Header"`
	NfeCabecMsg NfeCabecMsg `xml:"nfeCabecMsg"`
	Security   Security    `xml:"Security"`
}

type NfeCabecMsg struct {
	XMLName     xml.Name `xml:"nfeCabecMsg"`
	Xmlns       string   `xml:"xmlns,attr"`
	CUF         string   `xml:"cUF"`
	VersaoDados string   `xml:"versaoDados"`
	Token       string   `xml:"token"`
	CNPJ        string   `xml:"CNPJ"`
	DhRequisicao string  `xml:"dhRequisicao"`
}

type Security struct {
	XMLName       xml.Name      `xml:"Security"`
	XmlnsWSSe     string        `xml:"xmlns:wsse,attr"`
	UsernameToken UsernameToken `xml:"UsernameToken"`
}

type UsernameToken struct {
	Username string `xml:"Username"`
	Password string `xml:"Password"`
	Nonce    string `xml:"Nonce"`
	Created  string `xml:"Created"`
}

// ==================== ENVIO DE LOTE ====================

type EnviNFeRequest struct {
	Versao  string `xml:"versao,attr"`
	IdLote  string `xml:"idLote"`
	IndSinc int    `xml:"indSinc"`
	NFeXML  string `xml:"NFe"`
}

type ConsReciNFeRequest struct {
	Versao string `xml:"versao,attr"`
	Recibo string `xml:"tpAmb"`
}

// ==================== RESPOSTAS DA SEFAZ ====================

type RespostaEnvioLote struct {
	XMLName xml.Name `xml:"enviNFeResponse"`
	Versao  string   `xml:"versao,attr"`
	Recibo  string   `xml:"recibo"`
	CStat   string   `xml:"cStat"`
	XMotivo string   `xml:"xMotivo"`
}

type RespostaConsultaRecibo struct {
	XMLName xml.Name `xml:"consReciNFeResponse"`
	Versao  string   `xml:"versao,attr"`
	CStat   string   `xml:"cStat"`
	XMotivo string   `xml:"xMotivo"`
	ProtNFe ProtNFe  `xml:"protNFe"`
}

type ProtNFe struct {
	Versao  string  `xml:"versao,attr"`
	InfProt InfProt `xml:"infProt"`
}

type InfProt struct {
	ChNFe    string `xml:"chNFe"`
	DhRecbto string `xml:"dhRecbto"`
	NProt    string `xml:"nProt"`
	DigVal   string `xml:"digVal"`
	CStat    string `xml:"cStat"`
	XMotivo  string `xml:"xMotivo"`
}

// ==================== DADOS DA NOTA AUTORIZADA ====================

type RespostaSEFAZ struct {
	Status            string
	Protocolo         string
	ChaveAcesso       string
	DataAutorizacao   string
	Mensagem          string
	Recibo            string
	XmlAutorizado     string
	NumeroLote        string
	XmlAutorizadoPath string
}