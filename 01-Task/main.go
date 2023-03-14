package main

import (
	"github.com/Arakyn/Assignment/01-Task/Inits"
	"github.com/Arakyn/Assignment/01-Task/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	Inits.LoadEnvVariables()
	Inits.ConnectToDB()
}

func main() {

	r := gin.Default()
	// this method will create items that are stored in the database
	r.POST("/movies", controllers.MovieCreate)

	// this fetches all the movies in the database
	r.GET("/movies", controllers.MovieShow)

	// this fetches a movie from its ID
	r.GET("/movies/:id", controllers.MovieSingle)

	// this just updates the data stored in the database
	r.PUT("/movies/:id", controllers.MovieUpdate)

	// this deletes the data with the help of the primary key
	r.DELETE("/movies/:id", controllers.MovieDelete)
	r.Run()

}
