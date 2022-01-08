package inverted_index

import (
	"testing"
)

func TestAddWord(t *testing.T) {
	index := Index{}

	index.AddToIndex("asd", 1)

	if index["asd"][0] != 1 {
		t.FailNow()
	}
}
func TestGetAddedWord(t *testing.T) {
	index := Index{}
	word := "asd"
	position := 1

	index.AddToIndex(word, uint(position))

	occurrences := index.GetFromIndexByWord(word)
	if len(occurrences) > 1 || occurrences[0] != 1 {
		t.FailNow()
	}
}
func TestGetNonexistentWord(t *testing.T) {
	index := Index{}
	word := "asd"

	occurrences := index.GetFromIndexByWord(word)
	if len(occurrences) > 0 {
		t.FailNow()
	}
}
func TestGetWordByIndex(t *testing.T) {
	index := Index{}
	word := "asd"
	position := 1

	index.AddToIndex(word, uint(position))

	res, err := index.GetFromIndexByPosition(uint(position))
	if res != word || err != nil {
		t.FailNow()
	}
}
func TestGetWordByNonexistentIndex(t *testing.T) {
	index := Index{}
	position := 1

	_, err := index.GetFromIndexByPosition(uint(position))
	if err == nil {
		t.FailNow()
	}
}
