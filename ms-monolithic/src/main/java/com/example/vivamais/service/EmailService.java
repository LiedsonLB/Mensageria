package com.example.vivamais.service;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.MimeMessageHelper;
import org.springframework.stereotype.Service;
import org.thymeleaf.TemplateEngine;
import org.thymeleaf.context.Context;

import com.example.vivamais.entity.Pedido;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class EmailService {
    
    private static final Logger logger = LoggerFactory.getLogger(EmailService.class);
    
    private final JavaMailSender mailSender;
    private final TemplateEngine templateEngine;
    
    @Value("${spring.mail.username}")
    private String fromEmail;
    
    public void enviarConfirmacao(Pedido pedido) {
        try {
            logger.info("📧 Preparando e-mail para: {}", pedido.getEmail());
            
            // Prepara o template HTML
            Context context = new Context();
            context.setVariable("pedidoId", pedido.getId());
            context.setVariable("produto", pedido.getProduto());
            context.setVariable("valor", String.format("%.2f", pedido.getValor()));
            context.setVariable("data", pedido.getTimestamp());
            
            String html = templateEngine.process("email-template", context);
            
            // Envia o e-mail
            var message = mailSender.createMimeMessage();
            var helper = new MimeMessageHelper(message, true, "UTF-8");
            
            helper.setFrom(fromEmail);
            helper.setTo(pedido.getEmail());
            helper.setSubject("✅ Confirmação de Pedido #" + pedido.getId());
            helper.setText(html, true);
            
            mailSender.send(message);
            
            logger.info("✅ E-mail enviado com sucesso para: {}", pedido.getEmail());
            
        } catch (Exception e) {
            logger.error("❌ Erro ao enviar e-mail: {}", e.getMessage());
        }
    }
}