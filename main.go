package main

import (
	"go-web-native/config"
	categoriescontroller "go-web-native/controllers/categoriesController"
	homecontroller "go-web-native/controllers/homeController"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//1. homepage
	http.HandleFunc("/", homecontroller.Welcome)
	// 2. Categories
	http.HandleFunc("/categories", categoriescontroller.Index)
	http.HandleFunc("/categories/add", categoriescontroller.Add)
	http.HandleFunc("/categories/edit", categoriescontroller.Edit)
	http.HandleFunc("/categories/delete", categoriescontroller.Delete)

	log.Println("Server Running Port 666")
	http.ListenAndServe(":666", nil)
}
