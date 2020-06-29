package main

import (
	"testing"
)

func TestGetBooks(t *testing.T) {
	books := Books{
		Book{
			Author:      "A1",
			Title:       "Golang Book",
			ISBN:        "ABCD",
			ReleaseDate: "01-01-2020",
		},
	}
	if len(books) == 0 {
		t.Errorf("Failed! %v", books)
	} else {
		t.Logf("Success! %v", books)
	}
}
