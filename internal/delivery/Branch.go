package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
	"github.com/SicParv1sMagna/mdhh_backend/internal/pkg/middleware/decode"
	"github.com/SicParv1sMagna/mdhh_backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
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
			TalonCount:          branch.TalonCount,
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
			TalonCount:          branch.TalonCount,
		}

		branchResponses = append(branchResponses, response)
	}

	c.JSON(http.StatusOK, branchResponses)
}

func GetBranchById(repository *repository.Repository, c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var branch model.Branch

	branch, err = repository.GetBranchById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

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

	branchResponse := model.BranchResponse{
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
		TalonCount:          branch.TalonCount,
	}

	c.JSON(http.StatusOK, branchResponse)
}

func GetNearestBranches(repository *repository.Repository, latitude string, longitude string) ([]model.BusinessResponse, error) {
	var response []model.BusinessResponse

	for i := 1; i < 11; i++ {
		branch, err := repository.GetBranchById(i)
		if err != nil {
			return []model.BusinessResponse{}, err
		}

		response = append(response, model.BusinessResponse{ID: branch.Branch_ID, TalonCount: branch.TalonCount})
	}
	return response, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

//func toResponseBranch(branch model.Branch) (model.BranchResponse, error) {
//	openHours, err := decode.UnmarshalOpenHours(branch.OpenHours)
//	if err != nil {
//		return model.BranchResponse{}, err
//	}
//
//	openHoursIndividual, err := decode.UnmarshalOpenHours(branch.OpenHoursIndividual)
//	if err != nil {
//		return model.BranchResponse{}, err
//	}
//
//	return model.BranchResponse{
//		Branch_ID:           branch.Branch_ID,
//		SalePointName:       branch.SalePointName,
//		Address:             branch.Address,
//		Status:              branch.Status,
//		OpenHours:           openHours,
//		RKO:                 branch.RKO,
//		OpenHoursIndividual: openHoursIndividual,
//		OfficeType:          branch.OfficeType,
//		SalePointFormat:     branch.SalePointFormat,
//		SUOAvailability:     branch.SUOAvailability,
//		HasRamp:             branch.HasRamp,
//		Latitude:            branch.Latitude,
//		Longitude:           branch.Longitude,
//		MetroStation:        branch.MetroStation,
//		Distance:            branch.Distance,
//		KEP:                 branch.KEP,
//		MyBranch:            branch.MyBranch,
//		Network:             branch.Network,
//		SalePointCode:       branch.SalePointCode,
//		TalonCount:          branch.TalonCount,
//	}, nil
//}

func GetBranchesWithTalons(repository *repository.Repository, c *gin.Context) {
	latitude := c.DefaultQuery("latitude", "")
	longitude := c.DefaultQuery("longitude", "")

	nearestBranches, err := GetNearestBranches(repository, latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	// to run rabbit: docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management
	// rabbit connect
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "\"Failed to declare an exchange\"",
			"message": err.Error(),
		})
		return
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "\"Failed to declare a queue\"",
			"message": err.Error(),
		})
		return
	}
	for _, br := range nearestBranches {
		log.Printf("Binding queue %s to exchange %s with routing key %d",
			q.Name, "logs_direct", br.ID)

		err = ch.QueueBind(
			q.Name,                              // queue name
			strconv.FormatInt(int64(br.ID), 10), // routing key
			"logs_direct",                       // exchange
			false,
			nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "\"Failed to bind a queue\"",
				"message": err.Error(),
			})
			return
		}
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "\"Failed to register a consumer\"",
			"message": err.Error(),
		})
		return
	}

	// open websocket connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
	defer conn.Close()

	// send data
	jsonData, err := json.Marshal(nearestBranches)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Println("Error: ", jsonData)
	}
	log.Println("Sent: ", nearestBranches)

	// send new data (notifications)
	for msg := range msgs {
		err = conn.WriteMessage(websocket.TextMessage, msg.Body)
		if err != nil {
			log.Println("Error: ", err.Error())
			return
		}

		var branch model.Branch
		err = json.Unmarshal(msg.Body, &branch)
		if err != nil {
			log.Println("Error: ", err.Error())
			return
		}

		log.Println("Sent: ", branch)
	}

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	//c.JSON(http.StatusOK, gin.H{
	//	"status": "success",
	//	"id":     talon.ID,
	//})

}
