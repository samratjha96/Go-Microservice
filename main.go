package main
import (
		"fmt"
		"github.com/samratjha96/Go-Microservice/service"
        )
var appName = "accountservice"
func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}