package service

type File interface {
	SaveFile(data []byte) error
	ReadFile() ([]byte, error)
}
