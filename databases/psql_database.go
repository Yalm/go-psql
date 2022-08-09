package databases

import (
	"fmt"

	"github.com/Yalm/go-psql/config"
	"github.com/Yalm/go-psql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnectionToPostgreSql(appConfig *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		appConfig.PsqlConfig.Host,
		appConfig.PsqlConfig.Username,
		appConfig.PsqlConfig.Password,
		appConfig.PsqlConfig.Database,
		appConfig.PsqlConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	return db
}
