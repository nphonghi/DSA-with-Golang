package main

import "fmt"

type Node struct {
	key int
	parent, left, right *Node
}

type BST struct {
	root *Node
}

func (tree *BST) SearchNode(key int) (*Node, bool) {
	if tree.root == nil {
		return nil, false
	}
	currentNode := tree.root
	for currentNode != nil {
		if key == currentNode.key {
			return currentNode, true
		} else if key >= currentNode.key {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}
	}
	return nil, false
}

func (tree *BST) FindMin() *Node {
	if tree.root == nil {
		return nil
	}
	currentNode := tree.root
	for currentNode.left != nil {
		currentNode = currentNode.left
	}
	return currentNode
}

func (tree *BST) findMin(node *Node) *Node {
	currentNode := node
	for currentNode.left != nil {
		currentNode = currentNode.left
	}
	return currentNode
}

func (tree *BST) FindMax() *Node {
	if tree.root == nil {
		return nil
	}
	currentNode := tree.root
	for currentNode.right != nil {
		currentNode = currentNode.right
	}
	return currentNode
}

func (tree *BST) findMax(node *Node) *Node {
	currentNode := node
	for currentNode.right != nil {
		currentNode = currentNode.right
	}
	return currentNode
}

func (tree *BST) InorderTraversal(subtree *Node, callback func(int)) {
	if subtree.left != nil {
		tree.InorderTraversal(subtree.left, callback)
	}
	callback(subtree.key)
	if subtree.right != nil {
		tree.InorderTraversal(subtree.right, callback)
	}
}

func (tree *BST) PreorderTraversal(subtree *Node, callback func(int)) {
	callback(subtree.key)
	if subtree.left != nil {
		tree.PreorderTraversal(subtree.left, callback)
	}
	if subtree.right != nil {
		tree.PreorderTraversal(subtree.right, callback)
	}
}

func (tree *BST) PostorderTraversal(subtree *Node, callback func(int)) {
	if subtree.left != nil {
		tree.PostorderTraversal(subtree.left, callback)
	}
	if subtree.right != nil {
		tree.PostorderTraversal(subtree.right, callback)
	}
	callback(subtree.key)
}

func (tree *BST) Predecessor(node *Node) *Node {
	if node.left != nil {
		return tree.findMax(node.left)
	}
	parent := node.parent
	for parent != nil && node == parent.left {
		node = parent
		parent = parent.parent
	}
	return parent
}

func (tree *BST) Successor(node *Node) *Node {
	if node.right != nil {
		return tree.findMin(node.right)
	}
	parent := node.parent
	for parent != nil && node == parent.right {
		node = parent
		parent = parent.parent
	}
	return parent
}

func (tree *BST) InsertNode(key int) {
	if tree.root == nil {
		tree.root = &Node{key: key, parent: nil, left: nil, right: nil}
		return
	}
	currentNode := tree.root
	for {
		if key >= currentNode.key {
			if currentNode.right == nil {
				currentNode.right = &Node{key: key, parent: currentNode, left: nil, right: nil}
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = &Node{key: key, parent: currentNode, left: nil, right: nil}
				return
			}
			currentNode = currentNode.left
		}
	}
}

func (tree *BST) Transplant(u, v *Node) {
	if u.parent == nil {
		tree.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (tree *BST) DeleteNode(key int) {
	node, found := tree.SearchNode(key)
	if !found {
		fmt.Printf("%d không có trong cây", key)
		return
	}
	if node.left == nil {
		tree.Transplant(node, node.right)
	} else if node.right == nil {
		tree.Transplant(node, node.left)
	} else {
		successor := tree.findMin(node.right)
		if successor.parent != node {
			tree.Transplant(successor, successor.right)
			successor.right = node.right
			successor.right.parent = successor
		}
		tree.Transplant(node, successor)
		successor.left = node.left
		successor.left.parent = successor
	}
}

func main() {
	tree := &BST{}

	tree.InsertNode(15)
	tree.InsertNode(13)
	tree.InsertNode(7)
	tree.InsertNode(48)
	tree.InsertNode(109)
	tree.InsertNode(1)
	tree.InsertNode(151)
	tree.InsertNode(137)
	tree.InsertNode(5)

	// Search node
	var keyS int
	fmt.Println("Nhập key của node muốn tìm")
	fmt.Scanln(&keyS)
	if node, found := tree.SearchNode(keyS); found {
		fmt.Printf("Node %d đã được tìm thấy trong cây\n", node.key)
	} else {
		fmt.Printf("Node có key: %d không tồn tại trong cây", node.key)
	}

	minNode := tree.FindMin()
	maxNode := tree.FindMax()
	if minNode != nil {
		fmt.Printf("Min node: %d\n", minNode.key)
	}
	if maxNode != nil {
		fmt.Printf("Max node: %d\n", maxNode.key)
	}

	// Inorder traversal
	fmt.Println("Inorder traversal: ")
	tree.InorderTraversal(tree.root, func(key int) {
		fmt.Printf("%d ", key)
	})
	fmt.Println()

	// Predecessor & Successor
	if node, found := tree.SearchNode(13); found {
		predecessor := tree.Predecessor(node)
		successor := tree.Successor(node)
		if predecessor != nil {
			fmt.Printf("Predecessor of %d là: %d\n", node.key, predecessor.key)
		}
		if successor != nil {
			fmt.Printf("Successor of %d là: %d\n", node.key, successor.key)
		}
	}

	// Delete node
	var keyD int
	fmt.Println("Nhập key của node muốn xóa")
	fmt.Scanln(&keyD)
	tree.DeleteNode(keyD)
	fmt.Printf("Inorder traversal sau khi xóa node có key: %d là: ", keyD)
	tree.InorderTraversal(tree.root, func(key int) {
		fmt.Printf("%d ", key)
	})
	fmt.Println()
}