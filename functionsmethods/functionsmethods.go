package main

import (
	"askm/gettingstarted/functionsmethods/controllers"
	"fmt"
	"net/http"
)

func main() {
	title("Functions and Methods", true)

	port := 3000
	_, err := startWebServer(port) // _ ignores the first result not requiring the program to use the result
	fmt.Println(port, err)
}

func startWebServer(port int) (int, error) {
	title("Starting web server...", false)

	controllers.RegisterControllers()
	http.ListenAndServe(":3001", nil)

	title("Web server started on port "+fmt.Sprint(port), false)

	return port, nil
}

func title(title string, separator bool) {
	if separator {
		fmt.Println("--------------------------------------------------")
	}
	fmt.Println(title)
	if separator {
		fmt.Println("--------------------------------------------------")
	}
}
