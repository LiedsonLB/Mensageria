package com.example.vivamais.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Produto {

    private String codigo;
    private String nome;
    private String descricao;
    private String ncm;
    private String cfop;
    private String unidade;
    private Double quantidade;
    private Double valorUnitario;
}