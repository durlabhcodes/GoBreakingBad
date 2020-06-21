package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/magiconair/properties"
)

func helloWorld(writer http.ResponseWriter, request *http.Request) {
	log.Println("Method -->" + request.Method)
	fmt.Fprint(writer, "Hey You!!")
}

//HandleRequests is an amazing func
func HandleRequests() {
	LoadProperties()
	http.HandleFunc("/hello", helloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//LoadProperties loads properties
func LoadProperties() {
	props, _ := properties.LoadFile("application.properties", properties.UTF8)
	val, ok := props.Get("hello")
	if ok {
		fmt.Printf("Property is %s", val)
	} else {
		fmt.Println("Error in reading props")
	}
}
