package three

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type three struct {
	value int
	left  *three
	right *three
}

func (t three) GetLowestVal() int {

	branch := &t

	for branch.left != nil {
		branch = branch.left
	}

	return branch.value
}

func (t three) GetHighestVal() int {

	branch := &t

	for branch.right != nil {
		branch = branch.right
	}

	return branch.value
}

func (t three) Traverse(mode string) error {

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

func (t three) treverseInorder() {

	if t.left != nil {
		t.left.treverseInorder()
	}

	fmt.Printf("%d,", t.value)

	if t.right != nil {
		t.right.treverseInorder()
	}
}

func (t three) treversePreorder() {

	fmt.Printf("%d,", t.value)

	if t.left != nil {
		t.left.treversePreorder()
	}

	if t.right != nil {
		t.right.treversePreorder()
	}
}

func (t three) treversePostorder() {

	if t.left != nil {
		t.left.treversePostorder()
	}

	if t.right != nil {
		t.right.treversePostorder()
	}

	fmt.Printf("%d,", t.value)
}

func (t *three) TreverseAndUpdate(val int) {

	if t.left != nil {
		t.left.TreverseAndUpdate(val)
	}

	if t.right != nil {
		t.right.TreverseAndUpdate(val)
	}

	t.value = t.value + val
}

func (t three) findDepth(depth int) int {

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

func (t three) GetDepth() int {
	return t.findDepth(0)
}

func CreateExampleThree() *three {
	return &three{
		value: 9,
		left: &three{
			value: 6,
			left: &three{
				value: 4,
			},
			right: &three{
				value: 7,
			},
		},
		right: &three{
			value: 16,
			left: &three{
				value: 10,
			},
			right: &three{
				value: 19,
				right: &three{
					value: 23,
				},
			},
		},
	}
}

func (t *three) add(value int) *three {

	//Handles nil pointer case
	if t == nil {
		return &three{value: value}
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

func CreateThreeFromFile(filePath string) (*three, error) {

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	values := strings.Split(string(fileContent), ",")

	three := &three{}

	for _, v := range values {

		convertedValue, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		three.add(convertedValue)
	}

	return three, nil
}
