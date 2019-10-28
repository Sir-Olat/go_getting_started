package main

import (
	"net/http"
	"github.com/sir-olat/webservices/controllers"
	// "github.com/sir-olat/webservices/models"
	"fmt"
)

const (
	pi = 3.142
	first = iota
	second = iota
)

func main() {
	// variables()
	// array()
	// slice()
	// maps()
	// structs()
	// model()
	// _, err := startWebServer(3000)
	// fmt.Println(err)
	controllers.registerControllers()
	http.ListenAndServe(":3000", nil)
	// infiniteLoop()
	// loopSlice()
	// loopMap()
	// ifLogic()
	// removeItemFromSlice(5)
	// requestMethod()
}

// func model() {
// 	u := models.User{
// 		ID: 2,
// 		FirstName: "Bosun",
// 		LastName: "Ogunlana",
// 	}
// 	fmt.Println(u)
// }

// func variables(){
// 	var i int
// 	i = 42
// 	fmt.Println(i)

// 	var f float32 = 3.14
// 	fmt.Println(f)

// 	firstName := "Bosun"
// 	fmt.Println(firstName)

// 	ptr := &firstName
// 	fmt.Println(ptr, *ptr)

// 	firstName = "Tunbosun"
// 	fmt.Println(ptr, *ptr)

// 	b := true
// 	fmt.Println(b)

// 	c := complex(3, 4)
// 	r, im := real(c), imag(c)
// 	fmt.Println(r, im, c)

// 	// Pointers
// 	/* 
// 	Use * behind the var type to initialize a pointer
// 	to print the address, just print the var name
// 	to store value in the address, preeceed the assignment with *
// 	to print the content of a mem address, preceed the variable with *
// 	*/
// 	var lastname *string = new(string)
// 	*lastname = "Oguns"

// 	fmt.Println(*lastname)

// 	fmt.Println(pi)
// 	fmt.Println(first)
// 	fmt.Println(second)
// }

// func array() {
	
// 	var arr [3]int
// 	arr[0] = 1
// 	arr[1] = 2
// 	arr[2] = 3
// 	fmt.Println(arr)

// 	arr2 := [3]int{1,2,3}
// 	fmt.Println(arr2)
// }

// func slice() {
// 	slice := []int{1,2,3}
// 	fmt.Println(slice)

// 	slice = append(slice, 4, 2, 34, 43, 123)
// 	fmt.Println(slice)

// 	s2 := slice[1:]
// 	s3:= slice[:2]
// 	s4 := slice[1:2]

// 	fmt.Println(s2, s3, s4)
// }

// func maps(){
// 	m := map[string]int{"bar": 10}
// 	m["foo"] = 42
// 	fmt.Println(m)
// 	fmt.Println(m["foo"])

// 	delete(m, "foo")
// 	fmt.Println(m)
// }

// func structs(){
	
// 	type user struct {
// 		ID int
// 		Firstname string
// 		LastName string
// 	}
// 	// First initialization method
// 	var u user 
// 	u.ID = 1
// 	u.Firstname = "Bosun"
// 	u.LastName = "Oguns"
// 	fmt.Println(u)

// 	// Second initialization method
// 	u2 := user{ ID: 1, 
// 		Firstname: "Bosun", 
// 		LastName: "Oguns",
// 	}
// 	fmt.Println(u2)
// }

// func startWebServer(port int) (int, error) {
// 	fmt.Println("Starting server")
// 	// do something
// 	fmt.Println("Server started on port", port)
// 	// return errors.New("Something went wrong")
// 	return port, nil
// }

func looping() {
	for i := 0 ;i < 5; i++ {
		println(i)
	}

	var i int
	for i < 5 {
		println(i)
		i++
	}
	// println(i)
}

func infiniteLoop(){
	var i int 
	// for ; ; {
	// 	// first method
	// 	if i == 5 {
	// 		break
	// 	}
	// 	println(i)
	// 	i++
	// }

	for {
		// second method
		if i == 5 {
			break
		}
		println(i)
		i++
	}

}

func loopSlice() {
	slice := []int{1,2,3}
	for i, v := range slice {
		println(i, v)
	}
}

func loopMap(){
	ports := map[string]int{"http": 80, "https": 443}
	for k, v := range ports {
		if v == ports["http"] {
			continue
		}
		println(k,v)
	}
}

type User struct {
	ID int
	FirstName string
	LastName string
}

func ifLogic() {
	
	u1 := User{
		ID: 1,
		FirstName: "Bosun",
		LastName: "Rasaq",
	}

	u2 := User{
		ID: 2,
		FirstName: "Tunbosun",
		LastName: "Ogunlana",
	}

	if u1.ID == u2.ID {
		println("same user")
	} else if u1.FirstName == u2.FirstName{
		println("Same user")
	} else {
		println("diferent user")
	}

}

func removeItemFromSlice(item int) {
	slice :=  []int{1,2,3,4,4,5,6,6,7}
	for i, v := range slice {
		if v == item {
		 slice = append(slice[:i], slice[i+1:]...)
		}
	}
	fmt.Println(slice)
}

type HTTPRequest struct {
	Method string
}

func requestMethod() {
	r := HTTPRequest{ Method: "GET" }
	switch r.Method {
	case "GET":
		fmt.Println("Get request")
	case  "DELETE":
		fmt.Println("DELETE request")
	case "POST":
		fmt.Println("POST request")
	case "PUT":
		fmt.Println("PUT request")
	default:
		fmt.Println("unknown method")
	}
}