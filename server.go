package main


import (
	"net/http"

	"server.go/configs"
)


func main(){
	configs.LoadEnv()
	configs.ConnectDB()

	http.ListenAndServe(":8090", nil)
}
