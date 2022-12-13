package dbConnect

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "EventEvonik"
	password = "EventEvonik1213"
	dbname   = "EventEvonik"
)

type App struct {
	DB *sql.DB
}
type error interface {
	Error() string
}

func (a *App) Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

}

// func Start() {
// 	a := App{}

// 	a.port = ":9003"
// 	fmt.Print(a)

// 	a.init()
// 	a.Run()
// }
