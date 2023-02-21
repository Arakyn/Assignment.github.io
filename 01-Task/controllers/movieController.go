package controllers

import (
	"log"

	"github.com/arakyn/Assignment/01-Task/Inits"
	"github.com/arakyn/Assignment/01-Task/structures"
	"github.com/gin-gonic/gin"
)

//
//               WAS SENDING DATA THROUGH POSTMAN
//
//
func MovieCreate(c *gin.Context) {
	// Gettings things to change

	// declared a movie struct to bind incoming details to it
	var movie struct {
		Name     string
		Director string
	}
	c.Bind(&movie)
	// made a new struct of type structure.Movie as a template and all the field data would be directly stored in it
	// and it will be added to the database
	mov := structures.Movie{Name: movie.Name, Director: movie.Director}
	result := Inits.DB.Create(&mov)

	// if there is error then it will send status 400
	if result.Error != nil {
		c.Status(400)
		log.Fatal("Error while creating things")
		return
	}
	// this just shows the data you entered
	c.JSON(200, gin.H{
		"Movie": mov,
	})

}

// this shows all the movies that are in the database
func MovieShow(c *gin.Context) {
	// declaring a array of struct movies so it can store all the things in it
	var Movie []structures.Movie
	// Init.DB.Find basically searches all the data in the database and stores it in the array movie

	result := Inits.DB.Find(&Movie)
	if result.Error != nil {
		log.Fatal("Error at finding stuff")
		c.Status(500)
		return
	}
	// showing it on the page
	c.JSON(200, gin.H{
		"Movies": Movie,
	})

}

func MovieSingle(c *gin.Context) {
	// stripping the id from the URL
	id := c.Param("id")

	// declaring a empty var to store the fetched data
	var movie structures.Movie

	// fetching the data from the database
	Inits.DB.First(&movie, id)

	c.JSON(200, gin.H{
		"Movies": movie,
	})

}

func MovieUpdate(c *gin.Context) {

	// stripping the ID from the url
	id := c.Param("id")

	// making a empty like struct to store the updated catched values from the request
	var catchMovie struct {
		Name     string
		Director string
	}

	// binding the values to it
	c.Bind(&catchMovie)

	// fetching the data from the id and storing it somewhere empty
	var oldMovie structures.Movie
	Inits.DB.First(&oldMovie, id)

	// updating the data in the oldMovie and using the CatchedMovie data to store it there
	Inits.DB.Model(&oldMovie).Updates(structures.Movie{Name: catchMovie.Name, Director: catchMovie.Director})

	c.JSON(200, gin.H{
		"Movies": catchMovie,
	})
}

// func delete just removes it, so it cant be shown on the website or something but the actual data
// is present in the database for data recovery.
func MovieDelete(c *gin.Context) {
	// stripping the id of the url
	id := c.Param("id")

	// deleting the data using the ID which is the primary key set in gorm.Model
	Inits.DB.Delete(&structures.Movie{}, id)
}
