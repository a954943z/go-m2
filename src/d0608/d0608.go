package d0608

import (
	"fmt"
)

func Test1(arr []int) string {

	fmt.Print(arr)

	map1 := map[string]int{
		"abc": 1,
		"a": 2,
	}

	// list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 3,7,8,9,3,3,3}
	// list2 := []string{"v","a","l","u","e"}
	str := "value"

	// for i, _ := range arr {

	// 	map1[list1[i]]+= 1
	// 	// map1 [value] += 1
	// }

	// v:1 a:2 l:3 u:4 e:5

	for i := 0; i < len(str); i++ {
		value := str[i]
		tostr := string(value)
		// svalue := fmt.Sprint(value)
		map1[tostr] += 1

	}
	fmt.Print(map1)
	return ""

}

// 運用ＭＡＰ找出最大的ＶＡＬＵＥ，並ＲＥＴＵＲＮ ＫＥＹ出來