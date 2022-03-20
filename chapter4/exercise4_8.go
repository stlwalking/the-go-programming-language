// exercise 4.8, 统计输入数据中字母，数字和其他类型的字符数量

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	counts := map[string]int{}

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // 返回rune, bytes, error
		if r == 'q' {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if unicode.IsLetter(r) {
			counts["str"] += 1
			continue
		}
		if unicode.IsNumber(r) {
			counts["number"] += 1
			continue
		}
		if !(unicode.IsLetter(r) || unicode.IsNumber(r)) {
			counts["others"] += 1
		}
		if r == unicode.ReplacementChar && n == 1 {
			counts["invalid"] += 1
		}
	}

	for k, v := range counts {
		fmt.Printf("%q\t%d\n", k, v)
	}

}