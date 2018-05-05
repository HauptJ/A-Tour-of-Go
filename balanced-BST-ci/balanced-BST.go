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

/* Helper Functions */

// returns the smaller int
func min_val(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// returns the larger int
func max_val(a, b int) int {
  if a > b {
    return a
  }
  return b
}

type TNode struct {
  VALUE string
  DATA string
  LEFT *TNode
  RIGHT *TNode
  BAL int //hight diff betweeen the node's subtrees
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
    }
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
            node.balance_tree(node.LEFT)
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
          node.balance_tree(node.RIGHT)
        } else {
          node.BAL++
        }
      }
    }

    if node.BAL != 0 { // the tree's height has increased
      return true
    }

    return false // the tree's height has not increased
}

/* rebalance_tree helper methods */


/*
Takes a child node and rotates the child node's subtree to the left
*/
func (node *TNode) rotate_left(child *TNode) {
  right_child := child.RIGHT // save the right child node

  child.RIGHT = right_child.LEFT // right_child's left subtree is assigned to child

  right.LEFT = child // child becomes the left child of r

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

func (node *TNode) balance_tree(child *TNode){
  
}
