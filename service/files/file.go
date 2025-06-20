package files

type File interface {
	SaveFile(data []byte) error
	ReadFile() ([]byte, error)
}
