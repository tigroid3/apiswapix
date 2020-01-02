package v1

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tigroid3/apiswapix/v1/controllers"
	"log"
	"os"
)

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error godotenv %s", err)
	}

	serv := controllers.Server{}
	serv.InitDb(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	serv.Router = mux.NewRouter()
	serv.InitRoutes()
	serv.Start(":7777")
}
