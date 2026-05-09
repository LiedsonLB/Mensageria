package queue

import (
	"encoding/json"
	"log"

	amqp "github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	url  string
}

/*
====================================================
CONNECTION
====================================================
*/

func NewRabbitMQ(url string) *RabbitMQ {

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal("❌ Erro ao conectar RabbitMQ:", err)
	}

	// ← CORREÇÃO: Criar o channel principal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("❌ Erro ao criar channel:", err)
	}

	log.Println("✅ Conectado ao RabbitMQ")

	return &RabbitMQ{
		conn: conn,
		ch:   ch, // ← Inicializar o channel
		url:  url,
	}
}

// cria channel isolado (REGRA DE OURO)
func (r *RabbitMQ) newChannel() (*amqp.Channel, error) {
	ch, err := r.conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

/*
====================================================
BASIC QUEUE
====================================================
*/

func (r *RabbitMQ) Publish(queue string, body []byte) error {
	// ← CORREÇÃO: Verificar se o channel existe
	if r.ch == nil {
		log.Println("❌ Channel não inicializado")
		return nil
	}
	
	// Declarar fila
	_, err := r.ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println("Erro ao declarar fila:", err)
		return err
	}

	// Publicar mensagem
	err = r.ch.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	
	if err != nil {
		log.Println("Erro ao publicar mensagem:", err)
		return err
	}
	
	return nil
}

func (r *RabbitMQ) Consume(queue string) (<-chan amqp.Delivery, error) {

	ch, err := r.newChannel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}

/*
====================================================
QUEUE STATS
====================================================
*/

func (r *RabbitMQ) GetQueueStats(queueName string) (map[string]interface{}, error) {

	ch, err := r.newChannel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	stats := map[string]interface{}{
		"name":      q.Name,
		"messages":  q.Messages,
		"consumers": q.Consumers,
		"status":    "active",
	}

	return stats, nil
}

/*
====================================================
PUB / SUB SETUP
====================================================
*/

func (r *RabbitMQ) SetupPubSub() error {

	ch, err := r.newChannel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Exchange principal
	err = ch.ExchangeDeclare(
		"pedidos_exchange",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Dead Letter Exchange
	err = ch.ExchangeDeclare(
		"dlx_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Println("✅ Pub/Sub configurado")

	return nil
}

/*
====================================================
PUBLISH PEDIDO (PUBSUB)
====================================================
*/

func (r *RabbitMQ) PublicarPedido(pedido interface{}) error {

	ch, err := r.newChannel()
	if err != nil {
		return err
	}
	defer ch.Close()

	body, err := json.Marshal(pedido)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"pedidos_exchange",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)

	if err != nil {
		log.Println("❌ Erro publicar pedido:", err)
		return err
	}

	log.Println("📤 Pedido publicado")

	return nil
}

/*
====================================================
FILA PAGAMENTOS + DLQ
====================================================
*/

func (r *RabbitMQ) CriarFilaPagamento() error {

	ch, err := r.newChannel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"pagamentos",
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    "dlx_exchange",
			"x-dead-letter-routing-key": "pagamentos_dlq",
		},
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		"pagamentos",
		"",
		"pedidos_exchange",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		"pagamentos_dlq",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		"pagamentos_dlq",
		"pagamentos_dlq",
		"dlx_exchange",
		false,
		nil,
	)

	log.Println("✅ Fila pagamentos + DLQ criada")

	return err
}

/*
====================================================
FILA EMAIL
====================================================
*/

func (r *RabbitMQ) CriarFilaEmail() error {

	ch, err := r.newChannel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"emails",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		"emails",
		"",
		"pedidos_exchange",
		false,
		nil,
	)

	log.Println("✅ Fila emails criada")

	return err
}

/*
====================================================
FILA NOTA FISCAL (CORRIGIDO)
====================================================
*/

func (r *RabbitMQ) CriarFilaNotaFiscal() error {
	// ← CORREÇÃO: Usar newChannel() em vez de r.ch diretamente
	ch, err := r.newChannel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"notas_fiscais",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		"notas_fiscais",
		"",
		"pedidos_exchange",
		false,
		nil,
	)

	log.Println("✅ Fila de notas fiscais configurada")
	return err
}

/*
====================================================
CONSUME COM RETRY (ACK MANUAL)
====================================================
*/

func (r *RabbitMQ) ConsumeWithRetry(queue string) (<-chan amqp.Delivery, error) {

	ch, err := r.newChannel()
	if err != nil {
		return nil, err
	}

	// Prefetch = 1 (worker profissional)
	err = ch.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		return nil, err
	}

	return ch.Consume(
		queue,
		"",
		false, // manual ack
		false,
		false,
		false,
		nil,
	)
}

/*
====================================================
PUBLICAR EMAIL
====================================================
*/

func (r *RabbitMQ) PublicarEmail(emailData interface{}) error {

	ch, err := r.newChannel()
	if err != nil {
		return err
	}
	defer ch.Close()

	body, err := json.Marshal(emailData)
	if err != nil {
		return err
	}

	return ch.Publish(
		"pedidos_exchange",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)
}

/*
====================================================
CLOSE
====================================================
*/

func (r *RabbitMQ) Close() {
	if r.ch != nil {
		r.ch.Close()
	}
	if r.conn != nil {
		r.conn.Close()
		log.Println("🔌 RabbitMQ conexão encerrada")
	}
}