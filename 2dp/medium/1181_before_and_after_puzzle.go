package medium

import (
	"sort"
	"strings"
)

func beforeAndAfterPuzzles(phrases []string) []string {
	n := len(phrases)

	// sp[i][0] = first word, sp[i][1] = last word
	sp := make([][2]string, n)
	for i := 0; i < n; i++ {
		words := strings.Split(phrases[i], " ")
		sp[i][0] = words[0]
		sp[i][1] = words[len(words)-1]
	}

	m := make(map[string]struct{})

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if sp[i][0] == sp[j][1] {
				// combine phrases[j] + phrases[i] without overlapping first word
				combined := phrases[j] + phrases[i][len(sp[i][0]):]
				m[combined] = struct{}{}
			}
		}
	}

	// convert map keys to slice
	ret := make([]string, 0, len(m))
	for key := range m {
		ret = append(ret, key)
	}

	// sort lexicographically
	sort.Strings(ret)
	return ret
}
