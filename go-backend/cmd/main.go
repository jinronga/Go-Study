package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-study/go-backend/api/route"
	"github.com/go-study/go-backend/bootstrap"
	"time"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)

}
