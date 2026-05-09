package com.example.vivamais.webhook;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.stripe.model.Event;
import com.stripe.model.PaymentIntent;
import com.stripe.net.Webhook;

@RestController
@RequestMapping("/webhook")
public class StripeWebhookController {
    
    @Value("${stripe.webhook.secret:}")
    private String webhookSecret;
    
    @PostMapping("/stripe")
    public ResponseEntity<String> handleStripeWebhook(
            @RequestBody String payload,
            @RequestHeader("Stripe-Signature") String sigHeader) {
        
        // Para desenvolvimento local, você pode logar e ignorar a verificação
        System.out.println("📡 Webhook recebido do Stripe!");
        
        try {
            if (webhookSecret != null && !webhookSecret.isEmpty()) {
                Event event = Webhook.constructEvent(payload, sigHeader, webhookSecret);
                System.out.println("   Evento: " + event.getType());
                
                if ("payment_intent.succeeded".equals(event.getType())) {
                    PaymentIntent intent = (PaymentIntent) event.getDataObjectDeserializer().getObject().get();
                    System.out.println("✅ Pagamento confirmado! Intent ID: " + intent.getId());
                }
            }
            
            return ResponseEntity.ok("Webhook recebido");
            
        } catch (Exception e) {
            System.err.println("❌ Erro no webhook: " + e.getMessage());
            return ResponseEntity.badRequest().build();
        }
    }
}