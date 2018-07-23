package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    flag := true
    ch1 := make(chan int)
    ch2 := make(chan int)
    go Walk(p, ch1)
    go Walk(q, ch2)
    if q == nil && p != nil {
        return false
    }
    for i := range ch2{
        //fmt.Println("ch2:",i)
        tmp, ok := <-ch1
        //fmt.Println("ch1:",tmp)
        if ok == false {
            flag = false
            break
        }

        if  i != tmp {
            flag = false
            break
        }      
    }
    return flag
}

func Walk(t *TreeNode, ch chan int) {
    stack := make([]*TreeNode,0)
    stack = append(stack, t)
    fmt.Println(len(stack))
    for len(stack) != 0 {
        tmpnode := stack[0]
        stack = stack[1:]
        if tmpnode != nil {
            stack = append(stack, tmpnode.Left, tmpnode.Right)
            ch <- tmpnode.Val
        }else {
            ch <- -1
        }
    }
    close(ch)
}



func main() {
    t1 := TreeNode{1,&TreeNode{1,nil,nil},nil}
    t2 := TreeNode{1,nil,&TreeNode{1,nil,nil}}
    flag := isSameTree(&t1, &t2)
    fmt.Println(flag)
}