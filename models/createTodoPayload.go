package models

type Parent struct {
	DatabaseID string `json:"database_id"`
}

type CreateTodoPayload struct {
	Parent     Parent   `json:"parent"`
	Properties ItemData `json:"properties"`
}
