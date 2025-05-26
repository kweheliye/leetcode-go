package medium

type FirstUnique struct {
	isUnique map[int]bool
	queue    []int
}

func Constructor(nums []int) FirstUnique {
	fu := FirstUnique{
		isUnique: make(map[int]bool),
		queue:    []int{},
	}

	for _, num := range nums {
		fu.Add(num)
	}

	return fu
}

func (this *FirstUnique) ShowFirstUnique() int {
	for len(this.queue) > 0 && !this.isUnique[this.queue[0]] {
		this.queue = this.queue[1:]
	}

	if len(this.queue) == 0 {
		return -1
	}

	return this.queue[0]
}

func (this *FirstUnique) Add(value int) {
	if _, exists := this.isUnique[value]; !exists {
		// First time seeing this value, add to queue and mark as unique
		this.queue = append(this.queue, value)
		this.isUnique[value] = true
	} else {
		// Already seen this value, mark as not unique
		this.isUnique[value] = false
	}
}
