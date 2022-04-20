package seed

import (
	"log"

	"github.com/VnFood/final/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		Username: "LoiTPH123",
		Password: "LoiTPH123",
	},
	{
		Username: "LoiTPH234",
		Password: "LoiTPH234",
	},
}

var posts = []models.Post{
	{
		Title:   "Ayayayoooo 1",
		Content: "Ayayayoooo 1",
	},
	{
		Title:   "Ayayayoooo 2",
		Content: "Ayayayoooo 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
