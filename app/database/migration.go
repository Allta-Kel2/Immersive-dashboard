package database

import (
	"fmt"
	team "immersiveApp/features/teams/data"
	user "immersiveApp/features/users/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&team.Team{},
		&user.User{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
