package routes

import "net/http"

//Ping Method to check health
func Ping(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Success v2"))
}
