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
	baseURL := GetFromProperties("application.properties", "breaking.bad.api.baseurl")
	fmt.Println("Base url = ", baseURL)

	http.HandleFunc("/hello", helloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//LoadProperties loads properties
func LoadProperties(propertiesFileName string) *properties.Properties {
	props, error := properties.LoadFile(propertiesFileName, properties.UTF8)
	if error == nil {
		return props
	}

	log.Panicf("Property file with name %s not loaded with error : %s", propertiesFileName, error.Error())
	return nil
}

//GetFromProperties is GetFromProperties
func GetFromProperties(propertiesFileName string, propertyName string) string {
	defer recoverPropertyGet()
	properties := LoadProperties(propertiesFileName)
	if properties != nil {
		prop, ok := properties.Get(propertyName)

		if ok {
			return prop
		}

		log.Panicf("No property found with name %s", propertyName)
		return ""
	}
	log.Panicf("Properties not loaded from filename : %s", propertiesFileName)
	return ""
}

func recoverPropertyGet() {
	if r := recover(); r != nil {
		log.Println("Recovered from ", r)
	}
}
