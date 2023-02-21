package structures

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name     string
	Director string
}
