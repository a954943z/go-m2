package count

// 宣告
var playerMap map[string]int

func Init() {
	// 一定要初始化 map
	playerMap = make(map[string]int)
}

// 需要叫用時可直接使用
func Get() map[string]int {
	return playerMap
}

func Save(c string) int {
	// 取值
	times, exists := playerMap[c]
	if exists {
		times++
		if times > 10 {
			return 0
		}
		playerMap[c] = times

		return playerMap[c]
	} else {
		playerMap[c] = 1
		return playerMap[c]
	}

}

//製作一個ＡＰＩ，記錄玩家使用次數，並列印當前使用次數
//（寫三隻測試）
//不可超10次
//限制一分鐘只能用十次
