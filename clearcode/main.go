package main

import (
	"clear_code/database"
	handle "clear_code/handelers"
	"clear_code/routing"
	"clear_code/usecase"
)

func main() {
	conn := database.Connect()
	userUsecse := usecase.NewSrv(conn)
	handler := handle.NewHandler(userUsecse)
	routing.HandlerRouting(handler)
	//HandlerRouting()

}
