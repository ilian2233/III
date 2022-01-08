package inverted_index

import "fmt"

type Index map[string][]uint

func (i Index) AddToIndex(word string, position uint) {
	for _, v := range i[word] {
		if v == position {
			return
		}
	}

	i[word] = append(i[word], position)
}

func (i Index) GetFromIndexByWord(word string) []uint {
	return i[word]
}

func (i Index) GetFromIndexByPosition(position uint) (string, error) {
	for word, positions := range i {
		for _, currentPosition := range positions {
			if currentPosition == position {
				return word, nil
			}
		}
	}

	return "", fmt.Errorf("word not found")
}
