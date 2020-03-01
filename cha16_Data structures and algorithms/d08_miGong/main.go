package main

import "fmt"

//FindOut 走二维数组迷宫
func FindOut(mymap *[8][7]int, i, j int) bool { //*[8][7]int迷宫地图,保证同一个地图,引用传递,i,j表示点
	//1. 分析什么情况下找到出路 map[6][5]为出口,==2时候为找到出路
	if mymap[6][5] == 2 {
		return true
	}
	//继续找
	if mymap[i][j] == 0 { //说明该节点可以探测
		//先假设这个点可以通过,但是需要探测方向 -->  下右上左
		mymap[i][j] = 2
		if FindOut(mymap, i+1, j) { //下
			return true
		} else if FindOut(mymap, i, j+1) { //右
			return true
		} else if FindOut(mymap, i-1, j) { //上
			return true
		} else if FindOut(mymap, i, j-1) { //左
			return true
		} else { //死路
			mymap[i][j] = 3
			return false
		}
	} //else  //说明该节点不通,是墙
	return false

}

func main() {
	//二维数组模拟迷宫
	/* 规则
	1. 元素值为0, 表示没有走过的点
	2. 元素值为1, 表示是墙 ,无法通行
	3. 元素值为2, 表示是一个通路
	4. 元素值为3, 表示是走过的点,但是走不通,死路
	*/

	var myMap [8][7]int
	//地图最上最下设置为墙1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	//地图最左最右设置为墙1
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	myMap[1][2] = 1
	myMap[2][2] = 1

	//看地图
	for i := 0; i < 8; i++ {
		for x := 0; x < 7; x++ {
			fmt.Printf("%d  ", myMap[i][x])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	FindOut(&myMap, 1, 1)
	for i := 0; i < 8; i++ {
		for x := 0; x < 7; x++ {
			fmt.Printf("%d  ", myMap[i][x])
		}
		fmt.Println()
	}
}
