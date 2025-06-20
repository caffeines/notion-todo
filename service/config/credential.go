package config

import "github.com/caffeines/notion-todo/models"

type Credential interface {
	SetConfig(token string, databaseID string) error
	GetConfig() (*models.Config, error)
}
