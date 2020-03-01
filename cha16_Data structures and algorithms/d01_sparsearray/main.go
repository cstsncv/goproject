package main

import "fmt"

//二维数组的稀疏数组
//五子棋

//node结构体
type node struct {
	row int //行
	col int //列
	val int //值 interface{}可放任何
}

func main() {
	//1. 原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //蓝子
	//2. 遍历原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	//3.转成稀疏数组
	//3.1 遍历chessMap,如果有元素不等于0,创建一个node结构体,保存 行,列,数值
	//3.2 将该结构体放入至切片
	var sparseArr []*node
	val := &node{ //初始节点 记录稀疏数组的规模(总行,总列,默认值)
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, val)
	for r, v := range chessMap {
		for c, v2 := range v {
			if v2 != 0 {
				val := &node{
					row: r,
					col: c,
					val: v2,
				}
				sparseArr = append(sparseArr, val)
			}
		}
	}
	//输出稀疏数组
	for _, v := range sparseArr {
		fmt.Printf("%d %d %d \n", v.row, v.col, v.val)
	}
	/* 存盘为
	11 11 0
	1 2 1
	2 3 2
	*/
	//读取存盘直接恢复原二维数组
	var chessMap2 [11][11]int
	for i, v := range sparseArr { //用数组演示读取文件
		if i != 0 { //跳过第一行记录
			chessMap2[v.row][v.col] = v.val
		}
	}
	fmt.Println(chessMap2)
}
