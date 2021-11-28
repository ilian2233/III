package html_parser

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type terminalSymbol rune

func readTag(input, tag string)(string, string, error){
	//https://regex101.com/r/O2zjdn/1
	if len(regexp.
		MustCompile(`<(“[^”]*”|'[^’]*’|[^'”>])+>(...)*<\\/(“[^”]*”|'[^’]*’|[^'”>])+>|<(“[^”]*”|'[^’]*’|[^'”>])+\\/>`).
		FindAllString(input, -1)) > 1 {
		return "","",fmt.Errorf("tag syntax not found")
	}

	input = strings.TrimLeft(input, fmt.Sprintf("<%s", tag))
	endOfOpeningTag := strings.Index(input, ">")
	if endOfOpeningTag == -1 {
		return "","",fmt.Errorf("misformed tag")
	}

	tagAttributes := input[:endOfOpeningTag-1]

	closingTag := strings.LastIndex(input, "</"+tag+">")
	if endOfOpeningTag == -1 {
		return "","",fmt.Errorf("misformed tag")
	}

	tagBody := input[endOfOpeningTag:closingTag]

	return tagAttributes, tagBody, nil
}

func a_attributes(input string) ([]terminalSymbol, error){
	input = strings.Trim(input, " ")

	inputArr := strings.Split(input, "=")

	if inputArr[0] != "href" && inputArr[0] != "target"{
		return nil, fmt.Errorf("missing 'target' or 'href' parameter")
	}
	return attr_value(inputArr[1])
}

func alpha(input string) ([]terminalSymbol, error){
	if len(input)==1 && unicode.IsLetter(rune(input[0])){
		return []terminalSymbol{terminalSymbol(input[0])}, nil
	}

	return nil, fmt.Errorf("the input is not a letter")
}

func anchor(input string) ([]terminalSymbol, error){
	attributes, body, err := readTag(input, "a")
	if err != nil{
		return nil, err
	}
	attributeSymbols, err := a_attributes(attributes)
	if err != nil{
		return nil, err
	}
	bodySymbols, err := html(body)
	if err != nil{
		return nil, err
	}

	return append(attributeSymbols, bodySymbols...), nil
}

func attr_value(input string) ([]terminalSymbol, error){
	if input[0] == '"' && input[len(input)-1] == '"' {
		return urltext(strings.Trim(input, "\""))
	}

 	return urltext(input)
}

func body(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "body")
	if err != nil{
		return nil, err
	}

	return html(body)
}

func bold(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "b")
	if err != nil{
		return nil, err
	}

	return html(body)
}

func digit(input string) ([]terminalSymbol, error){
	if len(input)==1 && unicode.IsDigit(rune(input[0])){
		return []terminalSymbol{terminalSymbol(input[0])}, nil
	}

	return nil, fmt.Errorf("the input is not a letter")
}

func h1_tag(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "h1")
	if err != nil{
		return nil, err
	}

	return html(body)
}

func h2_tag(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "h2")
	if err != nil{
		return nil, err
	}

	return html(body)}

func h3_tag(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "h3")
	if err != nil{
		return nil, err
	}

	return html(body)}

func h4_tag(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "h4")
	if err != nil{
		return nil, err
	}

	return html(body)}

func h5_tag(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "h5")
	if err != nil{
		return nil, err
	}

	return html(body)}

func h6_tag(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "h6")
	if err != nil{
		return nil, err
	}

	return html(body)
}

func head(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "head")
	if err != nil{
		return nil, err
	}

	return head_content(body)
}

func head_content(input string) (symbols []terminalSymbol, err error){
	if symbols, err = meta(input); err == nil{
		return symbols, nil
	}

	if symbols, err = title(input); err == nil{
		return symbols, nil
	}

	return nil, err
}

func heading(input string) (symbols []terminalSymbol, err error){
	if symbols, err = h1_tag(input); err == nil{
		return symbols, nil
	}

	if symbols, err = h2_tag(input); err == nil{
		return symbols, nil
	}

	if symbols, err = h3_tag(input); err == nil{
		return symbols, nil
	}

	if symbols, err = h4_tag(input); err == nil{
		return symbols, nil
	}

	if symbols, err = h5_tag(input); err == nil{
		return symbols, nil
	}

	if symbols, err = h6_tag(input); err == nil{
		return symbols, nil
	}

	return nil, err
}

func html(input string) ([]terminalSymbol, error){
	return nil, nil
}

func html_document(input string) ([]terminalSymbol, error){
	return nil, nil
}

func italic(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "i")
	if err != nil{
		return nil, err
	}

	return html(body)
}

func item(input string) ([]terminalSymbol, error){
	return nil, nil
}

func itemlist(input string) ([]terminalSymbol, error){
	return nil, nil
}

func list(input string) ([]terminalSymbol, error){
	return nil, nil
}

func meta(input string) ([]terminalSymbol, error){
	return nil, nil
}

func olist(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "ol")
	if err != nil{
		return nil, err
	}

	return itemlist(body)
}

func oneitem(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "li")
	if err != nil{
		return nil, err
	}

	return html(body)
}

func physical_style(input string) ([]terminalSymbol, error){
	return nil, nil
}

func plaintext(input string) ([]terminalSymbol, error){
	return nil, nil
}

func prgf(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "p")
	if err != nil{
		return nil, err
	}

	return html(body)}

func table(input string) ([]terminalSymbol, error){
	return nil, nil
}

func tablecell(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "td")
	if err != nil{
		return nil, err
	}

	return html(body)}

func tablerow(input string) ([]terminalSymbol, error){
	return nil, nil
}

func title(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "title")
	if err != nil{
		return nil, err
	}

	return html(body)}

func ulist(input string) ([]terminalSymbol, error){
	_, body, err := readTag(input, "ul")
	if err != nil{
		return nil, err
	}

	return html(body)
}

func urltext(input string) ([]terminalSymbol, error){
	return nil, nil
}

