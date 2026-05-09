package com.example.vivamais.service;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import com.example.vivamais.entity.Pedido;
import com.stripe.Stripe;
import com.stripe.model.PaymentIntent;
import com.stripe.param.PaymentIntentCreateParams;

@Service
public class StripeService {
    
    @Value("${stripe.api.key}")
    private String stripeApiKey;
    
    /**
     * Cria um PaymentIntent no Stripe para processar o pagamento
     * @param pedido Pedido a ser processado
     * @return ClientSecret para o frontend usar
     */
    public String criarPaymentIntent(Pedido pedido) throws Exception {
        // Configura a chave da API
        Stripe.apiKey = stripeApiKey;
        
        // Converte valor para centavos (Stripe trabalha com a menor unidade da moeda)
        long valorCentavos = (long) (pedido.getValor() * 100);
        
        // Cria os parâmetros do PaymentIntent
        PaymentIntentCreateParams params = PaymentIntentCreateParams.builder()
                .setAmount(valorCentavos)
                .setCurrency("brl")
                .setDescription("Pedido: " + pedido.getId() + " - " + pedido.getProduto())
                .setReceiptEmail(pedido.getEmail())
                .putMetadata("pedido_id", pedido.getId())
                .putMetadata("produto", pedido.getProduto())
                .build();
        
        // Cria o PaymentIntent no Stripe
        PaymentIntent intent = PaymentIntent.create(params);
        
        System.out.println("💳 [STRIPE] PaymentIntent criado: " + intent.getId());
        System.out.println("   Status: " + intent.getStatus());
        System.out.println("   Valor: R$ " + pedido.getValor());
        
        return intent.getClientSecret();
    }
    
    /**
     * Confirma se o pagamento foi bem sucedido
     */
    public boolean confirmarPagamento(String paymentIntentId) throws Exception {
        Stripe.apiKey = stripeApiKey;
        PaymentIntent intent = PaymentIntent.retrieve(paymentIntentId);
        return "succeeded".equals(intent.getStatus());
    }
}