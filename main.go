package main

import (
	"errors"
	"github.com/SeppeSoete/queens/board"
    "fmt"
)

func main() {
    var n node
    var stack = make([]node,0)
    stack = append(stack, n)
    var solutions = make([]node, 0)
    stack, n, e := pop(stack)

    for ; e == nil; stack,n,e = pop(stack) {
        stack = append(stack, n.expand()...)
        if n.currentCol == 8 && !contains(solutions, n){
            solutions = append(solutions, n)
        }
    }
    for _,solution := range(solutions) {
        fmt.Printf(solution.b.ToString() + "\n")
    }
    fmt.Printf("found %d solutions.\n", len(solutions))
}

func contains(list []node, n node) bool {
    for _,listNode := range(list) {
        if listNode.b.Equals(n.b) {
            return true
        }
    }
    return false
}

type node struct {
    b board.Board
    currentCol int
}

func (n node) expand() []node{
    stack := make([]node, 0)
    if n.currentCol == 8 {
        return stack
    }
    for r:= 0; r < 8; r++ {
        var newNode node
        newNode.currentCol = n.currentCol + 1
        if n.b.CanPlaceQueen(r, newNode.currentCol - 1) {
            newNode.b = n.b.PlaceQueen(r, newNode.currentCol - 1)
            stack = append(stack, newNode)
        }
    }
//    for _,s := range(stack) {
//        fmt.Printf("Expanded\n" + s.b.ToString())
//    }
    return stack
}

func pop(s []node) ([]node , node, error) {
    if len(s) == 0 {
        return s,node{}, errors.New("Attempted to pop from an empty stack")
    }
    var n node = s[len(s) - 1]
    s = s[:len(s) - 1]
    //fmt.Printf("popped\n" + n.b.ToString())
    return s,n,nil
}
