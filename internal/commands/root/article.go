package root

import "fmt"

type article struct {
	url   string
	title string
}

func (a article) String() string {
	return fmt.Sprintf("title: %s\nurl: %s", a.title, a.url)
}
