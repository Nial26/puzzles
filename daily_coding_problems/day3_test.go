package puzzles

import (
	"reflect"
	"testing"
)

/*
		    root
		/        \
	  left        right
	/     \        /  \
 left.left NULL  NULL NULL
  /   \
NULL  NULL
*/

func TestTreeSerialization(t *testing.T) {
	createdTree := tree{
		val: "root",
		left: &tree{
			val: "left",
			left: &tree{
				val: "left.left",
			},
		},
		right: &tree{
			val: "right",
		},
	}
	expected := []string{"root", "left", "left.left", "NULL", "NULL", "NULL", "right", "NULL", "NULL"}
	if got := createdTree.serialize(); !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected : %v, got :%v", expected, got)
	}
	expectedString := "root, left, left.left, NULL, NULL, NULL, right, NULL, NULL"
	if got := createdTree.serializeToString(); got != expectedString {
		t.Fatalf("expected : %v, got :%v", expected, got)
	}
}

func TestDeserialization(t *testing.T) {
	serializedTree := "root, left, left.left, NULL, NULL, NULL, right, NULL, NULL"
	deserializedArrayTree := deserializeToArray(serializedTree)
	tree := *deserializeToTree(&deserializedArrayTree)
	if got := tree.left.left.val; got != "left.left" {
		t.Fatalf("expected : %v, got : %v", "left.left", got)
	}
}
