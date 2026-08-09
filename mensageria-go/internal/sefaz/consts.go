package sefaz

const (
	STATUS_LOTE_RECEBIDO = "103"
	STATUS_PROCESSANDO   = "105"
	STATUS_AUTORIZADA    = "100"
	STATUS_REJEITADA     = "110"
	STATUS_DUPLICIDADE   = "204"
	STATUS_XML_INVALIDO  = "215"
	STATUS_LOTE_SEM_NOTA = "217"
	STATUS_CANCELADA     = "101"
	
	SEFAZ_AMBIENTE_HOMOLOGACAO = "https://homologacao.sefaz.gov.br/nfe/NFeAutorizacao.asmx"
	SEFAZ_AMBIENTE_PRODUCAO    = "https://nfe.sefaz.gov.br/nfe/NFeAutorizacao.asmx"
)