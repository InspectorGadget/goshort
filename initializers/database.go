package initializers

import (
	"fmt"

	"github.com/InspectorGadget/goshort/helpers"
	"github.com/InspectorGadget/goshort/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() error {
	dsn := helpers.GetDatabaseDsn()
	if dsn == "" {
		return fmt.Errorf("database connection string is empty, please check your environment variables")
	}

	conn, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		return err
	}

	DB = conn

	return nil
}

func Migrate() error {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Url{},
		&models.Token{},
		&models.Role{},
		&models.RoleMap{},
	)

	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	fmt.Println("Database migration completed successfully")

	return nil
}
