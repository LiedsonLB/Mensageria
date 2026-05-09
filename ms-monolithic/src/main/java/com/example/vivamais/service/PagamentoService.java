package com.example.vivamais.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.example.vivamais.entity.Pedido;

@Service
public class PagamentoService {
    
    @Autowired
    private StripeService stripeService;
    
    @Autowired
    private NotaFiscalService notaFiscalService;
    
    @Autowired
    private EmailService emailService;
    
    /**
     * Processa o pagamento COMPLETO (Stripe + Nota Fiscal + Email)
     * Este método é chamado pelo Worker
     */
    public void processarPagamentoCompleto(Pedido pedido) throws Exception {
        System.out.println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
        System.out.println("💰 [PAGAMENTO] Iniciando processamento do pedido: " + pedido.getId());
        
        // 1. CRIA PAYMENTINTENT NO STRIPE
        System.out.println("💳 [1/3] Criando PaymentIntent no Stripe...");
        String clientSecret = stripeService.criarPaymentIntent(pedido);
        System.out.println("   ✅ PaymentIntent criado! ClientSecret: " + clientSecret.substring(0, 20) + "...");
        
        // NOTA: Em produção, o frontend usaria esse clientSecret para confirmar o pagamento
        // Para demonstração, vamos simular que o pagamento foi confirmado
        
        // 2. GERA NOTA FISCAL
        System.out.println("📄 [2/3] Gerando nota fiscal...");
        String caminhoNF = notaFiscalService.gerarNotaFiscal(pedido);
        System.out.println("   ✅ Nota fiscal salva em: " + caminhoNF);
        
        // 3. ENVIA E-MAIL DE CONFIRMAÇÃO
        System.out.println("📧 [3/3] Enviando e-mail de confirmação...");
        emailService.enviarConfirmacao(pedido);
        
        // Atualiza status do pedido
        pedido.setStatus("PAGO");
        
        System.out.println("✨ [PAGAMENTO] Pedido " + pedido.getId() + " processado com SUCESSO!");
        System.out.println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
    }
}