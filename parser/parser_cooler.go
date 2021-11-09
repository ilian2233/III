package parser

import (
	"fmt"
	"math"
	"strings"
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

func (dt *decimalToken) AddToToken(ch rune) error {
	if ch == ';' {
		dt.complete = true
	} else if !unicode.IsDigit(ch) {
		return fmt.Errorf("rune not digit")
	}

	dt.val *= 10
	dt.val += uint(ch - '0')

	if dt.val > math.MaxInt32 {
		return fmt.Errorf("int const too large")
	}

	return nil
}

func (dt decimalToken) isComplete() bool {
	return dt.complete
}

type stringToken struct {
	val      *strings.Builder
	complete bool
}

func (st *stringToken) AddToToken(ch rune) error {
	if unicode.IsDigit(ch) || unicode.IsLetter(ch) || ch == '"' {
		if _, err := st.val.WriteRune(ch); err != nil {
			return err
		}
		return nil
	} else if ch == '.' {
		if st.val.String()[st.val.Len()-1] == '"' {
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

type coolerParser []token

func createNewToken(ch rune) (token, error) {
	switch {
	case unicode.IsNumber(ch) || ch == ';':

		dt := decimalToken{}

		if err := dt.AddToToken(ch); err != nil {
			return nil, err
		}

		return &dt, nil
	case unicode.IsLetter(ch) || ch == '"' || ch == '.':

		st := stringToken{val: &strings.Builder{}}

		if err := st.AddToToken(ch); err != nil {
			return nil, err
		}

		return &st, nil
	default:
		return nil, fmt.Errorf("no token starts with that character")
	}
}

func (p *coolerParser) AddSymbol(ch rune) error {
	parser := *p
	last := len(parser) - 1

	//Checks if last token is complete or is nil
	if len(parser) < 1 || parser[last] == nil || parser[last].isComplete() {
		newToken, err := createNewToken(ch)
		if err != nil {
			return err
		}
		parser = append(parser, newToken)
		*p = parser
	} else {
		err := parser[last].AddToToken(ch)
		if err != nil {
			return err
		}
	}
	return nil
}

func Parse1() {
	p := coolerParser{}

	a := "123;631;\"hello\".5678;\"course1\".\"end\"."

	for _, v := range a {
		err := p.AddSymbol(v)
		if err != nil {
			fmt.Println(err)
		}
	}
	for _, v := range p {
		fmt.Println(v)
	}

}
