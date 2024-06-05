package main

import (
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/gami/ps2-gami-common/adapters/queue"
	"gitlab.id.vin/gami/ps2-gami-common/logger"

	"log"
	"net/http"
	"strconv"
	"test/src/kafka/config"
	"test/src/kafka/models"
)

var ErrUserNotFoundInProducer = errors.New("User Not Found")

func findUserByID(id int, users []models.User) (models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, ErrUserNotFoundInProducer
}

func getIDFromRequest(formValue string, ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.PostForm(formValue))
	if err != nil {
		return 0, fmt.Errorf(
			"failed to parse ID from form value %s: %w", formValue, err)
	}
	return id, nil
}

func sendKafkaMessage(producer queue.KafkaProducer, users []models.User, ctx *gin.Context, fromID, toID int) error {
	//message := ctx.PostForm("message")

	//fromUser, err := findUserByID(fromID, users)
	//if err != nil {
	//	return err
	//}
	//
	//toUser, err := findUserByID(toID, users)
	//if err != nil {
	//	return err
	//}
	//
	//notification := models.Notification{
	//	From:    fromUser,
	//	To:      toUser,
	//	Message: message,
	//}

	//notificationJSON, err := json.Marshal(notification)
	//if err != nil {
	//	return fmt.Errorf("failed to marshal notification: %w", err)
	//}
	msg := &queue.KafkaMessage{
		Topic: config.KafkaTopic,
		Key:   []byte("12424"),
		Value: []byte("!24"),
	}

	producer.SendMessageAsync(msg)
	return nil
}

func sendMessageHandler(producer queue.KafkaProducer, users []models.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fromID, err := getIDFromRequest("fromID", ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		toID, err := getIDFromRequest("toID", ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err = sendKafkaMessage(producer, users, ctx, fromID, toID)
		if errors.Is(err, ErrUserNotFoundInProducer) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Notification send successfully!",
		})
	}

}

func setupProducer() (sarama.SyncProducer, error) {
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	//producer, err := sarama.NewSyncProducer(, configs)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to setup producer: %w", err)
	//}
	//return producer, nil
	return nil, nil
}

func main() {
	users := []models.User{
		{ID: 1, Name: "Emma"},
		{ID: 2, Name: "Bruno"},
		{ID: 3, Name: "Rick"},
		{ID: 4, Name: "Lena"},
	}
	producerConfig := queue.ProducerConfig{
		SeedBrokers:    []string{"confluent-kafka-tw-1.int.onemount.dev:9093", "confluent-kafka-tw-2.int.onemount.dev:9093", "confluent-kafka-tw-3.int.onemount.dev:9093", "confluent-kafka-sg-4.int.onemount.dev:9093", "confluent-kafka-sg-5.int.onemount.dev:9093", "confluent-kafka-sg-6.int.onemount.dev:9093"},
		ClientKeyFile:  "/Users/huannguyen/Downloads/keys/tcb-promotion-rule-engine-clients.int.onemount.dev/tcb-promotion-rule-engine-clients.int.onemount.dev-key.pem",
		ClientCertFile: "/Users/huannguyen/Downloads/keys/tcb-promotion-rule-engine-clients.int.onemount.dev/tcb-promotion-rule-engine-clients.int.onemount.dev-cert.pem",
		ClientCAFile:   "/Users/huannguyen/Downloads/keys/truststore-tcb/truststore.pem",
	}

	producer, err := queue.NewKafkaProducer(producerConfig, true)
	if err != nil {
		logger.Fatalf(err, "Could not start kafka producer %v", err)
	}
	//producer, err := setupProducer()
	//if err != nil {
	//	log.Fatalf("failed to initialize producer: %v", err)
	//}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/send", sendMessageHandler(producer, users))

	fmt.Printf("Kafka PRODUCER started at http://localhost%s\n", config.ProducerPort)

	if err = router.Run(config.ProducerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
