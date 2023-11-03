package main

import (
	"fmt"
	"github.com/hudayberdipolatov/go-blog-web/models"
	"github.com/hudayberdipolatov/go-blog-web/routes"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	models.ConnectDataBase()
	log.Println("Server run port", portNumber)
	fmt.Println("----------------------------------")

	http.ListenAndServe(portNumber, routes.Routes())
}
