package main

import (
	"III/three"
	"fmt"
	"os"
)

func main() {

	t := three.CreateExampleTree()

	fmt.Printf("Zad 4:\n\n")
	fmt.Printf("The lowest val: %d\n", t.GetLowestVal())
	fmt.Printf("The highest val: %d", t.GetHighestVal())

	fmt.Printf("\n\nZad 5:\n\n")

	order := "Inorder"
	if len(os.Args) > 1 {
		order = os.Args[1]
	}

	if err := t.Traverse(order); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("\n\nZad 6:\n\n")
	t.TreverseAndUpdate(10)
	t.Traverse("Inorder")

	fmt.Printf("\n\nZad 7:\n\n")
	fmt.Printf("Depth: %d", t.GetDepth())

}
