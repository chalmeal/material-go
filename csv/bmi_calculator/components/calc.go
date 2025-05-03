package components

import "strconv"

// stringをintに変換
// エラーの場合は-1を返す
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}
