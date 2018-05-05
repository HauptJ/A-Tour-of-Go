package main

import (
  "fmt"
  "strings"
  "sort"
)

/* Node structure */
type TNode struct {
  VALUE string
  DATA string
  LEFT *TNode
  RIGHT *TNode
  BAL int //hight diff betweeen the node's subtrees
}

/* Tree structure */
type Tree struct {
  ROOT *TNode
}

/*
Insert method
Returns:
  true: if the tree height has increased
  false: otherwise
*/
func (node *TNode) insert_node(value, data string) bool {
  switch {
  case value == node.VALUE: // new search value is equal to the current node's search value
    node.DATA = data
    return false // node already exists, so change nothing
  case value < node.VALUE: // new search value is less than the current node's search value
    // if there is no left child, create a new one
    if node.LEFT == nil {
      node.LEFT = &TNode{VALUE: value, DATA: data} // create new node
      // if there is no right child, the new child node increased the hight of the subtree
      if node.RIGHT == nil {
        node.BAL = -1
      } else { // the new left child is the only child
        node.BAL = 0
      }
    } else { // there is a left and a right child
      if node.LEFT.insert_node(value, data) { // The left child exists / is not nill. Continue in left subtree
          if node.LEFT.BAL < -1 || node.LEFT.BAL > 1 { // The subtree's balance factor is either -2 or 2 so the tree must be rebalanced
            node.balance_tree_nodes(node.LEFT)
          } else {
            node.BAL-- // decrease the balance of the current node by one if no rebalancing occurred
          }
        }

    }
  case value > node.VALUE: // mirrored case for value < node.VALUE
    if node.RIGHT == nil {
      node.RIGHT = &TNode{VALUE: value, DATA: data}
      if node.LEFT == nil {
        node.BAL = 1
      } else {
        node.BAL = 0
      }
    } else {
      if node.RIGHT.insert_node(value, data) {
        if node.RIGHT.BAL < -1 || node.RIGHT.BAL > 1 {
          node.balance_tree_nodes(node.RIGHT)
        } else {
          node.BAL++
        }
      }
    }
  }

    if node.BAL != 0 { // the tree's height has increased
      return true
    }

    return false // the tree's height has not increased
}

/* rebalance_tree_nodes helper methods */


/*
Takes a child node and rotates the child node's subtree to the left
*/
func (node *TNode) rotate_node_left(child *TNode) {
  right_child := child.RIGHT // save the right child node

  child.RIGHT = right_child.LEFT // right_child's left subtree is assigned to child

  right_child.LEFT = child // child becomes the left child of right_child

  // Make parent / current node point to the new root node
  if child == node.LEFT {
    node.LEFT = right_child
  } else {
    node.RIGHT = right_child
  }

  child.BAL = 0
  right_child.BAL = 0
}

func (node *TNode) rotate_node_right(child *TNode) {
  fmt.Println("right rotation " + child.VALUE)
  left_child := child.LEFT
  child.LEFT = left_child.RIGHT
  left_child.RIGHT = child

  if child == node.LEFT {
    node.LEFT = left_child
  } else {
    node.RIGHT = left_child
  }

  child.BAL = 0
  left_child.BAL = 0
}

/*
Rotates the right child of the child node to the right, then rotates the child node to the left
*/
func (node *TNode) rotate_node_right_left(child *TNode) {
  child.RIGHT.LEFT.BAL = 1
  child.rotate_node_right(child.RIGHT)
  child.RIGHT.BAL = 1
  node.rotate_node_left(child)
}

/*
Rotates the left child of the child node to the left, then rotates the child node to the right
*/
func (node *TNode) rotate_node_left_right(child *TNode) {
  child.LEFT.RIGHT.BAL = -1
  child.rotate_node_left(child.LEFT)
  child.LEFT.BAL = -1
  node.rotate_node_right(child)
}

/*
Brings the subtree with root node child back to a balanced state
*/
func (node *TNode) balance_tree_nodes(child *TNode) {
  fmt.Println("balance " + child.VALUE)
  child.dump_node(0, "")
  switch {
  case child.BAL == -2 && child.LEFT.BAL == -1: // the left subtree is too high, and the left child node has a left child node
    node.rotate_node_right(child)
  case child.BAL == 2 && child.RIGHT.BAL == 1: // the right subtree is too high, and the right child node has a right child node
    node.rotate_node_left(child)
  case child.BAL == -2 && child.LEFT.BAL == 1: // the left subtree is too high, and the left child node has a right child node
    node.rotate_node_left_right(child)
  case child.BAL == 2 && child.RIGHT.BAL == -1: // The right subtree is too high, and the right child has a left child
    node.rotate_node_right_left(child)
  }
}

/*
Find the node with the specified string.
RETURNS the node's DATA as a string and presence of the node as a bool
*/
func (node *TNode) find_node(data string) (string, bool) {
  if node == nil { // node does not exist
    return "", false
  }

  switch {
  case data == node.VALUE: // node was found
    return node.DATA, true
  case data < node.VALUE:
    return node.LEFT.find_node(data)
  default:
    return node.RIGHT.find_node(data)
  }
}

/*
Dumps the structure of the subtree starting at node node
*/
func (node *TNode) dump_node(pos int, child string) {
  if node == nil {
    return
  }
  indent := ""
  if pos > 0 {
    indent = strings.Repeat(" ", (pos-1)*4) + "+" + child + "--"
  }
  fmt.Printf("%s%s[%d]\n", indent, node.VALUE, node.BAL)
  node.LEFT.dump_node(pos + 1, "LEFT")
  node.RIGHT.dump_node(pos + 1, "RIGHT")
}

/* Tree methods */

/*
Inserts a new node into the tree and sets it as the root of the tree
*/
func (tree *Tree) insert_tree(value, data string) {
  if tree.ROOT == nil {
    tree.ROOT = &TNode{VALUE: value, DATA: data}
    return
  }
  tree.ROOT.insert_node(value, data)

  if tree.ROOT.BAL < -1 || tree.ROOT.BAL > 1 { // rebalance tree
    tree.balance_tree()
  }
}

func (tree *Tree) balance_tree() {
  tmp_parent := &TNode{LEFT: tree.ROOT, VALUE: "temp parent"}
  tmp_parent.balance_tree_nodes(tree.ROOT)
  tree.ROOT = tmp_parent.LEFT // fetch the new root node from the temp parent node
}

func (tree *Tree) find_tree(root string) (string, bool) {
  if tree.ROOT == nil {
    return "", false
  }
  return tree.ROOT.find_node(root)
}

/*
Traverses the tree for a specified node
*/
func (tree *Tree) traverse_tree(node *TNode, find func(*TNode)) {
  if node == nil {
    return
  }
  tree.traverse_tree(node.LEFT, find)
  find(node)
  tree.traverse_tree(node.RIGHT, find)
}

/*
Dumps the tree structure
*/
func (tree *Tree) dump_tree() {
  tree.ROOT.dump_node(0, "")
}

func print_sorted_map(m map[string]string) {
  var keys[]string
  for k := range m {
    keys = append(keys, k)
  }
  sort.Strings(keys)
  for _, k := range keys {
    fmt.Println("Key:", k, "Value:", m[k])
  }
}

func print_balanced_BST(v, d *[]string) {
  for i := range *v {
    // a little pointer fun
    fmt.Println("Value:", (*v)[i], "Data", (*d)[i])
  }
}

func bst(values, data []string) ([]string, []string) {
  tree := &Tree{}
  for i := 0; i < len(values); i++ {
    fmt.Println("Insert " + values[i] + ": " + data[i])
    tree.insert_tree(values[i], data[i])
    tree.dump_tree()
    fmt.Println()
  }

  fmt.Println("Sorted values: | ")
  var v, d []string

  // arrays follow insertion order
  tree.traverse_tree(tree.ROOT, func(node *TNode) {v, d = append(v, node.VALUE), append(d, node.DATA)})
  //fmt.Printf("%v%v\n", v, d)
  //print_balanced_BST(&v, &d)

  // m := make(map[string]string)
  //
  // // golang maps do not follow insertion order :(
  // tree.traverse_tree(tree.ROOT, func(node *TNode) {m[node.VALUE] = node.DATA})
  // // sorting the map kind of defeats the purpose of this exercise
  // fmt.Printf("%v\n", m)
  // print_sorted_map(m)

  return v, d

}

func main() {
  //values := []string{"d", "b", "g", "g", "c", "e", "a", "h", "f", "i", "j", "l", "k"}
  //data := []string{"delta", "bravo", "golang", "golf", "charlie", "echo", "alpha", "hotel", "foxtrot", "india", "juliett", "lima", "kilo"}

  //v, d := bst(values, data)
  //print_balanced_BST(&v, &d)
  tst()

}
