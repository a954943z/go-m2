package bbb

//map

type qqq struct {
	a1 string
	A2 int
	a3 www
}
type www struct {
	a3 string
}

// l1 := []int{1, 2, 3}
// l2 := []string{"4", "5", "6"}
func Maptest(list []int, List2 []string) (map[string]qqq, int) {

	newmap := map[string]qqq{}

	for i := 0; i < len(list); i++ {
		newmap[List2[i]] = qqq{A2: list[i]}

	}

	return newmap, 100
}

// func Maptest() map[string]int {

// 	aa := map[string]int{}

// 	aa["abc"] = 1
// 	aa["def"] = 2

// 	return aa

// }
