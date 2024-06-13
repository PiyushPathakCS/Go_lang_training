package main

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.close()
}
