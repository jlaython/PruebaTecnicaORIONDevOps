package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type EventPayload struct {
	ID        string    `json:"id" binding:"required"`
	Message   string    `json:"message" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

var (
	cfg  Config
	ch   *amqp.Channel
	conn *amqp.Connection
	rdb  *redis.Client
)

func initRabbitMQ(rabbitUrl, queueName string) (*amqp.Connection, *amqp.Channel, error) {
	c, err := amqp.Dial(rabbitUrl)
	if err != nil {
		return nil, nil, err
	}

	channel, err := c.Channel()
	if err != nil {
		return nil, nil, err
	}

	_, err = channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	return c, channel, nil
}

func initRedis(addr, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	return rdb
}

func main() {
	// Load env
	cfg = LoadConfig()

	var err error

	/* RABBITMQ */
	conn, ch, err = initRabbitMQ(cfg.RabbitURL, cfg.QueueName)
	if err != nil {
		log.Fatalf("Error conectando a RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	/* REDIS */
	rdb = initRedis(cfg.RedisAddr, cfg.RedisPass)
	defer rdb.Close()

	r := gin.Default()

	// Actuator / Health Endpoint
	r.GET("/actuator/health", healthCheck)
	r.GET("/health", healthCheck)

	// API Endpoints
	r.POST("/api/v1/events", publishEvent)
	r.GET("/api/v1/events", getEvents)

	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Servidor Go corriendo en el puerto %s...", cfg.ServerPort)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

func healthCheck(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Check RabbitMQ
	rabbitStatus := "UP"
	if conn == nil || conn.IsClosed() || ch == nil || ch.IsClosed() {
		rabbitStatus = "DOWN"
	}

	// Check Redis via PING
	redisStatus := "UP"
	if err := rdb.Ping(ctx).Err(); err != nil {
		redisStatus = "DOWN"
	}

	globalStatus := "UP"
	httpCode := http.StatusOK

	if rabbitStatus == "DOWN" || redisStatus == "DOWN" {
		globalStatus = "DOWN"
		httpCode = http.StatusServiceUnavailable
	}

	c.JSON(httpCode, gin.H{
		"status": globalStatus,
		"components": gin.H{
			"rabbitmq": gin.H{
				"status": rabbitStatus,
			},
			"redis": gin.H{
				"status": redisStatus,
			},
		},
	})
}

func selectGlobalStatus(rabbitStatus string) string {
	if rabbitStatus == "DOWN" {
		return "DOWN"
	}
	return "UP"
}

func publishEvent(c *gin.Context) {
	var payload EventPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload.CreatedAt = time.Now().UTC()

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error serializando payload"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Save in Redis (Lista)
	if err := rdb.RPush(ctx, cfg.RedisListKey, bodyBytes).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falló al guardar en Redis"})
		return
	}

	// 2. Publish in RabbitMQ
	err = ch.PublishWithContext(
		ctx,
		"",
		cfg.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bodyBytes,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falló al publicar el mensaje en RabbitMQ"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "published & cached",
		"payload": payload,
	})
}

func getEvents(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get all elements
	rawEvents, err := rdb.LRange(ctx, cfg.RedisListKey, 0, -1).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener eventos de Redis"})
		return
	}

	events := make([]EventPayload, 0, len(rawEvents))
	for _, item := range rawEvents {
		var evt EventPayload
		if err := json.Unmarshal([]byte(item), &evt); err == nil {
			events = append(events, evt)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"count":  len(events),
		"events": events,
	})
}
