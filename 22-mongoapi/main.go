package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a7mod/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()

	fmt.Println("Server is getting starteddd....")
	log.Fatal(http.ListenAndServe(":4000", r))

	log.Println("Listening at port 4000 ...")

}
