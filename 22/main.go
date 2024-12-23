package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type seq [4]int
type votingInstant struct {
	n        int
	price    int
	diff     int
	sequence seq
}

func main() {
	data, _ := os.ReadFile("test_data")
	data, _ = os.ReadFile("data")
	input := string(data)

	initialSecrets := make([]int, 0)
	for _, s := range strings.Fields(input) {
		n, _ := strconv.Atoi(s)
		initialSecrets = append(initialSecrets, n)
	}

	// initialSecrets = []int{1, 2, 3, 2024}

	L := 2001
	sum := 0

	// L = 10
	fmt.Println(initialSecrets)
	secretLists := make([][]votingInstant, len(initialSecrets))

	maxBananas := 0
	var bestSeq seq
	maxBananasPerSeq := make(map[seq]int)

	for i, secret := range initialSecrets {
		secretLists[i] = make([]votingInstant, L)
		voterBananasPerSeq := make(map[seq]int)

		for j := 0; j < L; j++ {
			l := secretLists[i]
			if j == 0 {
				l[0] = votingInstant{
					n:     secret,
					price: secret % 10,
					diff:  0,
				}
				continue
			}

			n := next(l[j-1].n)
			price := n % 10
			diff := price - l[j-1].price
			var sequence seq
			if j > 3 {
				sequence = seq{l[j-3].diff, l[j-2].diff, l[j-1].diff, diff}
				if _, exists := voterBananasPerSeq[sequence]; !exists {
					voterBananasPerSeq[sequence] = price
					maxBananasPerSeq[sequence] += price
					maxBananasForSequence := maxBananasPerSeq[sequence]
					if maxBananasForSequence > maxBananas {
						maxBananas = maxBananasForSequence
						bestSeq = sequence
					}
				}
			}
			secretLists[i][j] = votingInstant{
				n,
				price,
				diff,
				sequence,
			}
		}
		sum += secretLists[i][L-1].n
	}

	// for _, s := range secretLists[0] {
	// 	fmt.Println(s)
	// }
	fmt.Println(sum)
	fmt.Println(maxBananas)
	fmt.Println(bestSeq)
}

func next(n int) int {
	n = (n ^ (n << 6)) % 16777216
	n = (n ^ (n >> 5)) % 16777216
	n = (n ^ (n << 11)) % 16777216
	return n
}
