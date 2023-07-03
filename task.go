package main

import (
	"fmt"
)

func findLCS(str1 string, str2 string) string {
	m := len(str1)
	n := len(str2)
	lcs := make([][]string, m+1)
	for i := 0; i <= m; i++ {
		lcs[i] = make([]string, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + string(str1[i-1])
			} else {
				lcs[i][j] = maxLenString(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	return lcs[m][n]
}

func maxLenString(str1, str2 string) string {
	if len(str1) > len(str2) {
		return str1
	}
	return str2
}

func main() {
	string1 := "AGGTAB"
	string2 := "GXTXAYB"
	lcs := findLCS(string1, string2)
	fmt.Println("Наибольшая общая подпоследовательность для строк", string1, "и", string2, ":", lcs)
}
