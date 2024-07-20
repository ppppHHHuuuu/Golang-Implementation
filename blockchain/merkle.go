package blockchain

import "fmt"

type node struct {
	left, right *node
	key         int
	value       int
}

// Bst Why use pointer here
/*
	1. Reuse value but not creating a copy of object => reduce memory usage
	2. pass by reference
	3. Dynamically allocation
*/
type Bst struct {
	head *node
}

var (
	maxim = 10
	size  = 0
	arr   []*Bst
)

/*
 * this function finds the given key in the Binary Tree
 * returns the node containing the key
 * returns NULL in case key is not present
 */
func find(tree *node, key int) (node *node) {
	if tree == nil {
		return nil
	}
	if tree.key == key {
		return tree
	} else if tree.left.key < key {
		return find(tree.left, key)
	} else {
		return find(tree.right, key)
	}
}
func hash(key int) int {
	return key % maxim
}
func displayTree(tree *node) {
	if tree == nil {
		return
	}
	fmt.Println(tree.key, " ", tree.value)
	if tree.left != nil {
		displayTree(tree.left)
	}
	if tree.right != nil {
		displayTree(tree.right)
	}
}

func Init() {
	i := 0
	arr = make([]*Bst, maxim)
	for i < maxim {
		fmt.Println("Init")
		arr[i].head = nil
		i++
	}
}
func Display() {
	for i := 0; i < maxim; i++ {
		tree := arr[i].head
		if tree == nil {
			fmt.Printf("arr[%v] has no %v\n", tree)
		} else {
			fmt.Printf("arr[%v] has %v\n", tree)
			displayTree(tree)
		}
	}
}

func SizeOfHashTree() int {
	return size
}

func InsertElement(tree *node, item *node) {
	if item.key < tree.key {
		if tree.left == nil {
			tree.left = item
		} else {
			InsertElement(tree.left, item)
		}
	} else {
		if tree.right == nil {
			tree.right = item
		} else {
			InsertElement(tree.right, item)
		}
	}
}

func Add(key int, value int) {
	index := hash(key)
	tree := arr[index].head
	new_item := &node{
		key:   key,
		value: value,
	}
	if tree == nil {
		fmt.Printf("Inserting %v and %v\n", key, value)
		arr[index].head = new_item
	} else {
		temp := find(tree, key)
		if temp == nil {
			fmt.Println("No key available")
			fmt.Printf("Inserting %v and %v\n", key, value)
		} else {
			fmt.Println("Update key ")
			temp.value = value
		}
	}
}

func Delete(key int) {
	index := hash(key)
	tree := arr[index].head
	if tree == nil {
		fmt.Printf("Key %v doesn't exist", key)
		return
	} else {

	}
}
