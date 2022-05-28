package main

import "fmt"

func main() {
	// ランダムに並べられた重複のない整数の配列
	array := []int{5, 4, 6, 2, 1, 9, 8, 3, 7, 10}
	// ソート実行
	sortedArray := sort(array)
	// 結果出力
	for _, v := range sortedArray {
		fmt.Println(v)
	}
}

func sort(array []int) []int {
	// 要素が一つの場合はソートの必要がないので、そのまま返却
	if len(array) == 1 {
		return array
	}

	// 配列の先頭を基準値とする
	pivot := array[0]

	// ここから記述
	pl := 0
	right := len(array) - 1
	pr := right

	for pl <= pr {
		for array[pl] < pivot {
			pl++
		}
		for array[pr] > pivot {
			pr--
		}

		if pl <= pr {
			array[pr], array[pl] = array[pl], array[pr]
			pl++
			pr--
		}
	}

	fmt.Println("-----------------")

	for i := 0; i <= pl-1; i++ {
		fmt.Println(array[i])
	}

	fmt.Println("-----------------")

	for i := pr + 1; i < len(array)-1; i++ {
		fmt.Println(array[i])
	}

	fmt.Println("-----------------")

	if pr > 0 {
		sort(array[0:pr])
	}

	if pl < right {
		sort(array[pl:right])
	}

	return array

	// ここまで記述

}
