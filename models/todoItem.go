package models

type Text struct {
	Content string `json:"content"`
}

type TextTitle struct {
	Text Text `json:"text"`
}

type Title struct {
	Titles []TextTitle `json:"title"`
}
type Select struct {
	Name string `json:"name"`
}
type Status struct {
	Select Select `json:"select"`
}

type ItemData struct {
	ItemName   Title  `json:"Title"`
	ItemStatus Status `json:"Status"`
}

type Properties struct {
	Item ItemData `json:"properties"`
}

// NewProperties returns a new Properties
func NewProperties() ItemData {
	return ItemData{
		ItemName: Title{
			Titles: []TextTitle{
				{
					Text: Text{
						Content: "Hello World",
					},
				},
			},
		},
		ItemStatus: Status{
			Select: Select{
				Name: "Todo",
			},
		},
	}
}
