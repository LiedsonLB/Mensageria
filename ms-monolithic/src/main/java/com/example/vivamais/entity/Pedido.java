package com.example.vivamais.entity;

import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.AllArgsConstructor;
import java.io.Serializable;
import java.time.LocalDateTime;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Pedido implements Serializable {
    private static final long serialVersionUID = 1L;
    
    private String id;
    private String produto;
    private Double valor;
    private String email;
    private String status;
    private String nomeUsuario;
    private String codigo;
    private LocalDateTime timestamp;
    
    public Pedido(String produto, Double valor, String email) {
        this.id = java.util.UUID.randomUUID().toString().substring(0, 8).toUpperCase();
        this.produto = produto;
        this.valor = valor;
        this.email = email;
        this.status = "PENDENTE";
        this.timestamp = LocalDateTime.now();
    }
}