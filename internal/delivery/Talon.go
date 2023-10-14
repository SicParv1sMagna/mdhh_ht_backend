package delivery

import (
	"context"
	"encoding/json"
	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
	"github.com/SicParv1sMagna/mdhh_backend/internal/repository"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"strconv"
	"time"
)

//func getFailJson(msg string) ([]byte, error) {
//	str := `{"status":  "fail", "message": ` + msg + `}`
//
//	if jsonString, err := json.Marshal(str); err != nil {
//		return []byte{}, err
//	} else {
//		return jsonString, nil
//	}
//}

func pushNotification(talon model.Talon) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(talon)
	failOnError(err, "Failed to marchall body")

	err = ch.PublishWithContext(ctx,
		"logs_direct", // exchange
		strconv.FormatInt(int64(talon.BranchID), 10), // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %d", body)
}

func AddTalon(repository *repository.Repository, c *gin.Context) {
	var talon = model.Talon{}

	var err error
	if err = c.ShouldBindJSON(&talon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err = repository.AddTalon(&talon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	go pushNotification(talon)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"id":     talon.ID,
	})

}

func DeleteTalon(repository *repository.Repository, c *gin.Context) {
	id, err := strconv.Atoi(c.DefaultQuery("id", ""))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if id < 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err = repository.DeleteTalon(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
