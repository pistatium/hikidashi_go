package entities

type Item struct {
	Path string
	Value string
	ContentType string
}

func NewItem(path string, value string) Item {
	return Item{
		Path: path,
		Value: value,
		ContentType: "text/plain",
	}
}
