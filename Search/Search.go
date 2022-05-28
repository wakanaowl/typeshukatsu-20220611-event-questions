package main

import "fmt"

func main() {
	// 昇順にソートされた配列
	sortedArray := []int{1, 2, 3, 5, 12, 7890, 12345}
	// 探索対象の番号
	targetNumber := 7890
	// 探索実行
	targetIndex := serchIndex(sortedArray, targetNumber)
	// 結果出力
	fmt.Println(targetIndex)
}

func serchIndex(sortedArray []int, targetNumber int) int {

	// ここから記述
	left := 0                     //配列の左端のindex
	right := len(sortedArray) - 1 //右端のindex

	var mid int = left + (right-left)/2 //真ん中のindex

	for left <= right {

		mid = left + (right-left)/2 //探索範囲の真ん中を更新

		if sortedArray[mid] == targetNumber {
			return mid //探索対象のindexを返却
		} else if sortedArray[mid] > targetNumber {
			right = mid - 1 //探索範囲を左半分に
		} else if sortedArray[mid] < targetNumber {
			left = mid + 1 //探索範囲を右半分に
		}
	}
	// ここまで記述

	// 探索対象が存在しない場合、-1を返却
	return -1
}
