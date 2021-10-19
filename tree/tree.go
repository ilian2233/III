package tree

import (
	"errors"
	"fmt"
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

func CreateExampleTree() tree {
	return tree{
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
