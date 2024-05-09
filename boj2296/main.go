package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Building struct {
	X int
	Y int
	C int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)

	b := []Building{}
	for i := 0; i < n; i++ {
		var x, y, c int
		fmt.Fscan(r, &x, &y, &c)
		b = append(b, Building{x, y, c})
	}

	sort.Slice(b, func(i, j int) bool {
		return b[i].X < b[j].X
	})

	dp := [1000][2]int{}
	for i := 0; i < n; i++ {
		dp[i][0] = b[i].C
		dp[i][1] = b[i].C
	}
	answer := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if b[j].Y < b[i].Y {
				dp[i][0] = max(dp[i][0], dp[j][0]+b[i].C)
			} else {
				dp[i][1] = max(dp[i][1], dp[j][1]+b[i].C)
			}
		}
		answer = max(answer, dp[i][0], dp[i][1])
	}
	fmt.Printf("%d\n", answer)
}
