package main

import (
	"github.com/InspectorGadget/goshort/initializers"
	"github.com/InspectorGadget/goshort/models"
)

func init() {
	if err := initializers.ConnectToDB(); err != nil {
		panic(err)
	}
}

func main() {
	if err := initializers.DB.AutoMigrate(
		&models.User{},
		&models.Url{},
		&models.Token{},
		&models.Role{},
		&models.RoleMap{},
	); err != nil {
		panic(err)
	}
}
