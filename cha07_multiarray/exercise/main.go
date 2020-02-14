package main

import (
	"fmt"
	"math/rand"
	"time"
)

//二维数组存三个班,随机取每班5同学成绩
//求每班总分,平均分,所有总分,平均分

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr [3][5]int
	for i, v := range arr {
		for k, _ := range v {
			arr[i][k] = rand.Intn(101)
		}
	}
	fmt.Printf("采集的所有成绩=%v\n", arr)

	allsum := 0.0
	for i, v := range arr {
		sum := 0.0
		for _, l := range v {
			sum += float64(l)
		}
		allsum += sum
		fmt.Printf("第%v个班总分=%.2f,班级平均分=%.2f\n", i+1, sum, sum/float64(len(v)))
	}
	fmt.Printf("总分=%.2f,总平均分=%.2f", allsum, allsum/float64(len(arr)*len(arr[0])))
}
