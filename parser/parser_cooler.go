package main

import (
	"fmt"
	"math"
	"unicode"
)

type token interface {
	AddToToken(ch rune) error
	isComplete() bool
}

type decimalToken struct {
	val      uint
	complete bool
}

func (dt decimalToken) AddToToken(ch rune) error {

	if ch == ';' {
		dt.complete = true
	}

	if !unicode.IsDigit(ch) {
		return fmt.Errorf("rune not digit")
	}

	dt.val *= 10
	dt.val += uint(ch)

	if dt.val > math.MaxInt32 {
		return fmt.Errorf("int const too large")
	}

	return nil
}

func (dt decimalToken) isComplete() bool {
	return dt.complete
}

type stringToken struct {
	val      string
	complete bool
}

func (st stringToken) AddToToken(ch rune) error {

	if unicode.IsDigit(ch) || unicode.IsLetter(ch) || ch == '"' {
		st.val += string(ch)
		return nil
	} else if ch == '.' {
		if st.val[len(st.val)-1] == '"' {
			st.complete = true
			return nil
		} else {
			return fmt.Errorf("word seems incomplete")
		}
	}
	return fmt.Errorf("this symbol cant be added to word")
}

func (st stringToken) isComplete() bool {
	return st.complete
}

type tokenCreator interface {
	AddSymbol(ch string) error
}

type coolerParser struct {
	tokens []token
}

func createNewToken(ch rune) (token, error) {
	switch {
	case unicode.IsDigit(ch):

		dt := decimalToken{}

		if err := dt.AddToToken(ch); err != nil {
			return nil, err
		}

		return dt, nil
	case ch == '"':

		st := stringToken{}

		if err := st.AddToToken(ch); err != nil {
			return nil, err
		}

		return st, nil
	default:
		return nil, fmt.Errorf("no token starts with that character")
	}
}

func (p coolerParser) AddSymbol(ch rune) error {
	last := len(p.tokens) - 1

	//Checks if last token is complete or is nil
	if p.tokens == nil || p.tokens[last] == nil || p.tokens[last].isComplete() {

		newToken, err := createNewToken(ch)
		if err != nil {
			return err
		}
		p.tokens = append(p.tokens, newToken)
	} else {

		err := p.tokens[last].AddToToken(ch)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	p := coolerParser{}

	for _, v := range "123;631;\"hello\".5678;\"course1\".\"end\"." {
		err := p.AddSymbol(v)
		if err != nil {
			return
		}
	}

	for _, v := range p.tokens {
		fmt.Println(v)
	}

}
