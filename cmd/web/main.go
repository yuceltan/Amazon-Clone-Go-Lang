package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/janpavtel/site/internal/drivers"
	"bitbucket.org/janpavtel/site/internal/handlers"
	"bitbucket.org/janpavtel/site/internal/mails"
	"bitbucket.org/janpavtel/site/internal/routes"
)

const portNumber = ":8080"

func main() {

	db, err := drivers.ConnectSQL("host=localhost port=5432 dbname=postgres user=postgres password=ebiri")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying ...")
	}
	defer db.SQL.Close()

	mailChan := make(chan mails.MailData, 1)
	defer close(mailChan)

	mails.ListenForMail(mailChan)

	fmt.Printf("Starting application on port %s \n", portNumber)

	view := handlers.NewView(db, mailChan)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes.CreateRoutes(view),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
