package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-study/go-backend/bootstrap"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewLoginRouter(env, timeout, db, publicRouter)
}
