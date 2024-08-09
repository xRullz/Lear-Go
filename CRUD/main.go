package main

import (
	"crud/config"
	categorycontroller "crud/controllers/CategoryController"
	homecontroller "crud/controllers/HomeController"
	productcontroller "crud/controllers/ProductController"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//Category
	http.HandleFunc("/category", categorycontroller.Index)
	http.HandleFunc("/category/create", categorycontroller.Create)
	http.HandleFunc("/category/update", categorycontroller.Update)
	http.HandleFunc("/category/delete", categorycontroller.Delete)

	//Product
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/create", productcontroller.Create)
	http.HandleFunc("/product/update", productcontroller.Update)
	http.HandleFunc("/product/delete", productcontroller.Delete)
	http.HandleFunc("/product/detail", productcontroller.Detail)


    log.Println("Server is listening on port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}