package services

type IObject interface {
	create(name string, params map[string]interface{})
	update(id string, params map[string]interface{})
	delete(id string)
}

