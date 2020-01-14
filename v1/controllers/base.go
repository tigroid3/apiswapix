package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tigroid3/apiswapix/v1/middlewares"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) InitDb(DbDriver, DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	server.DB, err = gorm.Open(DbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DbDriver)
		log.Fatal("This is the error:", err)
	}
}

func (s *Server) InitRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.Home))).Methods("GET")

	s.Router.HandleFunc("/test", middlewares.SetMiddlewareJSON(s.Test)).Methods("POST", "GET")

	s.Router.HandleFunc("/user/register", middlewares.SetMiddlewareJSON(s.Register)).Methods("POST")
	s.Router.HandleFunc("/user/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/user/refresh-token", middlewares.SetMiddlewareJSON(s.RefreshToken)).Methods("POST")

	s.Router.HandleFunc("/ad", middlewares.SetMiddlewareJSON(s.AdIndex)).Methods("GET")

	////Users routes

	//s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
}

func (server *Server) Start(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
