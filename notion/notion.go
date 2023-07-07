package notion

type Notion interface {
	AddPage(title string) error
}
