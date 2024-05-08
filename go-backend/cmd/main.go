package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-study/go-backend/api/route"
	"github.com/go-study/go-backend/bootstrap"
	_ "github.com/go-study/go-backend/cmd/docs"
	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

// @title			启动文件
// @version		1.0
// @description	启动文件
// @host			localhost:8080
// @BasePath
func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	//ginSwagger.WrapHandler(swaggerfiles.Handler,
	//	ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
	//	ginSwagger.DefaultModelsExpandDepth(-1))
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)

}
