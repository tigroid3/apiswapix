package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.ru/new-swapix/api/v4/middlewares"
	"log"
	"net/http"
	"os"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) InitDb(DbDriver, DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	server.DB, err = gorm.Open(DbDriver, DBURL)
	server.DB.LogMode(os.Getenv("APP_ENV") != "prod")
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DbDriver)
		log.Fatal("This is the error:", err)
	}
}

//посмотреть способы красивой группировки методов с middleware
func (s *Server) InitRoutes() {
	//s.Router.Use(mux.CORSMethodMiddleware(s.Router))

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/test", middlewares.SetMiddlewareJSON(s.Test)).Methods("POST", "GET")

	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.UserCreate)).Methods("POST")
	s.Router.HandleFunc("/users/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.UserView)).Methods("GET")
	s.Router.HandleFunc("/users/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.UserUpdate)).Methods("PUT")

	s.Router.HandleFunc("/auth/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/auth/refreshToken", middlewares.SetMiddlewareJSON(s.RefreshToken)).Methods("POST")

	s.Router.HandleFunc("/ads/search", middlewares.SetMiddlewareJSON(s.AdSearch)).Methods("GET")
	s.Router.HandleFunc("/ads/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.AdView)).Methods("GET")
	s.Router.HandleFunc("/ads", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.AdCreate))).Methods("POST")
	s.Router.HandleFunc("/ads", middlewares.SetMiddlewareJSON(s.AdDelete)).Methods("DELETE")
	s.Router.HandleFunc("/ads", middlewares.SetMiddlewareJSON(s.AdIndex)).Methods("GET")
	s.Router.HandleFunc("/ads", middlewares.SetMiddlewareJSON(s.AdUpdate)).Methods("PUT")

	s.Router.HandleFunc("/adImages", middlewares.SetMiddlewareJSON(s.AdImageCreate)).Methods("POST")
	s.Router.HandleFunc("/adImages/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.AdImageDelete)).Methods("DELETE")
	s.Router.HandleFunc("/adImages/default/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.AdImageSetDefault)).Methods("PUT")
	s.Router.HandleFunc("/adImages/ad/{id:[0-9]+}", middlewares.SetMiddlewareJSON(s.AdImageGetForAd)).Methods("GET")
}

func (s *Server) Start(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.Router))
}

//func (s *Server) InitChannels() {
//
//}
