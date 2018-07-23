//leetcode上的简单题，我想试试用goroutine来并发遍历两棵树，然后判断是否是相同的树
//结果发现比不用并发的方法还慢，
//主要是因为我遍历其中一棵树的节点时要阻塞等待另一棵树遍历节点的结果
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
        tmp, ok := <-ch1
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
    stack = append(stack, t)    //用slice实现一个简单的栈来进行层次遍历
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
