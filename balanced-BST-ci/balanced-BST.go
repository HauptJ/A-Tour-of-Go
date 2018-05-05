/*
1) Get the Middle of the array and make it root.
2) Recursively do same for left half and right half.
      a) Get the middle of left half and make it left child of the root
          created in step 1.
      b) Get the middle of right half and make it right child of the
          root created in step 1.
*/

package main

import (
  "fmt"
  "strings"
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

/* Helper Functions */

// returns the smaller int
// func min_val(a, b int) int {
//   if a < b {
//     return a
//   }
//   return b
// }
//
// // returns the larger int
// func max_val(a, b int) int {
//   if a > b {
//     return a
//   }
//   return b
// }


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
func (node *TNode) rotate_left(child *TNode) {
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

func (node *TNode) rotate_right(child *TNode) {
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
func (node *TNode) rotate_right_left(child *TNode) {
  child.RIGHT.LEFT.BAL = 1
  child.rotate_right(child.RIGHT)
  child.RIGHT.BAL = 1
  node.rotate_left(child)
}

/*
Rotates the left child of the child node to the left, then rotates the child node to the right
*/
func (node *TNode) rotate_left_right(child *TNode) {
  child.LEFT.RIGHT.BAL = -1
  child.rotate_left(child.LEFT)
  child.LEFT.BAL = -1
  node.rotate_right(child)
}

/*
Brings the subtree with root node child back to a balanced state
*/
func (node *TNode) balance_tree_nodes(child *TNode) {
  fmt.Println("balance " + child.VALUE)
  child.dump_subtree(0, "")
  switch {
  case child.BAL == -2 && child.LEFT.BAL == -1: // the left subtree is too high, and the left child node has a left child node
    node.rotate_right(child)
  case child.BAL == 2 && child.RIGHT.BAL == 1: // the right subtree is too high, and the right child node has a right child node
    node.rotate_left(child)
  case child.BAL == -2 && child.LEFT.BAL == 1: // the left subtree is too high, and the left child node has a right child node
    node.rotate_left_right(child)
  case child.BAL == 2 && child.RIGHT.BAL == -1: // The right subtree is too high, and the right child has a left child
    node.rotate_right_left(child)
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
func (node *TNode) dump_subtree(pos int, child string) {
  if node == nil {
    return
  }
  indent := ""
  if pos > 0 {
    indent = strings.Repeat(" ", (pos-1)*4) + "+" + child + "--"
  }
  fmt.Printf("%s%s[%d]\n", indent, node.VALUE, node.BAL)
  node.LEFT.dump_subtree(pos + 1, "LEFT")
  node.RIGHT.dump_subtree(pos + 1, "RIGHT")
}

/* Tree methods */

func (tree *Tree) insert_tree(value, data string) {
  if tree.ROOT == nil {
    tree.ROOT = &TNode{VALUE: value, DATA: data}
    return
  }
  //tree.ROOT.insert_tree(value, data)
  tree.ROOT.insert_node(value, data)

  if tree.ROOT.BAL < -1 || tree.ROOT.BAL > 1 {
    tree.balance_tree()
  }
}

func (tree *Tree) balance_tree() {
  tmp_parent := &TNode{LEFT: tree.ROOT, VALUE: "temp parent"}
  //tmp_parent.balance_tree(tree.ROOT)
  tmp_parent.balance_tree_nodes(tree.ROOT)
  tree.ROOT = tmp_parent.LEFT // fetch the new root node from the temp parent node
}

func (tree *Tree) find_tree(root string) (string, bool) {
  if tree.ROOT == nil {
    return "", false
  }
  //return tree.ROOT.find_tree(root)
  return tree.ROOT.find_node(root)
}

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
  //tree.ROOT.dump_tree(0, "")
  tree.ROOT.dump_subtree(0, "")
}

func main() {
  values := []string{"d", "b", "g", "g", "c", "e", "a", "h", "f", "i", "j", "l", "k"}
  data := []string{"delta", "bravo", "golang", "golf", "charlie", "echo", "alpha", "hotel", "foxtrot", "india", "juliett", "lima", "kilo"}

  tree := &Tree{}
  for i := 0; i < len(values); i++ {
    fmt.Println("Insert " + values[i] + ": " + data[i])
    tree.insert_tree(values[i], data[i])
    tree.dump_tree()
    fmt.Println()
  }

  fmt.Println("Sorted values: | ")
  tree.traverse_tree(tree.ROOT, func(node *TNode) {fmt.Print(node.VALUE, ": ", node.DATA, " | ") })
}
