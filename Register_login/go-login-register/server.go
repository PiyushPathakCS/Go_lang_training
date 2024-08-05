package main

import (
	_ "go-login-register/docs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func HandlerRouting() {

	r := mux.NewRouter()
	r.HandleFunc("/signup", Register).Methods("POST")
	r.HandleFunc("/login", Authenticate).Methods("POST")
	r.HandleFunc("/logout", Logout).Methods("POST")
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8000", r))

}
