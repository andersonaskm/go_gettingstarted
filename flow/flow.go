package main

import "fmt"

func main() {

	title("Flow execution - Loop, panics, if, switches", false)

	loopTillCondition()
	loopWithPostClause()
	infiniteLoop()
	loopOverObjects()
	throwPanic()
	switchStatements()
}

func loopTillCondition() {

	title("Loop Till Condition", true)

	var i int
	for i < 5 {

		fmt.Println(i)
		i++

		if i == 3 {
			continue
		}

		fmt.Println("continue...")

	}
}

func loopWithPostClause() {

	title("Loop with post clause", true)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func infiniteLoop() {

	title("Infinite Loop", true)

	var i int
	for {
		if i == 5 {
			break
		}

		fmt.Println(i)

		i++
	}
}

func loopOverObjects() {
	title("Loop over objects", true)

	slice := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for i, v := range slice {
		println(i, v)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}

	for i, v := range wellKnownPorts {
		println(i, v)
	}

	for _, v := range wellKnownPorts { // only print the values
		println(v)
	}

}

func throwPanic() {

	title("Throwing panic", true)
	//panic("Something bad happened")
}

type HTTPRequest struct {
	Method string
}

func switchStatements() {

	title("Switch Statements", true)

	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("Get Request")
	case "DELETE":
		println("Delete Request")
	case "POST":
		println("Post Request")
	case "PUT":
		println("Put Request")
	default:
		println("Unhandled method")
	}

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
