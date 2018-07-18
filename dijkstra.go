package main 

import "fmt"

var Max_int int = 9999

//dijkstra是要点是遍历源点与其它点的距离数组，每次选取数组里的最小值且不重复选取，
//由选取到的最小值去更新距离数组
//（更新规则是源点经由最小值点去其它点的距离小于原来的就更新）
func dijkstra(arc [][]int, n int, start int) ([]int, [][]int){
	//用源点的那一行数组来保留源点到各个点的最短距离可以省一个n数组的空间
	var visited = make([]int, n)//已确定最短距离点集合
	for i := 0; i < n; i++ {
		visited[i] = 0
	}

	var path = make([][]int, n)
	for i := 0; i < n; i++ {
		path[i] = append(path[i], start)
	}

	var stepnode int
	visited[start] = 1 //源点置为已访问
	for i := 0; i < n-1; i++ {
		min := Max_int
		for j := 0; j < n; j++ {
			if(visited[j] == 0 && arc[start][j] < min){
				min = arc[start][j]
				stepnode = j
			}
		}
		visited[stepnode] = 1
		path[stepnode] = append(path[stepnode], stepnode)
		for k := 0; k < n; k++ {
			if arc[stepnode][k] < Max_int {
				if arc[start][k] > arc[start][stepnode] + arc[stepnode][k] {
					arc[start][k] = arc[start][stepnode] + arc[stepnode][k]
					path[k] = path[stepnode] // 前k-1个点的路径
					//源点路径为源点本身，如果所求点到达不了源点，那么该点路径也为源点本身
				}
			}
		}
	}
	return	arc[start], path
}

func main() {
	var n, m, start int
	var arc [][]int
	fmt.Scan(&n, &m, &start)
	arc = make([][]int, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				tmp[j] = 0
			}else {
				tmp[j] = Max_int
			}
		}
		arc[i] = tmp
	}

	for i := 0; i < m; i++ {
		var x, y, w int
		fmt.Scan(&x, &y,&w)
		arc[x][y] = w
	}

	dis, path := dijkstra(arc, n, start)
	fmt.Println(dis, path)
}