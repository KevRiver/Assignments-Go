package boj2805

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	r := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(r, &n, &m)

	trees := []int{}
	var maxHeight int = 0
	for i := 0; i < n; i++ {
		var height int
		fmt.Fscan(r, &height)
		maxHeight = max(maxHeight, height)
		trees = append(trees, height)
	}

	var low int = 0
	var high int = maxHeight + 1
	var answer int = 0
	for low < high {
		h := (low + high) / 2

		var gain = 0
		for _, treeHeight := range trees {
			gain += max(treeHeight-h, 0)
		}

		if gain >= m {
			answer = h
			low = h + 1
		} else {
			high = h
		}
	}

	fmt.Printf("%d\n", answer)
}
