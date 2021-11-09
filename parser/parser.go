package parser

import (
	"fmt"
	"math"
	"unicode"
)

type word string

func (p parser) tokenize() ([]interface{}, error) {
	var tokens []interface{}

	for i := 0; i < len(p.sentence); {
		switch {
		case p.sentence[i] == '"' || p.sentence[i] == ';' || p.sentence[i] == '.':

			tokens = append(tokens, p.sentence[i])
			i++
		case unicode.IsDigit(rune(p.sentence[i])):

			number := uint(0)

			for unicode.IsDigit(rune(p.sentence[i])) {
				number *= 10
				number += uint(p.sentence[i])
				i++
			}

			if number > math.MaxInt32 {
				return nil, fmt.Errorf("int const too large")
			}

			tokens = append(tokens, number)
		case unicode.IsLetter(rune(p.sentence[i])):

			for unicode.IsDigit(rune(p.sentence[i])) || unicode.IsLetter(rune(p.sentence[i])) {
				i++
			}

			tokens = append(tokens, word(""))
		default:

			return nil, fmt.Errorf("unknown chacter")
		}
	}

	return tokens, nil
}

type parser struct {
	sentence string
}

func CreateParser(input string) parser {
	return parser{sentence: input}
}

func (p parser) IsValid() bool {
	_, err := p.tokenize()
	return err == nil
}

func Parse() {
	fmt.Println(CreateParser("123;631;\"hello\".5678;\"course1\".\"end\".").IsValid())
}
