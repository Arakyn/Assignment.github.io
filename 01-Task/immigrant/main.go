package main

import (
	"github.com/Arakyn/Assignment/01-Task/Inits"
	"github.com/Arakyn/Assignment/01-Task/structures"
)

func init() {
	Inits.LoadEnvVariables()
	Inits.ConnectToDB()
}

func main() {
	// Inits.DB.AutoMigrate(&structures.Movie{})
	Inits.DB.AutoMigrate(&structures.Book{})
}
