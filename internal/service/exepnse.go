package service

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Expense struct {
	gorm.Model
	Date time.Time
	Name string
	Tags []Tag `gorm:"foreignKey:ID"`
}

func GetAllExpenses(db *DBService) ([]Expense, error) {
	var expenses []Expense
	if result := db.DB.Find(&expenses); result.Error != nil {
		return expenses, result.Error
	}
	return expenses, nil
}

func PostExpenses(db *DBService) (Expense, error) {
	expense := Expense{
		Name: "Dal Rice",
		Date: time.Now(),
		Tags: []Tag{
			{
				Name: "Lunch",
			},
		},
	}
	if result := db.DB.Create(&expense); result.Error != nil {
		return expense, result.Error
	}
	return expense, nil
}
