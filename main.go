package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "blog_db"
)

type Theme struct {
	id          int
	ageCategory int
	name        string
	publicId    string
}

func printTheme(theme Theme) {
	fmt.Println("id: ", theme.id)
	fmt.Println("ageCategory: ", theme.ageCategory)
	fmt.Println("name: ", theme.name)
	fmt.Println("public id: ", theme.publicId)
}

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	rows, err := db.Query(`select id, age_category, name, public_id from themes`)
	CheckError(err)

	for rows.Next() {
		var theme Theme

		err = rows.Scan(&theme.id, &theme.ageCategory, &theme.name, &theme.publicId)
		CheckError(err)

		printTheme(theme)
	}

	err = db.Ping()
	CheckError(err)

	//	app.Start()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
