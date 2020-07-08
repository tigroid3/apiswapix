package v4

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gitlab.ru/new-swapix/api/v4/controllers"
	"os"
)

func Run() {
	//var err error
	err = godotenv.Load("/go/bin/.env")
	//if err != nil {
	//	log.Fatalf("Error godotenv %s", err)
	//}

	serv := controllers.Server{}
	serv.InitDb(
		"postgres",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	serv.Router = mux.NewRouter()
	serv.InitRoutes()
	serv.Start(":8080")
}
