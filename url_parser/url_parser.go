package url_parser

import (
	"fmt"
	"strings"
)

type terminalSymbol struct{}

func url(input string) (val []terminalSymbol, err error) {
	if val, err = generic(input); err == nil {
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
	if len(v) == 1 {
		return nil, fmt.Errorf("missing ':' ")
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

func httpAddress(input string) ([]terminalSymbol, error) {
	return nil, fmt.Errorf("not implemented")
}

func ftpAddress(input string) ([]terminalSymbol, error) {
	return nil, fmt.Errorf("not implemented")
}

func login(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func hostport(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func host(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func hostname(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func hostnumber(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func port(input string) (val []terminalSymbol, err error) {
	return digits(input)
}

func path(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func user(input string) (val []terminalSymbol, err error) {
	return xalphas(input)
}

func password(input string) (val []terminalSymbol, err error) {
	return xalphas(input)
}

func xalpha(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func xalphas(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func xpalpha(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func xpalphas(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func ialpha(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func alpha(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func digit(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func safe(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func extra(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func escape(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func hex(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func digits(input string) (val []terminalSymbol, err error) {
	return nil, fmt.Errorf("not implemented")
}

func Parse(input string) {
	parsed, err := url(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsed)
}
