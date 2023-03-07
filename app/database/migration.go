package database

import (
	"fmt"
	class "immersiveApp/features/classes/data"
	status "immersiveApp/features/statuses/data"
	team "immersiveApp/features/teams/data"
	user "immersiveApp/features/users/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
		&team.Team{},
		&class.Class{},
		&status.Status{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
