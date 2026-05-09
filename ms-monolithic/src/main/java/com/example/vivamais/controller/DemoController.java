package com.example.vivamais.controller;

import java.util.Map;

import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.example.vivamais.entity.Pedido;
import com.example.vivamais.entity.dto.PedidoRequest;

@RestController
@RequestMapping("/pedido")
public class DemoController {
    
    @Autowired
    private RabbitTemplate rabbitTemplate;
    
    @Value("${rabbitmq.queue.name}")
    private String queueName;
    
    @PostMapping("/async")
    public ResponseEntity<?> processarAssincrono(@RequestBody PedidoRequest request) {
        
        Pedido pedido = new Pedido(request.getProduto(), request.getValor(), request.getEmail());
        
        rabbitTemplate.convertAndSend(queueName, pedido);
        
        return ResponseEntity.accepted().body(Map.of(
            "mensagem", "✅ Pedido recebido! Você receberá um e-mail em breve.",
            "pedidoId", pedido.getId(),
            "tempoResposta", "50ms"
        ));
    }
}