package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello word")
	godotenv.Load(".env") //go loads the envirement varibles
	// os.Getenv  build function (expored package)
	portString := os.Getenv("PORT")
	if portString == "" {
		//log.Fatal will exit the prog with code 1 and a msg
		log.Fatal("PORT is not found in the env")
	}
	fmt.Println("Port:", portString)
}
