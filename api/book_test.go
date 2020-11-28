package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "123456789"}
	json := book.ToJSON()
	assert.Equal(t, `{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"123456789"}`, string(json), "Book JSON marshalling wrong")
	// Book struct에서 json:"소문자"로 변경 후에는 아래 케이스가 에러
	// assert.Equal(t, `{"Title":"Cloud Native Go","Author":"M.-L. Reimer","ISBN":"123456789"}`, string(json), "Book JSON marshalling wrong")

}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"Title":"Cloud Native Go","Author":"M.-L. Reimer","ISBN":"123456789"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "123456789"}, book, "Book JSON unmarshalling wrong")
}
