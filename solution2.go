
package main

import "fmt"

type Treenode struct{
	data string
	left *Treenode
	right *Treenode
}
func preorder (head *Treenode){
	if head==nil {
		return
	}
	fmt.Print(head.data)
	fmt.Print(" ")
	preorder(head.left)
	preorder(head.right)
}
func postorder (head *Treenode){
	if head == nil {
		return
	}
	postorder(head.left)
	postorder(head.right)
	fmt.Print(head.data)
	fmt.Print(" ")
}
func main(){
	head := &Treenode{"+",nil,nil}
	head.left=&Treenode{"a",nil,nil}
	head.right =& Treenode{"-",nil,nil}
	head.right.left = &Treenode{"b",nil,nil}
	head.right.right = &Treenode{"c",nil,nil}
	fmt.Print("preorder of the given expression tree is ")
	preorder(head)
	fmt.Println()
	fmt.Print("postorder of the given expression tree is ")
	postorder(head)
}