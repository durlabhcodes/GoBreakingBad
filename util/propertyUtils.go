package util

import (
	"log"

	"github.com/magiconair/properties"
)

//LoadProperties loads properties
func LoadProperties(propertiesFileNamesArr []string) *properties.Properties {
	props, error := properties.LoadFiles(propertiesFileNamesArr, properties.UTF8, true)
	if error == nil {
		return props
	}

	log.Panicf("Property file with name %s not loaded with error : %v", propertiesFileNamesArr, error.Error())
	return nil
}

//GetFromProperties is GetFromProperties
func GetFromProperties(propertiesFileNameArr []string, propertyName string) string {
	defer recoverPropertyGet()
	properties := LoadProperties(propertiesFileNameArr)
	if properties != nil {
		prop, ok := properties.Get(propertyName)

		if ok {
			return prop
		}

		log.Panicf("No property found with name %s", propertyName)
		return ""
	}
	log.Panicf("Properties not loaded from filenames : %v", propertiesFileNameArr)
	return ""
}

func recoverPropertyGet() {
	if r := recover(); r != nil {
		log.Println("Recovered from ", r)
	}
}
