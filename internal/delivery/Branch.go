package delivery

import (
	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
	"github.com/SicParv1sMagna/mdhh_backend/internal/pkg/middleware/decode"
	"github.com/SicParv1sMagna/mdhh_backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func GetAllBranches(repository *repository.Repository, c *gin.Context) {
	branches, err := repository.GetAllBranches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var branchResponses []model.BranchResponse

	for _, branch := range branches {
		openHours, err := decode.UnmarshalOpenHours(branch.OpenHours)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		openHoursIndividual, err := decode.UnmarshalOpenHours(branch.OpenHoursIndividual)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		response := model.BranchResponse{
			Branch_ID:           branch.Branch_ID,
			SalePointName:       branch.SalePointName,
			Address:             branch.Address,
			Status:              branch.Status,
			OpenHours:           openHours,
			RKO:                 branch.RKO,
			OpenHoursIndividual: openHoursIndividual,
			OfficeType:          branch.OfficeType,
			SalePointFormat:     branch.SalePointFormat,
			SUOAvailability:     branch.SUOAvailability,
			HasRamp:             branch.HasRamp,
			Latitude:            branch.Latitude,
			Longitude:           branch.Longitude,
			MetroStation:        branch.MetroStation,
			Distance:            branch.Distance,
			KEP:                 branch.KEP,
			MyBranch:            branch.MyBranch,
			Network:             branch.Network,
			SalePointCode:       branch.SalePointCode,
		}

		branchResponses = append(branchResponses, response)
	}

	c.JSON(http.StatusOK, branchResponses)
}

func GetBranchBySearch(repository *repository.Repository, c *gin.Context) {
	var branches []model.Branch

	search := c.Param("query")

	branches, err := repository.GetBranchBySearch(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	var branchResponses []model.BranchResponse

	for _, branch := range branches {
		openHours, err := decode.UnmarshalOpenHours(branch.OpenHours)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		openHoursIndividual, err := decode.UnmarshalOpenHours(branch.OpenHoursIndividual)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		response := model.BranchResponse{
			Branch_ID:           branch.Branch_ID,
			SalePointName:       branch.SalePointName,
			Address:             branch.Address,
			Status:              branch.Status,
			OpenHours:           openHours,
			RKO:                 branch.RKO,
			OpenHoursIndividual: openHoursIndividual,
			OfficeType:          branch.OfficeType,
			SalePointFormat:     branch.SalePointFormat,
			SUOAvailability:     branch.SUOAvailability,
			HasRamp:             branch.HasRamp,
			Latitude:            branch.Latitude,
			Longitude:           branch.Longitude,
			MetroStation:        branch.MetroStation,
			Distance:            branch.Distance,
			KEP:                 branch.KEP,
			MyBranch:            branch.MyBranch,
			Network:             branch.Network,
			SalePointCode:       branch.SalePointCode,
		}

		branchResponses = append(branchResponses, response)
	}

	c.JSON(http.StatusOK, branchResponses)
}

func GetBranchById(repository *repository.Repository, c *gin.Context) {

}

func GetNearestBranches(latitude string, longitude string) ([]model.Branch, error) {
	return []model.Branch{
		model.Branch{Branch_ID: 1},
	}, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func GetBranchesWithTalons(repository *repository.Repository, c *gin.Context) {
	latitude := c.DefaultQuery("latitude", "")
	longitude := c.DefaultQuery("longitude", "")

	branches, err := GetNearestBranches(latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	// TODO GetTalons

	// to run rabbit docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management
	ampqConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer ampqConn.Close()

	ch, err := ampqConn.Channel()
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

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for _, br := range branches {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_direct", br)

		err = ch.QueueBind(
			q.Name, // queue name
			strconv.FormatInt(int64(br.Branch_ID), 10), // routing key
			"logs_direct", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for msg := range msgs {
		err = conn.WriteMessage(websocket.TextMessage, msg.Body)
		if err != nil {
			log.Println("Error: ", msg)
		}
		return
	}

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	//c.JSON(http.StatusOK, gin.H{
	//	"status": "success",
	//	"id":     talon.ID,
	//})

}
