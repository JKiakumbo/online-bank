package account

import (
	"account-service/config"
	"account-service/internal/api"
	"account-service/internal/db"
	"account-service/internal/kafka"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	session, err := db.NewCassandraSession(cfg.DBHost)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	producer, err := kafka.NewProducer(cfg.KafkaBrokers)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	router := gin.Default()
	api.RegisterRoutes(router, session, producer)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.ServicePort
	}

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}
