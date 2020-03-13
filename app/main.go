package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"app/src/controllers"
	"app/src/routes"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func load() {
	if r, _ := exists("/api/data/initialData.txt"); r {
		fmt.Println("CSV file already loaded")
		return
	}

	f, err := os.OpenFile("/api/data/initialData.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	if err != nil {
		fmt.Println("Failed on loading initial data!", err)
		return
	}

	controller := controllers.CompanyController{}

	controller.LoadCompanies("q1_catalog.csv")
	f.Write([]byte("Success"))

	log.Print("Loaded")
}

func main() {
	load()
	router := routes.NewRouter()
	log.Print("Iniciou")
	log.Fatal(http.ListenAndServe(":5000", router))
}
