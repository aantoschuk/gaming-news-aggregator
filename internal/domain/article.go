package domain

import "fmt"

// Article is an object have data about single article.
type Article struct {
	// link to the article
	Url string
	// title of the article
	Title string
}

// Stringer implementation
func (a Article) String() string {
	return fmt.Sprintf("title: %s\nurl: %s", a.Title, a.Url)
}
