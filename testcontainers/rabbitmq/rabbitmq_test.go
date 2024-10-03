package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
	"go.uber.org/zap"
	"os"
	"testing"
)

const (
	defaultUser     = "guest"
	defaultPassword = "guest"
)

var (
	logger         *zap.Logger
	amqpConnection *amqp.Connection
	amqpChannel    *amqp.Channel
)

type Message struct {
	Url        string `json:"blob_url"`
	Channel    uint8  `json:"channel"`
	BurstId    uint16 `json:"burst_id"`
	DatatakeId uint32 `json:"datatake_id"`
}

func TestMain(m *testing.M) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, _ = zap.Config.Build(config)
	ctx := context.Background()
	rabbitmqContainer, err := rabbitmq.Run(ctx,
		"rabbitmq:3.12.11-management-alpine",
		rabbitmq.WithAdminUsername(defaultUser),
		rabbitmq.WithAdminPassword(defaultPassword),
	)
	if err != nil {
		logger.Fatal("failed to start container: %s", zap.Error(err))
	}
	rabbitmqURL, err := rabbitmqContainer.AmqpURL(ctx)
	if err != nil {
		logger.Fatal("Failed to get AmqpURL from container", zap.Error(err))
	}
	amqpConnection, err = amqp.Dial(rabbitmqURL)
	if err != nil {
		logger.Fatal("Failed to connect to Rabbitmq", zap.Error(err))
	}
	amqpChannel, err = amqpConnection.Channel()
	if err != nil {
		logger.Fatal("Failed to open a Rabbitmq channel", zap.Error(err))
	}

	exitVal := m.Run()

	_ = amqpChannel.Close()
	_ = amqpConnection.Close()
	if err = rabbitmqContainer.Terminate(ctx); err != nil {
		logger.Fatal("Failed to terminate container", zap.Error(err))
	}
	os.Exit(exitVal)
}

func TestSimplePubSub(t *testing.T) {
	queueName := "data1"
	message := &Message{
		Url:        "chunk_1_datatakeId_1-burstId_1-channelId_1.bin",
		Channel:    1,
		BurstId:    1,
		DatatakeId: 1,
	}

	// Create and check queue
	queue, err := amqpChannel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		t.Error("Failed to declare queue", zap.Error(err))
	}
	if queue.Messages != 0 {
		t.Errorf("Queue: '%s' was not empty", queueName)
	}

	// Pub
	messageBytes, err := json.Marshal(message)
	if err != nil {
		t.Error("Failed to marshal to json", zap.Error(err))
	}
	err = amqpChannel.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        messageBytes,
	})
	if err != nil {
		t.Error("Failed to publish to queue", zap.Error(err))
	}
	queInspect, err := amqpChannel.QueueInspect(queueName)
	if err != nil {
		t.Error("Failed to inspect queue", zap.Error(err))
	}
	if queInspect.Messages != 1 {
		t.Errorf("Number of messages in '%s' expected to be '1' but was '%d'", queueName, queInspect.Messages)
	}

	// Sub
	msg, _, err := amqpChannel.Get(queueName, true)
	if err != nil {
		t.Error("Error when get message", zap.Error(err))
	}
	subMessage := new(Message)
	err = json.Unmarshal(msg.Body, subMessage)
	if err != nil {
		t.Error("Error reading json", zap.Error(err))
	}
	result := *subMessage
	expected := Message{
		Url:        "chunk_1_datatakeId_1-burstId_1-channelId_1.bin",
		Channel:    1,
		BurstId:    1,
		DatatakeId: 1,
	}
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, expected: %v.", result, expected)
	}

	// Inspect queue after
	queInspect, err = amqpChannel.QueueInspect(queueName)
	if err != nil {
		t.Error("Failed to inspect queue", zap.Error(err))
	}
	if queInspect.Messages != 0 {
		t.Errorf("Number of messages in '%s' expected to be '1' but was '%d'", queueName, queInspect.Messages)
	}
}

func TestDuplicateMessages(t *testing.T) {
	queueName := "data2"
	message := &Message{
		Url:        "chunk_1_datatakeId_1-burstId_1-channelId_1.bin",
		Channel:    1,
		BurstId:    1,
		DatatakeId: 1,
	}

	// Create and check queue
	queue, err := amqpChannel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		t.Error("Failed to declare queue", zap.Error(err))
	}
	if queue.Messages != 0 {
		t.Errorf("Queue: '%s' was not empty", queueName)
	}

	// Pub
	messageBytes, err := json.Marshal(message)
	if err != nil {
		t.Error("Failed to marshal to json", zap.Error(err))
	}

	for i := 0; i < 2; i++ {
		err = amqpChannel.Publish("", queueName, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        messageBytes,
		})
		if err != nil {
			t.Error("Failed to publish to queue", zap.Error(err))
		}
	}
	queInspect, err := amqpChannel.QueueInspect(queueName)
	if err != nil {
		t.Error("Failed to inspect queue", zap.Error(err))
	}
	if queInspect.Messages != 1 {
		t.Errorf("Number of messages in '%s' expected to be '1' but was '%d'", queueName, queInspect.Messages)
	}
}
