package tree

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type tree struct {
	value int
	left  *tree
	right *tree
}

func (t tree) GetLowestVal() int {

	branch := &t

	for branch.left != nil {
		branch = branch.left
	}

	return branch.value
}

func (t tree) GetHighestVal() int {

	branch := &t

	for branch.right != nil {
		branch = branch.right
	}

	return branch.value
}

func (t tree) Traverse(mode string) error {

	fmt.Print(mode + ": ")

	switch mode {
	case "Inorder":
		{
			t.treverseInorder()
			return nil
		}
	case "Preorder":
		{
			t.treversePreorder()
			return nil
		}
	case "Postorder":
		{
			t.treversePostorder()
			return nil
		}
	default:
		return errors.New("No such mode found")
	}
}

func (t tree) treverseInorder() {

	if t.left != nil {
		t.left.treverseInorder()
	}

	fmt.Printf("%d,", t.value)

	if t.right != nil {
		t.right.treverseInorder()
	}
}

func (t tree) treversePreorder() {

	fmt.Printf("%d,", t.value)

	if t.left != nil {
		t.left.treversePreorder()
	}

	if t.right != nil {
		t.right.treversePreorder()
	}
}

func (t tree) treversePostorder() {

	if t.left != nil {
		t.left.treversePostorder()
	}

	if t.right != nil {
		t.right.treversePostorder()
	}

	fmt.Printf("%d,", t.value)
}

func (t *tree) TreverseAndUpdate(val int) {

	if t.left != nil {
		t.left.TreverseAndUpdate(val)
	}

	if t.right != nil {
		t.right.TreverseAndUpdate(val)
	}

	t.value = t.value + val
}

func (t tree) findDepth(depth int) int {

	leftDepth, rightDepth := depth, depth

	if t.left != nil {
		leftDepth = t.left.findDepth(depth + 1)
	}

	if t.right != nil {
		rightDepth = t.right.findDepth(depth + 1)
	}

	if leftDepth > rightDepth {
		return leftDepth
	} else {
		return rightDepth
	}
}

func (t tree) GetDepth() int {
	return t.findDepth(0)
}

func CreateExampleTree() *tree {
	return &tree{
		value: 9,
		left: &tree{
			value: 6,
			left: &tree{
				value: 4,
			},
			right: &tree{
				value: 7,
			},
		},
		right: &tree{
			value: 16,
			left: &tree{
				value: 10,
			},
			right: &tree{
				value: 19,
				right: &tree{
					value: 23,
				},
			},
		},
	}
}

func (t *tree) add(value int) *tree {

	//Handles nil pointer case
	if t == nil {
		return &tree{value: value}
	}

	//Handles empty three case
	if t.value == 0 && t.left == nil && t.right == nil {
		t.value = value
		return t
	}

	//Checks if value is smaller or bigger and adds it in the left or right recursively
	if t.value > value {
		t.left = t.left.add(value)
	} else if t.value < value {
		t.right = t.right.add(value)
	}

	//If value is equal doesn't add it
	return t
}

func CreateTreeFromFile(filePath string) (*tree, error) {

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	values := strings.Split(string(fileContent), ",")

	tree := &tree{}

	for _, v := range values {

		convertedValue, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		tree.add(convertedValue)
	}

	return tree, nil
}
