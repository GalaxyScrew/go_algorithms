package main 

import "fmt"
import "math/rand"

//建堆：从i = length / 2 - 1 开始，从右向左开始调整，
//即for i = length / 2 - 1；i >= 0；i-- {
//		go_down()
//	}
func init_heap(arr []int, length int) bool{

	for i := length/2 - 1; i >= 0; i-- {
		go_down(arr, i, length)
	}

	return true
}

//下沉：从当前节点的左右子节点（即2 * index + 1和2 * index + 2）中选出较大者，
//若较大者比当前节点大就交换，然后再从交换者的位置开始下沉，否则return
func go_down(arr []int, nownode int, length int) {
	
	leftchild := 2 * nownode + 1	//左子节点2 * index + 1
	rightchild := 2 * nownode + 2	//右子节点2 * index + 2

	if leftchild > length - 1 {		//判断是否有子节点，无就return
		return
	}

	if rightchild <= length - 1 && arr[rightchild] > arr[leftchild] {	//选出左右子节点较大者
		leftchild = rightchild
	} 

	if arr[leftchild] <= arr[nownode] { 	//当前节点比左右子节点都大就直接return
		return
	}

	arr[leftchild], arr[nownode] = arr[nownode], arr[leftchild]	//当前节点比子节点小就交换

	go_down(arr, leftchild, length)		//对交换了的子节点再进行下沉
}

//堆排序：将根节点与末尾节点交换（即0与index-1交换），
//然后index--，从根节点开始下沉操作
//最后得到的数组是最大堆-递增数组、最小堆-递减数组
func heap_sort(arr []int, length int) {

	for i := length - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		go_down(arr, 0, i)
	}
	return
}

func main() {
    var length = 10
    var tree []int

    //构建一个随机数组
    for i := 0; i < length; i++ {
        tree = append(tree, int(rand.Intn(100)))
    }

    fmt.Println(tree)

    init_heap(tree, length)

    fmt.Println(tree)

    heap_sort(tree, length)

    fmt.Println(tree)
}