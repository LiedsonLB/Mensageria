package com.example.vivamais.service;

import java.io.File;
import java.io.FileWriter;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

import org.springframework.stereotype.Service;

import com.example.vivamais.entity.Pedido;

@Service
public class NotaFiscalService {
    
    private static final String DIRETORIO_NF = "notas-fiscais";
    
    public NotaFiscalService() {
        // Cria o diretório na inicialização
        File dir = new File(DIRETORIO_NF);
        if (!dir.exists()) {
            dir.mkdirs();
            System.out.println("📁 Diretório criado: " + dir.getAbsolutePath());
        }
    }
    
    /**
     * Gera uma nota fiscal REAL em formato XML
     * @param pedido Pedido para gerar a nota
     * @return Caminho do arquivo gerado
     */
    public String gerarNotaFiscal(Pedido pedido) throws Exception {
        String timestamp = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyyMMddHHmmss"));
        String nomeArquivo = String.format("NF_%s_%s.xml", pedido.getId(), timestamp);
        String caminho = DIRETORIO_NF + File.separator + nomeArquivo;
        
        String xml = gerarXMLNotaFiscal(pedido, timestamp);
        
        try (FileWriter writer = new FileWriter(caminho)) {
            writer.write(xml);
        }
        
        System.out.println("📄 [NF] Nota Fiscal gerada: " + new File(caminho).getAbsolutePath());
        return caminho;
    }
    
    private String gerarXMLNotaFiscal(Pedido pedido, String timestamp) {
        return String.format("""
            <?xml version="1.0" encoding="UTF-8"?>
            <NFe xmlns="http://www.portalfiscal.inf.br/nfe">
                <infNFe versao="4.00" Id="NFe%s">
                    <ide>
                        <cNF>%s</cNF>
                        <nNF>%d</nNF>
                        <serie>1</serie>
                        <dhEmi>%s</dhEmi>
                        <tpNF>1</tpNF>
                        <idDest>1</idDest>
                        <tpEmis>1</tpEmis>
                        <finNFe>1</finNFe>
                        <indFinal>1</indFinal>
                        <indPres>1</indPres>
                        <procEmi>0</procEmi>
                        <verProc>Sistema Mensageria 1.0</verProc>
                    </ide>
                    <emit>
                        <xNome>Sistema Mensageria LTDA</xNome>
                        <CNPJ>12.345.678/0001-90</CNPJ>
                    </emit>
                    <dest>
                        <email>%s</email>
                    </dest>
                    <det nItem="1">
                        <prod>
                            <cProd>%s</cProd>
                            <xProd>%s</xProd>
                            <NCM>99999999</NCM>
                            <CFOP>5102</CFOP>
                            <uCom>UN</uCom>
                            <qCom>1.0000</qCom>
                            <vUnCom>%.2f</vUnCom>
                            <vProd>%.2f</vProd>
                        </prod>
                    </det>
                    <total>
                        <vNF>%.2f</vNF>
                    </total>
                    <transp>
                        <modFrete>0</modFrete>
                    </transp>
                    <pag>
                        <tPag>01</tPag>
                        <vPag>%.2f</vPag>
                    </pag>
                </infNFe>
            </NFe>
            """,
            pedido.getId(),
            pedido.getId().substring(0, Math.min(8, pedido.getId().length())),
            System.currentTimeMillis() % 1000000,
            LocalDateTime.now().format(DateTimeFormatter.ISO_LOCAL_DATE_TIME),
            pedido.getEmail(),
            pedido.getId(),
            pedido.getProduto(),
            pedido.getValor(),
            pedido.getValor(),
            pedido.getValor(),
            pedido.getValor()
        );
    }
}