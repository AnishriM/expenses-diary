package tag

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name string
}

type TagService interface {
	GetTagByID(db *DBService, ID uint) (Tag, error)
	GetAllTags(db *DBService) ([]Tag, error)
	UpdateTag(ID uint, name string, db *DBService) (Tag, error)
	DeleteTag(ID uint, db *DBService) (Tag, error)
	CreateTag(name string, db *DBService) (Tag, error)
}

func GetTagByID(db *DBService, ID uint) (Tag, error) {
	tag := Tag{}
	if result := db.DB.First(&tag, ID); result.Error != nil {
		return tag, result.Error
	}
	return tag, nil
}

func GetAllTags(db *DBService) ([]Tag, error) {
	var tags []Tag
	if result := db.DB.Find(&tags); result.Error != nil {
		return tags, result.Error
	}
	return tags, nil
}

func UpdateTag(ID uint, name string, db *DBService) (Tag, error) {
	var tag Tag
	var err error
	tag, err = GetTagByID(db, ID)
	if err != nil {
		return tag, err
	}

	tag.Name = name
	if result := db.DB.Save(&tag); result.Error != nil {
		return tag, result.Error
	}
	return tag, nil
}

func DeleteTag(ID uint, db *DBService) (Tag, error) {
	var tag Tag
	var err error
	tag, err = GetTagByID(db, ID)
	if err != nil {
		return tag, err
	}
	if result := db.DB.Delete(&tag, ID); result.Error != nil {
		return tag, result.Error
	}
	return tag, nil
}

func CreateTag(name string, db *DBService) (Tag, error) {
	tag := Tag{
		Name: name,
	}

	if result := db.DB.Create(&tag); result.Error != nil {
		return tag, result.Error
	}
	return tag, nil
}
