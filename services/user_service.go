package services

import types "server.go/models"

type user = types.User

func (u *user) create(params map[string]interface{}){
	u.Id = params["_id"]
	u.Name = params["name"]
}

