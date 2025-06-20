package config

import (
	"encoding/json"
	"errors"
	"github.com/caffeines/notion-todo/models"
	"github.com/caffeines/notion-todo/service/files"
)

type credentialImpl struct {
	file files.File
}

var (
	credential Credential
)

func NewCredentialSvc(file files.File) Credential {
	if file == nil {
		panic("file storage not initialized")
	}
	if credential == nil {
		credential = &credentialImpl{
			file: file,
		}
	}
	return credential
}

func (c *credentialImpl) SetConfig(token string, databaseID string) error {
	if token == "" || databaseID == "" {
		return errors.New("token or databaseID cannot be empty")
	}
	if c.file == nil {
		return errors.New("file storage not initialized")
	}

	cfg := models.Config{
		Token:      token,
		DatabaseID: databaseID,
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = c.file.SaveFile(data)
	return err
}

func (c *credentialImpl) GetConfig() (*models.Config, error) {
	if c.file == nil {
		return nil, errors.New("file storage not initialized")
	}
	data, err := c.file.ReadFile()
	if err != nil {
		return nil, err
	}
	var cfg models.Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
