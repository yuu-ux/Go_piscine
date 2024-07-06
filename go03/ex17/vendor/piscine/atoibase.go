// 第1引数の数字を第2引数で受け取った文字列の桁数によってx進数として、第2引数の文字に沿って10進数に変換する
// 2進数から10進数への変換
// 桁ごとに1,2,4,8のように重み付けをして1が上がっていればかける

package piscine

func CountDigit(base string) int {
	var digit int
	for range base {
		digit++
	}
	return digit
}

func CheckBase(base []rune) bool {
	for i, c1 := range base {
		if c1 == '+' || c1 == '-' {
			return false
		}
		for _, c2 := range base[i+1:] {
			if c1 == c2 {
				return false
			}
		}
	}
	return true
}

// インデックス番号を調べる
func SearchIndex (c rune, base string) int {
	for i, c1 := range base {
		if c == c1 {
			return i
		}
	}
	return 0
}

// stringからintへキャストする関数
func StringToInt(c rune) int {
	return (int(c) - '0')
}

func AtoiBase(s string, base string) int {
	var digit = CountDigit(base)
	var result int
	var index int

	if !(CheckBase([]rune(base))) || (digit < 2) {
		return 0
	}

	for _, c := range s {
		if !(c >= '0' && c <= '9') {
			index = SearchIndex(c, base)
		} else {
			index = StringToInt(c)
		}
		result = (result * digit) + index
	}
	return result
}
