package main

import (
	"askm/gettingstarted/collections/models"
	"fmt"
)

func main() {

	fmt.Println("------------------------------")
	fmt.Println("Collections - Arrays, Slices, Maps, Structs")
	fmt.Println("------------------------------")

	arrays()
	slices()
	maps()
	structs()

	u := models.User{
		ID:        1,
		FirstName: "Anderson",
		LastName:  "Szalai",
	}

	fmt.Println(u)
}

func arrays() {

	fmt.Println("------------------------------")
	fmt.Println("Arrays")
	fmt.Println("------------------------------")

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)

	arr2 := [3]int{4, 5, 6}
	fmt.Println(arr2)
}

func slices() {
	fmt.Println("------------------------------")
	fmt.Println("Slices")
	fmt.Println("------------------------------")

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice := arr[:]
	fmt.Println(slice)

	sliceDynamic := []int{1, 2, 3, 4, 5}
	fmt.Println(sliceDynamic)

	sliceDynamic = append(sliceDynamic, 6, 7)
	fmt.Println(sliceDynamic)

	slice2 := slice[1:]
	slice3 := slice[:6]
	slice4 := slice[1:5]

	fmt.Println(slice2, slice3, slice4)
}

func maps() {
	fmt.Println("------------------------------")
	fmt.Println("Maps")
	fmt.Println("------------------------------")

	m := map[string]int{"foo": 20}

	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 40
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

}

func structs() {
	fmt.Println("------------------------------")
	fmt.Println("Structs")
	fmt.Println("------------------------------")

	type user struct {
		id        int
		firstName string
		lastName  string
	}

	var u user
	u.id = 10
	u.firstName = "Anderson"
	u.lastName = "Szalai"

	fmt.Println(u)

	u2 := user{id: 1,
		firstName: "Karen",
		lastName:  "Motta"}
	fmt.Println(u2)

}
