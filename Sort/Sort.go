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
	right := len(array) - 1 //配列の右端
	pr := right             //右カーソル
	pl := 0                 //左カーソル

	for pl <= pr {
		for array[pl] < pivot {
			pl++ //基準値より大きい値を左から探す
		}
		for array[pr] > pivot {
			pr-- //基準値より小さい値を右から探す
		}

		if pl <= pr {
			array[pr], array[pl] = array[pl], array[pr] //配列の要素入れ替え
			pl++
			pr--
		}
	}

	if pr >= 0 {
		sort(array[0 : pr+1]) //基準値未満の配列でsortを呼び出す
	}

	if pl <= right {
		sort(array[pl : right+1]) //基準値以上の配列でsortを呼び出す
	}

	return array

	// ここまで記述

}
