package initializers

import (
	"fmt"

	"github.com/InspectorGadget/goshort/helpers"
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
