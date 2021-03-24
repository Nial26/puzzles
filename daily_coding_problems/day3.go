package puzzles

import (
	"strings"
)

/*
Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

For example, given the following Node class

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
The following test should pass:

node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'
*/

type tree struct {
	val   string
	left  *tree
	right *tree
}

func (t *tree) serialize() []string {
	if t == nil {
		return []string{"NULL"}
	}
	leftSide := t.left.serialize()
	rightSide := t.right.serialize()
	var result []string
	result = append(result, t.val)
	result = append(result, leftSide...)
	result = append(result, rightSide...)
	return result
}

func (t *tree) serializeToString() string {
	s := t.serialize()
	return strings.Join(s, ", ")
}

func deserializeToArray(treeString string) []string {
	return strings.Split(treeString, ", ")
}

func deserializeToTree(serializedTree *[]string) *tree {
	k := *serializedTree
	if len(k) == 0 {
		return nil
	}

	val := k[0]
	*serializedTree = k[1:]
	if val == "NULL" {
		return nil
	}

	t := tree{val: val}
	t.left = deserializeToTree(serializedTree)
	t.right = deserializeToTree(serializedTree)
	return &t
}

// This went on some random direction, initially I was thinking I would serialize it with the help of brackets and spent
// considerable amount of time, figuring out how to match brackets using regex and then trying to write a parantheses matcher etc.
// I knew I had to do something recursive, but I wasn't able to think of how to break down the recursive function to left and right with a value
// i.e, 2 recursive functions
// Eventually ended up checking the solution, and the answer was just go mutate the original list every time you consume a value, that's not very
// satisfying, but it works 🤷‍♀️
// Ended up implementing it in python first, and then porting over to go
