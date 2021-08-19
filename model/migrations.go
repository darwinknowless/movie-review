package model

import "movie-review/connection"

func Migration() {
	db := connection.GetConnection()

	db.Migrator().DropTable(&Movie{})
	db.Migrator().DropTable(&Author{})

	checkMovie := db.Migrator().HasTable(&Movie{})
	if !checkMovie {
		db.Migrator().CreateTable(&Movie{})
	}
	checkAuthor := db.Migrator().HasTable(&Author{})
	if !checkAuthor {
		db.Migrator().CreateTable(&Author{})
	}
}
