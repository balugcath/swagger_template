package main

import (
	"log"
	"net/http"

	"swagger_template/src/main_service/api_handler"
	"swagger_template/src/main_service/client"
	"swagger_template/src/main_service/middleware"
	"swagger_template/src/main_service/sql"
	"swagger_template/src/main_service/swagger/restapi"
)

func main() {

	pg, err := storage.NewPG().Open("conn string")
	if err != nil {
		log.Fatalln(err)
	}

	var cli client.Client

	mw := middleware.NewMiddleware(pg, cli)
	api := apihandler.NewTodo(mw)

	h, err := restapi.Handler(restapi.Config{TodosAPI: api})
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatal(http.ListenAndServe(":8080", h))
}
