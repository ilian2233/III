package html_parser

import (
	"fmt"
	"golang.org/x/text/unicode/rangetable"
	"strings"
	"unicode"
)

type terminalSymbol rune

//TODO: Test
func readTag(input, tag string) (string, string, error) {
	if !strings.HasPrefix(input, "<"+tag) {
		return "", "", fmt.Errorf("misformed start of tag")
	}
	input = strings.TrimPrefix(input, "<"+tag)

	if strings.HasSuffix(input, "/>") {
		input = strings.TrimSuffix(input, "/>")

		return input, "", nil
	}

	closingTagPos := strings.IndexAny(input, ">")
	if closingTagPos == -1 {
		return "", "", fmt.Errorf("unclosed open tag")
	}

	if !strings.HasSuffix(input, "</"+tag+">") {
		return "", "", fmt.Errorf("misformed closing tag")
	}
	input = strings.TrimSuffix(input, "</"+tag+">")

	return input[closingTagPos:], input[:closingTagPos], nil
}

func a_attributes(input string) ([]terminalSymbol, error) {
	input = strings.Trim(input, " ")

	inputArr := strings.Split(input, "=")

	if inputArr[0] != "href" && inputArr[0] != "target" {
		return nil, fmt.Errorf("missing 'target' or 'href' parameter")
	}
	return attr_value(inputArr[1])
}

func alpha(input string) ([]terminalSymbol, error) {
	if len(input) == 1 && unicode.IsLetter(rune(input[0])) {
		return []terminalSymbol{terminalSymbol(input[0])}, nil
	}

	return nil, fmt.Errorf("the input is not a letter")
}

func anchor(input string) ([]terminalSymbol, error) {
	attributes, body, err := readTag(input, "a")
	if err != nil {
		return nil, err
	}
	attributeSymbols, err := a_attributes(attributes)
	if err != nil {
		return nil, err
	}
	bodySymbols, err := html(body)
	if err != nil {
		return nil, err
	}

	return append(attributeSymbols, bodySymbols...), nil
}

func attr_value(input string) ([]terminalSymbol, error) {
	if input[0] == '"' && input[len(input)-1] == '"' {
		return urltext(strings.Trim(input, "\""))
	}

	return urltext(input)
}

func body(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "body")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func bold(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "b")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func digit(input string) ([]terminalSymbol, error) {
	if len(input) == 1 && unicode.IsDigit(rune(input[0])) {
		return []terminalSymbol{terminalSymbol(input[0])}, nil
	}

	return nil, fmt.Errorf("the input is not a letter")
}

func h1_tag(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "h1")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func h2_tag(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "h2")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func h3_tag(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "h3")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func h4_tag(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "h4")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func h5_tag(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "h5")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func h6_tag(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "h6")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func head(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "head")
	if err != nil {
		return nil, err
	}

	return head_content(body)
}

func head_content(input string) (symbols []terminalSymbol, err error) {
	if symbols, err = meta(input); err == nil {
		return symbols, nil
	}

	if symbols, err = title(input); err == nil {
		return symbols, nil
	}

	return nil, err
}

func heading(input string) (symbols []terminalSymbol, err error) {
	if symbols, err = h1_tag(input); err == nil {
		return symbols, nil
	}

	if symbols, err = h2_tag(input); err == nil {
		return symbols, nil
	}

	if symbols, err = h3_tag(input); err == nil {
		return symbols, nil
	}

	if symbols, err = h4_tag(input); err == nil {
		return symbols, nil
	}

	if symbols, err = h5_tag(input); err == nil {
		return symbols, nil
	}

	if symbols, err = h6_tag(input); err == nil {
		return symbols, nil
	}

	return nil, err
}

func html(input string) ([]terminalSymbol, error) {
	inputArr := strings.SplitN(input, " ", 2)

	itemSymbols, err := item(inputArr[0])
	if err != nil{
		return nil, err
	}

	htmlSymbols, err := html(inputArr[1])
	if err != nil{
		return nil, err
	}

	return append(itemSymbols,htmlSymbols...), nil
}

func html_document(input string) ([]terminalSymbol, error) {

	endOfHead := strings.IndexAny(input, "</head>")

	headSymbols, err := head(input[:endOfHead])
	if err != nil{
		return nil, err
	}

	bodySymbols, err := body(input[endOfHead:])
	if err != nil{
		return nil, err
	}

	return append(headSymbols, bodySymbols...), nil
}

func italic(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "i")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func item(input string) (symbols []terminalSymbol,err error) {
	if symbols, err = physical_style(input); err == nil {
		return symbols, nil
	}

	if symbols, err = prgf(input); err == nil {
		return symbols, nil
	}

	if symbols, err = plaintext(input); err == nil {
		return symbols, nil
	}

	if symbols, err = anchor(input); err == nil {
		return symbols, nil
	}

	if symbols, err = heading(input); err == nil {
		return symbols, nil
	}

	if symbols, err = table(input); err == nil {
		return symbols, nil
	}

	if symbols, err = list(input); err == nil {
		return symbols, nil
	}

	return nil, err
}

func itemlist(input string) ([]terminalSymbol, error) {
	items := strings.Split(input, "<li>")

	var symbols []terminalSymbol
	for _,v := range items {
		currentSymbols, err := oneitem("<li>"+v)
		if err != nil {
			return nil, err
		}

		symbols = append(symbols, currentSymbols...)
	}

	return symbols, nil
}

func list(input string) (symbols []terminalSymbol,err error) {
	if symbols, err := ulist(input); err == nil{
		return symbols, nil
	}

	if symbols, err := olist(input); err == nil{
		return symbols, nil
	}

	return nil, err
}

func meta(input string) ([]terminalSymbol, error) {
	attributes, _, err := readTag(input, "meta")
	if err != nil {
		return nil, err
	}

	attributeArr := strings.Split(attributes, "\"")

	if attributeArr[0] != "name = " && attributeArr[2] != "content = "{
		return nil, fmt.Errorf("required attributes not found")
	}

	nameSymbols, err := plaintext(attributeArr[1])
	if err != nil {
		return nil, err
	}

	textSymbols, err := plaintext(attributeArr[3])
	if err != nil {
		return nil, err
	}

	return append(nameSymbols, textSymbols...), nil
}

func olist(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "ol")
	if err != nil {
		return nil, err
	}

	return itemlist(body)
}

func oneitem(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "li")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func physical_style(input string) (symbols []terminalSymbol,err error) {
	if symbols, err := bold(input); err == nil{
		return symbols, nil
	}

	if symbols, err := italic(input); err == nil{
		return symbols, nil
	}

	return nil, err
}

func plaintext(input string) ([]terminalSymbol, error) {
	if strings.Contains(input, "<"){
		return nil, fmt.Errorf("should not contain <")
	}

	symbols := make([]terminalSymbol, len(input))

	for i,v := range input {
		symbols[i] = terminalSymbol(v)
	}

	return symbols, nil
}

func prgf(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "p")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func table(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "table")
	if err != nil {
		return nil, err
	}

	tableRows := strings.Split(body, "<tr>")

	var symbols []terminalSymbol

	for _,v := range tableRows{
		localSymbols, err := tablerow("<tr>"+v)
		if err != nil{
			return nil, err
		}

		symbols = append(symbols, localSymbols...)
	}

	return symbols, nil
}

func tablecell(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "td")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func tablerow(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "table")
	if err != nil {
		return nil, err
	}

	tableRows := strings.Split(body, "<td>")

	var symbols []terminalSymbol

	for _,v := range tableRows{
		localSymbols, err := tablecell("<td>"+v)
		if err != nil{
			return nil, err
		}

		symbols = append(symbols, localSymbols...)
	}

	return symbols, nil
}

func title(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "title")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func ulist(input string) ([]terminalSymbol, error) {
	_, body, err := readTag(input, "ul")
	if err != nil {
		return nil, err
	}

	return html(body)
}

func urltext(input string) (symbols []terminalSymbol, err error) {

	urlAllowedSymbols := rangetable.New('.', '/', ':', '?')

	if !unicode.IsOneOf([]*unicode.RangeTable{urlAllowedSymbols}, rune(input[0])) {
		return nil, fmt.Errorf("expected single safe character, recieved: %v", input)
	}

	for i,v := range input {
		currentSymbols, err := alpha(string(v))

		if i!=0 {

			if err != nil{
				return nil, err
			}

			currentSymbols, err = digit(string(v))
			if err != nil{
				return nil, err
			}

			if !unicode.IsOneOf([]*unicode.RangeTable{urlAllowedSymbols}, rune(v)) {
				return nil, fmt.Errorf("expected single safe character, recieved: %v", input)
			}else{
				currentSymbols = []terminalSymbol{terminalSymbol(v)}
			}
		}

		symbols = append(symbols, currentSymbols...)
	}

	return symbols, nil
}
