package tag

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBService struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *DBService {
	return &DBService{
		DB: db,
	}
}

func NewDatabase() (*gorm.DB, error) {
	println("Setting up database")

	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_table := os.Getenv("DB_TABLE")

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", db_host, db_port, db_username, db_table, db_password)

	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateDB(db *gorm.DB) error {
	var tag Tag
	if result := db.AutoMigrate(&tag); result.Error != nil {
		return result.Error
	}
	return nil
}
