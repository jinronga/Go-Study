package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-study/go-backend/api/middleware"
	"github.com/go-study/go-backend/bootstrap"
	"github.com/go-study/go-backend/mongo"
	swaggerFiles "github.com/swaggo/files"
	ginSwgger "github.com/swaggo/gin-swagger"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")

	publicRouter.GET("/swgger/*any", ginSwgger.WrapHandler(swaggerFiles.Handler))

	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
}
