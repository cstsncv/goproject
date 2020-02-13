package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1. 定义byte类型26个元素的数组.分别放置 'A'-'Z' , 字符运算 'A' + 1 ==> 'B'
	var myChar [26]byte
	for i := 0; i < 26; i++ {
		myChar[i] = 'A' + byte(i) //注意!!需要将  i ==> byte
	}

	for index, value := range myChar {
		fmt.Printf("index = %d, value = %c\t\n", index, value) // %c按字符输出
	}

	//2. 生成一个数组 [5]int ,随机生成数字填充, 找出最大值
	// 假定第一个元素是最大值,下标0,  然后从第二个开始比较,如发现有更大, 则交换
	var numArr [15]int
	var numArr1 [len(numArr)]int
	lang := len(numArr)
	rand.Seed(time.Now().UnixNano()) //初始化rand的Seed
	for i := 0; i < lang; i++ {
		numArr[i] = rand.Intn(222)
	}
	fmt.Printf("numArr = %v\n", numArr)
	//求最大元素
	maxIndex := 0
	for i := 1; i < lang; i++ {
		if numArr[maxIndex] < numArr[i] {
			maxIndex = i
		}
	}
	//求元素总和
	numSum := 0
	for _, value := range numArr {
		numSum += value
	}
	//反转元素
	for index := 0; index < lang; index++ {
		numArr1[index] = numArr[lang-1-index]
	}
	fmt.Printf("numArr数组内最大值 = %v\n", numArr[maxIndex])
	fmt.Printf("numArr元素和 = %d, numArr 元素平均值 = %.2f\n", numSum, float64(numSum)/float64(lang)) //int转float64
	fmt.Printf("numArr1 = %v\n", numArr1)
}
