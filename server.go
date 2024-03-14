package main


import (
	"net/http"

	"server.go/configs"
)


func main(){
	configs.ConnectDB()
	configs.LoadEnv()

	http.ListenAndServe(":8090", nil)
}
