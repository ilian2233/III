package urlparser

import (
	"fmt"
	"golang.org/x/text/unicode/rangetable"
	"strings"
	"unicode"
)

type terminalSymbol rune

func url(input string) (val []terminalSymbol, err error) {
	if val, err = generic(input); err == nil {
		return val, nil
	}
	if val, err = httpsAddress(input); err == nil {
		return val, nil
	}
	if val, err = httpAddress(input); err == nil {
		return val, nil
	}
	if val, err = ftpAddress(input); err == nil {
		return val, nil
	}
	return nil, fmt.Errorf("%w\ninvalid URL", err)
}

func generic(input string) ([]terminalSymbol, error) {
	v := strings.Split(input, ":")
	if len(v) != 1 {
		return nil, fmt.Errorf("missing or too many occurences of ':' ")
	}

	v1, e1 := scheme(v[0])
	if e1 != nil {
		return nil, e1
	}

	v2, e2 := path(v[1])
	if e2 != nil {
		return nil, e2
	}

	return append(v1, v2...), nil
}

func scheme(input string) (val []terminalSymbol, err error) {
	return ialpha(input)
}

func httpsAddress(input string) ([]terminalSymbol, error) {
	if !strings.Contains(input, "https://") {
		return nil, fmt.Errorf("missing 'https://'")
	}

	input = strings.Replace(input, "https://", "", 1)

	arr := strings.SplitN(input, "/", 2)

	symbols, err := hostport(arr[0])
	if err != nil {
		return nil, err
	}

	terminalSymbols, err := path(arr[1])
	if err != nil {
		return nil, err
	}

	return append(symbols, terminalSymbols...), nil
}

func httpAddress(input string) ([]terminalSymbol, error) {
	if !strings.Contains(input, "http://") {
		return nil, fmt.Errorf("missing 'http://'")
	}

	input = strings.Replace(input, "http://", "", 1)

	arr := strings.SplitN(input, "/", 2)

	symbols, err := hostport(arr[0])
	if err != nil {
		return nil, err
	}

	terminalSymbols, err := path(arr[1])
	if err != nil {
		return nil, err
	}

	return append(symbols, terminalSymbols...), nil
}

func ftpAddress(input string) ([]terminalSymbol, error) {
	if !strings.Contains(input, "ftp://") {
		return nil, fmt.Errorf("missing 'ftp://'")
	}

	input = strings.Replace(input, "ftp://", "", 1)

	arr := strings.Split(input, "/")

	loginSymbols, err := login(arr[0])
	if err != nil {
		return nil, err
	}

	pathSymbols, err := path(arr[1])
	if err != nil {
		return nil, err
	}

	return append(loginSymbols, pathSymbols...), nil
}

func login(input string) (val []terminalSymbol, err error) {
	arr := strings.SplitN(input, "@", 2)

	val, err = hostport(arr[len(arr)-1])
	if err != nil {
		return nil, err
	}

	if len(arr) > 1 {
		userAndPassword := strings.SplitN(arr[0], ":", 2)

		userSymbols, err := user(userAndPassword[0])
		if err != nil {
			return nil, err
		}
		val = append(val, userSymbols...)

		if len(userAndPassword) > 1 {
			passwordSymbols, err := password(userAndPassword[1])
			if err != nil {
				return nil, err
			}
			val = append(val, passwordSymbols...)
		}
	}

	return val, nil
}

func hostport(input string) (val []terminalSymbol, err error) {
	arr := strings.SplitN(input, ":", 2)

	val, err = host(arr[0])
	if err != nil {
		return nil, err
	}

	if len(arr) > 1 {
		portSymbols, err := port(arr[1])
		if err != nil {
			return nil, err
		}
		val = append(val, portSymbols...)
	}

	return val, nil
}

func host(input string) (val []terminalSymbol, err error) {
	val, err = hostname(input)
	if err != nil {
		return hostnumber(input)
	}

	return val, nil
}

func hostname(input string) (val []terminalSymbol, err error) {
	arr := strings.SplitN(input, ".", 2)

	val, err = ialpha(arr[0])
	if err != nil {
		return nil, err
	}

	if len(arr) > 1 {
		hostnameSymbols, err := hostname(arr[1])
		if err != nil {
			return nil, err
		}
		val = append(val, hostnameSymbols...)
	}

	return val, nil
}

func hostnumber(input string) (val []terminalSymbol, err error) {
	arr := strings.SplitN(input, ".", 4)

	for _, v := range arr {
		digitSymbols, err := digits(v)
		if err != nil {
			return nil, err
		}

		val = append(val, digitSymbols...)
	}

	return val, nil
}

func port(input string) (val []terminalSymbol, err error) {
	return digits(input)
}

func path(input string) (val []terminalSymbol, err error) {
	if input == "" {
		return val, nil
	}

	arr := strings.SplitN(input, "/", 2)
	val, err = xpalphas(arr[0])
	if err != nil {
		return nil, err
	}

	if len(arr) > 1 {
		pathSymbols, err := path(arr[1])
		if err != nil {
			return nil, err
		}
		val = append(val, pathSymbols...)
	}

	return val, nil
}

func user(input string) (val []terminalSymbol, err error) {
	return xalphas(input)
}

func password(input string) (val []terminalSymbol, err error) {
	return xalphas(input)
}

func xalpha(input string) (val []terminalSymbol, err error) {
	if val, err = alpha(input); err == nil {
		return val, nil
	}
	if val, err = digit(input); err == nil {
		return val, nil
	}
	if val, err = safe(input); err == nil {
		return val, nil
	}
	if val, err = extra(input); err == nil {
		return val, nil
	}
	if val, err = escape(input); err == nil {
		return val, nil
	}
	return nil, err
}

func xalphas(input string) (val []terminalSymbol, err error) {
	for _, v := range input {
		xalphaSymbol, err := xalpha(string(v))
		if err != nil {
			return nil, err
		}
		val = append(val, xalphaSymbol...)
	}

	return val, nil
}

func xpalpha(input string) (val []terminalSymbol, err error) {
	if input == "+" {
		return append(val, terminalSymbol('+')), nil
	}

	return xalpha(input)
}

func xpalphas(input string) (val []terminalSymbol, err error) {
	val, err = xpalpha(string(input[0]))
	if err != nil {
		return nil, err
	}

	if len(input) > 1 {
		xalphasSymbol, err := xpalphas(input[1:])
		if err != nil {
			return nil, err
		}
		val = append(val, xalphasSymbol...)
	}

	return val, nil
}

func ialpha(input string) (val []terminalSymbol, err error) {
	val, err = alpha(string(input[0]))
	if err != nil {
		return nil, err
	}

	if len(input) > 1 {
		xalphasSymbol, err := xalphas(input[1:])
		if err != nil {
			return nil, err
		}
		val = append(val, xalphasSymbol...)
	}

	return val, nil
}

func alpha(input string) (val []terminalSymbol, err error) {
	if len(input) > 1 || !unicode.IsLetter(rune(input[0])) {
		return nil, fmt.Errorf("expected single letter, recieved: %v", input)
	}

	return append(val, terminalSymbol(input[0])), nil
}

func digit(input string) (val []terminalSymbol, err error) {
	if len(input) > 1 || !unicode.IsDigit(rune(input[0])) {
		return nil, fmt.Errorf("expected single digit, recieved: %v", input)
	}

	return append(val, terminalSymbol(input[0])), nil
}

func safe(input string) ([]terminalSymbol, error) {
	safeRangeTable := rangetable.New('$', '-', '\\', '@', '.', '&')

	if !unicode.IsOneOf([]*unicode.RangeTable{safeRangeTable}, rune(input[0])) {
		return nil, fmt.Errorf("expected single safe character, recieved: %v", input)
	}

	return []terminalSymbol{terminalSymbol(input[0])}, nil
}

func extra(input string) ([]terminalSymbol, error) {
	extraRangeTable := rangetable.New('!', '*', '"', '\'', '(', ')', ':', ';', ',', '#', ' ')

	if !unicode.IsOneOf([]*unicode.RangeTable{extraRangeTable}, rune(input[0])) {
		return nil, fmt.Errorf("expected single extra character, recieved: %v", input)
	}

	return []terminalSymbol{terminalSymbol(input[0])}, nil
}

func escape(input string) ([]terminalSymbol, error) {
	if !strings.Contains(input, "\\%") {
		return nil, fmt.Errorf("missing escape character, recieved: %v", input)
	}

	input = strings.Replace(input, "\\%", "", 1)

	if len(input) != 2 {
		return nil, fmt.Errorf("expected two xeh characters, recieved: %v", input)
	}

	hexSymbol1, err := hex(string(input[0]))
	if err != nil {
		return nil, err
	}

	hexSymbol2, err := hex(string(input[1]))
	if err != nil {
		return nil, err
	}

	return append(hexSymbol1, hexSymbol2...), nil
}

func hex(input string) ([]terminalSymbol, error) {
	if len(input) > 1 {
		return nil, fmt.Errorf("expected single hex, recieved: %v", input)
	}

	if unicode.IsDigit(rune(input[0])) {
		return digit(input)
	}

	hexRangeTable := rangetable.New('a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F')

	if !unicode.IsOneOf([]*unicode.RangeTable{hexRangeTable}, rune(input[0])) {
		return nil, fmt.Errorf("expected single hex, recieved: %v", input)
	}

	return []terminalSymbol{terminalSymbol(input[0])}, nil
}

func digits(input string) (val []terminalSymbol, err error) {
	val, err = digit(string(input[0]))
	if err != nil {
		return nil, err
	}

	if len(input) > 1 {
		digitsSymbols, err := digits(input[1:])
		if err != nil {
			return nil, err
		}
		val = append(val, digitsSymbols...)
	}

	return val, nil
}

func Parse(input string) error {
	parsed, err := url(input)
	if err != nil {
		return err
	}
	fmt.Println(parsed)
	return nil
}
