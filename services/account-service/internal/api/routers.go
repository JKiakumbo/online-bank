package api

import (
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func RegisterRoutes(*gin.Engine, *gocql.Session, sarama.SyncProducer) {
	// Route registration logic goes here
}
