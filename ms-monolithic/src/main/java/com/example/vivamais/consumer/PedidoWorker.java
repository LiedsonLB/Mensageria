package com.example.vivamais.consumer;

import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import com.example.vivamais.entity.Pedido;
import com.example.vivamais.service.PagamentoService;

@Component
public class PedidoWorker {
    
    @Autowired
    private PagamentoService pagamentoService;  // ← Serviço REAL agora!
    
    @RabbitListener(queues = "${rabbitmq.queue.name}")
    public void processarPedido(Pedido pedido) {
        
        System.out.println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
        System.out.println("📦 [WORKER] Pedido recebido da fila: " + pedido.getId());
        System.out.println("   👤 Cliente: " + pedido.getEmail());
        System.out.println("   📦 Produto: " + pedido.getProduto());
        System.out.println("   💰 Valor: R$ " + pedido.getValor());
        
        try {
            // PROCESSAMENTO REAL!
            pagamentoService.processarPagamentoCompleto(pedido);
            
            System.out.println("✨ [WORKER] Pedido FINALIZADO com sucesso!");
            System.out.println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
            
        } catch (Exception e) {
            System.err.println("❌ [WORKER] Erro ao processar pedido: " + e.getMessage());
            e.printStackTrace();
        }
    }
}