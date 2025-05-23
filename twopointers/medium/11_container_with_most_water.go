package medium

// Version 1 Brute force
func maxAreaV1(height []int) int {
	res := 0
	n := len(height)

	for left := 0; left < n; left++ {
		for right := left + 1; right < n; right++ {
			area := (right - left) * min(height[left], height[right])
			res = max(res, area)
		}
	}

	return res
}

func maxAreaV2(height []int) int {

	p1, p2 := 0, len(height)-1
	result := 0

	for p1 < p2 {
		area := (p2 - p1) * min(height[p1], height[p2])
		result = max(result, area)

		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}

	}

	return result
}
