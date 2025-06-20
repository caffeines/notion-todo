package files

import (
	"fmt"
	"os"
	"os/exec"
)

const AppName = "notion-todo"

type fileImpl struct {
	fileName string
}

func NewFileService(fileName string) File {
	return &fileImpl{
		fileName: fileName,
	}
}

func (f *fileImpl) getBasePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil
}

func (f *fileImpl) createDirIfNotExists() error {
	basePath, err := f.getBasePath()
	if err != nil {
		return err
	}
	dirPath := fmt.Sprintf("%s/.%s", basePath, AppName)
	_, err = exec.Command("mkdir", "-p", dirPath).Output()
	if err != nil {
		return err
	}
	return nil
}

func (f *fileImpl) getPath() (string, error) {
	basePath, err := f.getBasePath()
	if err != nil {
		return "", err
	}
	path := fmt.Sprintf("%s/.%s/%s", basePath, AppName, f.fileName)
	return path, nil
}

func (f *fileImpl) SaveFile(data []byte) error {
	err := f.createDirIfNotExists()
	if err != nil {
		return err
	}
	path, err := f.getPath()
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0644)
	return err
}

func (f *fileImpl) ReadFile() ([]byte, error) {
	path, err := f.getPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
