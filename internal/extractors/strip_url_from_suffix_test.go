package extractors

import (
	"testing"
)

func TestStripUrlFromSuffix(t *testing.T) {
	s, _ := stripUrlFromSuffix("https://www.ign.com/news/", "news")

	if *s != "https://www.ign.com" {
		t.Fatalf("test failed: want=%s got=%s", "https://www.ign.com", *s)
	}

	s, _ = stripUrlFromSuffix("https://www.ign.com/news", "news")
	if *s != "https://www.ign.com" {
		t.Fatalf("test failed: want=%s got=%s", "https://www.ign.com", *s)
	}

	s, _ = stripUrlFromSuffix("", "news")
	if s != nil {
		t.Fatalf("test failed: string should be nil, but got=%s", *s)
	}
}
