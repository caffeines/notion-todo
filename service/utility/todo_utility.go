package utility

import (
	"bytes"
	"encoding/json"

	"github.com/caffeines/notion-todo/models"
)

// NewTodoProperties returns a new Properties
func newTodoProperties(title, date string) models.ItemData {
	return models.NewProperties(title, date)
}

// GetCreateTodoBody returns a new Todo body
func GetCreateTodoBody(title, date, databaseId string) (*bytes.Buffer, error) {
	itemData := newTodoProperties(title, date)
	payload := models.CreateTodoPayload{
		Parent: models.Parent{
			DatabaseID: databaseId,
		},
		Properties: itemData,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(jsonPayload), nil
}
