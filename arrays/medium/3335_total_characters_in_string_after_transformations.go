package medium

const mod = 1000000007

func lengthAfterTransformations(s string, t int) int {
	cnt := [26]int{}

	for _, c := range s {
		cnt[c-'a']++
	}

	for round := 0; round < t; round++ {
		next := [26]int{}
		for i := 0; i < 25; i++ {
			next[i+1] = cnt[i] % mod
		}
		next[0] = (next[0] + cnt[25]) % mod
		next[1] = (next[1] + cnt[25]) % mod
		cnt = next
	}
	var ans int
	for _, c := range cnt {
		ans = (ans + c) % mod
	}

	return ans
}
