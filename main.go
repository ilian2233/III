package main

import (
	urlparser "III/url_parser"
	"fmt"
)

func main() {
	// t := tree.CreateExampleTree()

	// fmt.Printf("Zad 4:\n\n")
	// fmt.Printf("The lowest val: %d\n", t.GetLowestVal())
	// fmt.Printf("The highest val: %d", t.GetHighestVal())

	// fmt.Printf("\n\nZad 5:\n\n")

	// order := "Inorder"
	// if len(os.Args) > 1 {
	// 	order = os.Args[1]
	// }

	// if err := t.Traverse(order); err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Printf("\n\nZad 6:\n\n")
	// t.TreverseAndUpdate(10)
	// t.Traverse(order)

	// fmt.Printf("\n\nZad 7:\n\n")
	// fmt.Printf("Depth: %d", t.GetDepth())

	// fmt.Printf("\n\nZad 8:\n\n")
	// filePath := "./tree.txt"
	// tree, err := tree.CreateTreeFromFile(filePath)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Three loaded from file %v: \n", filePath)
	// tree.Traverse(order)

	//parser.Parse1()

	err := urlparser.Parse("https://pkg.go.dev/golang.org/x/text/unicode/rangetable#New")
	if err != nil {
		fmt.Println(err)
		return
	}
}
