package main

import (
	"learning-go/build-a-rest-api-with-golang-from-scratch-postgresql-with-gorm-and-gin-web-framework/controllers"
	"learning-go/build-a-rest-api-with-golang-from-scratch-postgresql-with-gorm-and-gin-web-framework/models"
	"learning-go/build-a-rest-api-with-golang-from-scratch-postgresql-with-gorm-and-gin-web-framework/util"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cann't load config", err)
	}

	r := gin.Default()

	db := models.SetupModels(config.DBUrl)

	// provide db variaable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
