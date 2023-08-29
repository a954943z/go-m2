package work

import "fmt"

//設定一個function,參數為abclist陣列,回傳值為map[int]int型態
func Homework1(abclist[]int)map[int]int  {
	//宣告一個mapresult,及定義map[int]int型別
	mapresult := map[int]int{
		//map內容值可以一開始就設定,會自動與下方迴圈值合併
		3:3,
		4:2,
	}
	for _, v := range abclist {
		mapresult[v]++
	}
	
	maxCount := 0
	maxNumber := 0

	for number, count := range mapresult {
		if count > maxCount {
			maxCount = count
			maxNumber = number
		}
	}

	fmt.Printf("出現次數最多的數字是: %d，出現次數為: %d\n", maxNumber, maxCount)

	return mapresult
}

func QQQ(int16)int16{
	

return 111
}